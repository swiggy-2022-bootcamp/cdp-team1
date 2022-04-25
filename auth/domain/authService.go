package domain

import (
	"authService/config"
	"authService/errs"
	"fmt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

type AuthService interface {
	HealthCheck() *errs.AppError
	Login(string, string, string) (string, *errs.AppError)
	CreateSession(string, string, string) *errs.AppError
	Logout(string) *errs.AppError
	VerifyAuthToken(string, string, string) (bool, *errs.AppError)
	GenerateJWT(string, string) (string, *errs.AppError)
	ExtractTokenSign(string) string
	ParseAuthToken(string) (string, string, *errs.AppError)
}

type Claims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

type DefaultAuthService struct {
	AuthDB  AuthRepositoryDB
	AdminDB UserRepositoryDB
	CustDB  UserRepositoryDB
}

func NewAuthService(authDB AuthRepositoryDB, adminDB UserRepositoryDB, custDB UserRepositoryDB) AuthService {
	return &DefaultAuthService{
		AuthDB:  authDB,
		AdminDB: adminDB,
		CustDB:  custDB,
	}
}

func (d DefaultAuthService) HealthCheck() *errs.AppError {

	if err := d.AuthDB.RepoHealthCheck(); err != nil {
		return err
	}
	if err := d.AdminDB.RepoHealthCheck(); err != nil {
		return err
	}
	if err := d.CustDB.RepoHealthCheck(); err != nil {
		return err
	}
	return nil
}

func (d DefaultAuthService) Login(username string, email string, password string) (string, *errs.AppError) {

	if (len(username) == 0 && len(email) == 0) || (len(username) > 0 && len(email) > 0) {
		return "", errs.NewValidationError("only one of username or email must be present")
	}

	var (
		user *User
		err  *errs.AppError
		role string
	)
	if len(username) > 0 {
		user, err = d.AdminDB.FetchByUsername(username)
		role = "admin"
	} else {
		user, err = d.CustDB.FetchByEmail(email)
		role = "user"
	}

	if err != nil {
		return "", err
	}
	user.Role = role
	if err1 := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err1 != nil {
		return "", errs.NewAuthenticationError("invalid credentials")
	}

	token, err := d.GenerateJWT(user.UserID, user.Role)
	if err != nil {
		return "", err
	}

	tokenSign := d.ExtractTokenSign(token)
	err = d.CreateSession(tokenSign, user.UserID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (d DefaultAuthService) CreateSession(tokenSign string, userID string, role string) *errs.AppError {

	session := NewSession(tokenSign, userID, role, true)
	err := d.AuthDB.SaveSession(*session)
	return err
}

func (d DefaultAuthService) Logout(authToken string) *errs.AppError {

	tokenSign := d.ExtractTokenSign(authToken)
	err := d.AuthDB.UpdateSessionStatusByTokenSign(tokenSign)
	return err
}

func (d DefaultAuthService) GenerateJWT(userID string, role string) (string, *errs.AppError) {

	expirationTime := time.Now().Add(time.Duration(config.EnvVars.TokenDuration) * time.Minute)

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}).SignedString(config.EnvVars.SecretBytes)
	if err != nil {
		return token, errs.NewUnexpectedError("Error generating token")
	}
	return token, nil
}

func (d DefaultAuthService) ExtractTokenSign(token string) string {

	splitToken := strings.Split(token, ".")
	tokenSign := splitToken[len(splitToken)-1]
	return tokenSign
}

func (d DefaultAuthService) ParseAuthToken(token string) (string, string, *errs.AppError) {

	tokenBytes, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		// validate the alg is what is expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return config.EnvVars.SecretBytes, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return "", "", errs.NewAuthenticationError("expired token")
			}
			return "", "", errs.NewAuthenticationError("invalid token")
		}
		return "", "", errs.NewAuthenticationError(err.Error())
	}
	if claims, ok := tokenBytes.Claims.(jwt.MapClaims); ok && tokenBytes.Valid {
		return claims["user_id"].(string), claims["role"].(string), nil
	}
	return "", "", errs.NewAuthenticationError("cannot process token")
}

func (d DefaultAuthService) VerifyAuthToken(token string, userID string, role string) (bool, *errs.AppError) {

	tokenSign := d.ExtractTokenSign(token)
	session, err := d.AuthDB.FetchSessionByTokenSign(tokenSign)
	if err != nil {
		return false, err
	}
	if !session.IsActive {
		return false, errs.NewAuthenticationError("inactive session")
	}

	claimsUserID, claimsRole, err := d.ParseAuthToken(token)

	if err != nil {
		if strings.Compare(err.Message, "expired token") == 0 {
			d.Logout(token)
		}
		return false, err
	}

	if (len(userID) > 0 && userID != claimsUserID) || (len(role) > 0 && role != claimsRole) {
		return false, errs.NewAuthorizationError("invalid token")
	}

	return true, nil
}
