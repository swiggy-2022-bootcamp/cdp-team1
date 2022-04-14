package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/transaction/domain/repository"
)

type HealthCheckHandler struct {
	transactionRepository repository.TransactionRepository
}

func NewHealthCheckHandler(transactionRepository repository.TransactionRepository) HealthCheckHandler {
	return HealthCheckHandler{transactionRepository: transactionRepository}
}

type HealthCheckResponse struct {
	Server   string `json:"server"`
	Database string `json:"database"`
}

// HealthCheck godoc
// @Summary To check if the service is running or not.
// @Description This request will return 200 OK if server is up..
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	Service is up
// @Router / [GET]
func (hc HealthCheckHandler) HealthCheck(c *gin.Context) {

	//Checking for database status.
	if hc.transactionRepository.DBHealthCheck() {
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
