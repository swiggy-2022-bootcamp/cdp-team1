package handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"qwik.in/payment-mode/domain/services"
	"qwik.in/payment-mode/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewServer(paymentRepository *mocks.MockPaymentRepository) *gin.Engine {
	healthCheckHandler := NewHealthCheckHandler(paymentRepository)
	paymentService := services.NewPaymentServiceImpl(paymentRepository)
	paymentHandler := NewPaymentHandler(paymentService)

	server := gin.Default()
	router := server.Group("payment-mode/api")

	router.GET("/", healthCheckHandler.HealthCheck)
	router.POST("/paymentmethods/:userId", paymentHandler.AddPaymentMode)
	router.GET("/paymentmethods/:userId", paymentHandler.GetPaymentMode)

	return server
}
