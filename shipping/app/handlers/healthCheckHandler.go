package handlers

import (
	"net/http"

	"qwik.in/shipping/internal/repository"

	"github.com/gin-gonic/gin"
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

//HealthCheckHandlerFunc ..
func HealthCheckHandlerFunc(shippingAddrRepoObj repository.ShippingAddrRepo) HealthCheckHandler {
	return HealthCheckHandler{shippingAddrRepoObj: shippingAddrRepoObj}
}

// HealthCheck godoc
// @Summary      To check if the service is running or not.
// @Description  This request will return 200 OK, if server is up and running.
// @Tags         HealthCheckResponse
// @Schemes
// @Accept   json
// @Produce  json
// @Success  200  {object}  HealthCheckResponse
// @Failure  500  {object}  HealthCheckResponse
// @Router   / [GET]
func (hc HealthCheckHandler) HealthCheck(c *gin.Context) {
	if hc.shippingAddrRepoObj.DBHealthCheck() {
		response := &HealthCheckResponse{
			Server:   "Server is Up and running ! 🏃💨💨",
			Database: "Database is Up and running ! 🏃💨💨",
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := &HealthCheckResponse{
			Server:   "Server is Up and running ! 🏃💨💨",
			Database: "Database is down ! 🛌💤💤",
		}
		c.JSON(http.StatusBadRequest, response)
	}
}
