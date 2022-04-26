package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/repository"
	"qwik.in/account-frontstore/domain/service"
	"qwik.in/account-frontstore/internal/errors"
)

type AccountControllerInterface interface {
	RegisterAccount(c *gin.Context)
	GetAccountById(c *gin.Context)
	UpdateAccount(c *gin.Context)
}

type AccountController struct {
	accountService service.AccountServiceInterface
}

func InitAccountController(accountServiceToBeInjected service.AccountServiceInterface) AccountControllerInterface {
	accountController := new(AccountController)
	if accountServiceToBeInjected == nil {
		accountController.accountService = service.InitAccountService(&repository.AccountRepository{}, &repository.GrpcClientRepository{})
	} else {
		accountController.accountService = accountServiceToBeInjected
	}
	return accountController
}

// RegisterAccount godoc
// @Summary Register Customer
// @Description Lets register customer in front store.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param req body model.Account true "New account to add"
// @Success	200  {object} model.Account
// @Router /customer [POST]
func (accountController *AccountController) RegisterAccount(c *gin.Context) {
	newAccount := model.Account{}
	json.NewDecoder(c.Request.Body).Decode(&newAccount)
	createdAccount, err := accountController.accountService.CreateAccount(newAccount)

	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	c.JSON(200, *createdAccount)
}

// GetAccountById godoc
// @Summary Get account by ID
// @Description Lets front store get customer account by customer id.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param accessorId path string true "accessor id"
// @Success	200  {object} model.Account
// @Router /account/{accessorId} [GET]
func (accountController *AccountController) GetAccountById(c *gin.Context) {
	userId, _, err := extractUser(c.GetHeader("Authorization"))
	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	fetchedAccount, err := accountController.accountService.GetAccountById(userId)
	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedAccount)
}

// UpdateAccount godoc
// @Summary Update Account
// @Description Lets front store update account by id.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param accessorId path string true "accessorId id"
// @Param req body model.Account true "Updated account"
// @Success	200  {object} model.Account
// @Router /account/{accessorId} [PUT]
func (accountController *AccountController) UpdateAccount(c *gin.Context) {
	userId, _, err := extractUser(c.GetHeader("Authorization"))
	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	account := model.Account{}
	json.NewDecoder(c.Request.Body).Decode(&account)
	updatedAccount, err := accountController.accountService.UpdateAccount(userId, account)

	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	c.JSON(200, *updatedAccount)
}

func extractUser(token string) (string, string, error) {
	tokenBytes, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {

		// validate the alg is what is expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("sUpErCaLiFrAgIlIsTiCeXpIaLiDoCiOuS"), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				// Token is either expired or not active yet
				return "", "", errors.NewAuthenticationError("expired token")
			}
			return "", "", errors.NewAuthenticationError("invalid token")
		}
		return "", "", errors.NewAuthenticationError(err.Error())
	}
	if claims, ok := tokenBytes.Claims.(jwt.MapClaims); ok && tokenBytes.Valid {
		return claims["user_id"].(string), claims["role"].(string), nil
	}
	return "", "", errors.NewAuthenticationError("cannot process token")
}
