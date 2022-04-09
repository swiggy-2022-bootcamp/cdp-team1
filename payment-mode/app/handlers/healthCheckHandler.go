package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheck godoc
// @Summary To check if the service is running or not.
// @Description This request will return 200 OK if server is up..
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	Service is up
// @Router / [GET]
func HealthCheck(c *gin.Context) {
	//Check to be added for database.
	c.JSON(http.StatusOK, gin.H{"message": "Service is Up"})
}
