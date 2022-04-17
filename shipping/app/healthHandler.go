package app

import (
	"github.com/ashwin2125/qwk/shipping/domain/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

//HealthCheckHandler ..
type HealthCheckHandler struct {
	shippingAddrRepoObj repository.ShippingAddrRepo
}

//HealthCheckResponse ..
type HealthCheckResponse struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}

// HealthCheck go doc
// @Summary      Health of shipping service
// @Description  used to check whether shipping service is up and running or not
// @Tags         Health
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /    [get]
func HealthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Server is Up and running ! 🏃💨💨"})
	}
}
