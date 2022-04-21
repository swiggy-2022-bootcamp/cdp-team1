package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/payment-mode/domain/repository"
)

type HealthCheckHandler struct {
	paymentRepository repository.PaymentRepository
}

// Response for health check API call
type HealthCheckResponse struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}

func NewHealthCheckHandler(paymentRepository repository.PaymentRepository) HealthCheckHandler {
	return HealthCheckHandler{paymentRepository: paymentRepository}
}

// HealthCheck godoc
// @Summary To check if the service is running or not.
// @Description This request will return 200 OK if server is up..
// @Tags HealthCheckResponse
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {object} 	HealthCheckResponse
// @Failure 500  {object} 	HealthCheckResponse
// @Router / [GET]
func (hc HealthCheckHandler) HealthCheck(c *gin.Context) {

	// Checks if database is active or not
	if hc.paymentRepository.DBHealthCheck() {
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
