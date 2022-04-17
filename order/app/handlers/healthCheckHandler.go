package handlers

import (
	"net/http"
	"orderService/domain/repository"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	orderRepository repository.OrderRepositoryDB
}

type HealthCheckResponse struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}

func NewHealthCheckHandler(orderRepository repository.OrderRepositoryDB) HealthCheckHandler {
	return HealthCheckHandler{orderRepository: orderRepository}
}

// HealthCheck godoc
// @Summary To check if the service is running or not.
// @Description This request will return 200 OK if server is up
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	Service is up
// @Router / [GET]
func (hc HealthCheckHandler) HealthCheck(c *gin.Context) {

	//Check to be added for database.
	if hc.orderRepository.DBHealthCheck() {
		response := &HealthCheckResponse{
			Server:   "Server is up",
			Database: "Database is up",
		}
		c.JSON(http.StatusOK, response)
	} else {
		response := &HealthCheckResponse{
			Server:   "Server is up",
			Database: "Database is down",
		}
		c.JSON(http.StatusInternalServerError, response)
	}
}
