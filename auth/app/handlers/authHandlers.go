package handlers

import (
	"authService/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthHandlers struct {
}

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
	c.JSON(http.StatusOK, response)
}
