package handlers

import (
	"authService/config"
	"authService/domain"
	"authService/dto"
	"authService/errs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandlers struct {
	AuthSvc domain.AuthService
}

// HealthCheckHandler
// @Schemes
// @Description Heath Check
// @Tags auth
// @Produce json
// @Success		200	{object}	dto.HealthResponseDTO	"Service is healthy"
// @Failure		404	{object}	dto.HealthResponseDTO	"Service is unavailable"
// @Router /auth [get]
func (ah *AuthHandlers) HealthCheckHandler(c *gin.Context) {

	response := dto.HealthResponseDTO{
		Status: "OK",
	}
	code := http.StatusOK
	if err := ah.AuthSvc.HealthCheck(); err != nil {
		response.Status = err.Message
		code = err.Code
	}
	c.JSON(code, response)
}

// LoginHandler
// @Schemes
// @Description Login
// @Tags auth
// @Produce json
// @Accept json
// @Param        dto.LoginRequestDTO  body dto.LoginRequestDTO true "User login"
// @Success		200	{object}	dto.LoginResponseDTO
// @Failure		404	{object}	errs.AppError
// @Router /auth/login [post]
func (ah *AuthHandlers) LoginHandler(c *gin.Context) {

	var loginDto dto.LoginRequestDTO

	if err := c.Bind(&loginDto); err != nil {
		//log.Fatal("Invalid request body")
		err := errs.NewValidationError("Invalid request body")
		c.JSON(err.Code, err.AsMessage())
		return
	}

	token, err1 := ah.AuthSvc.Login(loginDto.Username, loginDto.Email, loginDto.Password)
	if err1 != nil {
		c.JSON(err1.Code, err1.AsMessage())
		return
	}

	c.SetCookie("Authorization", token, config.EnvVars.TokenDuration*1000, "/", "", false, true)

	c.JSON(http.StatusOK, &dto.LoginResponseDTO{AuthToken: token})
}

// LogoutHandler
// @Schemes
// @Description Logout
// @Tags auth
// @Produce json
// @Success		200	{object}	dto.LogoutResponseDTO
// @Failure		404	{object}	errs.AppError
// @Security Bearer Token
// @Router /auth/logout [post]
func (ah *AuthHandlers) LogoutHandler(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if isValid, err := ah.AuthSvc.VerifyAuthToken(token, "", ""); !isValid || err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}
	if err := ah.AuthSvc.Logout(token); err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}
	c.JSON(http.StatusOK, &dto.LogoutResponseDTO{Message: "logged out successfully from this device"})
}

// VerificationHandler
// @Schemes
// @Description Login
// @Tags auth
// @Produce json
// @Accept json
// @Param        dto.VerificationRequestDTO  body dto.VerificationRequestDTO true "Verify Token"
// @Success		200
// @Failure		404	{object}	errs.AppError
// @Security Bearer Token
// @Router /auth/verify [post]
func (ah *AuthHandlers) VerificationHandler(c *gin.Context) {

	token := c.GetHeader("Authorization")

	var verifyReqDto dto.VerificationRequestDTO
	if err := c.Bind(&verifyReqDto); err != nil {
		err1 := errs.NewValidationError("invalid request body")
		c.JSON(err1.Code, err1.AsMessage())
		return
	}

	if _, err := ah.AuthSvc.VerifyAuthToken(token, verifyReqDto.UserID, verifyReqDto.Role); err != nil {
		c.JSON(err.Code, err.AsMessage())
		return
	}
	c.Status(http.StatusOK)
}
