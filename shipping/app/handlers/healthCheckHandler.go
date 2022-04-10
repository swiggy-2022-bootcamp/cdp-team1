package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//HealthCheck ..
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Shipping Server - Up and Running !",
	})
}
