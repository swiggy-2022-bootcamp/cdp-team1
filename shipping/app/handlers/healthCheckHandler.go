package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"qwik.in/shipping/domain/repository"
)

//HealthCheck ..
// func HealthCheck(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Shipping Server - Up and Running !",
// 	})
// }

//HealthCheckHandler ..
type HealthCheckHandler struct {
	shippingAddressRepository repository.ShippingAddressRepository
}

//HealthCheckResponse ..
type HealthCheckResponse struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}

//NewHealthCheckHandler ..
func NewHealthCheckHandler(shippingAddressRepo repository.ShippingAddressRepository) HealthCheckHandler {
	return HealthCheckHandler{shippingAddressRepository: shippingAddressRepo}
}

//HealthCheck ..
func (hc HealthCheckHandler) HealthCheck(c *gin.Context) {
	if hc.shippingAddressRepository.DBHealthCheck() {
		response := &HealthCheckResponse{
			Server:   "Server is Up and running ! 🏃🏼💨",
			Database: "Database is Up and running ! 🏃🏼💨",
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := &HealthCheckResponse{
			Server:   "Server is Up and running ! 🏃🏼💨",
			Database: "Database is down ! 🛏️💤",
		}
		c.JSON(http.StatusBadRequest, response)
	}
}
