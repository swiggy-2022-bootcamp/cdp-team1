package handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"qwik.in/payment-mode/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

// Creating test server with PaymentService mock.
func NewServer(paymentService *mocks.MockPaymentService) *gin.Engine {
	paymentHandler := NewPaymentHandler(paymentService)

	server := gin.Default()
	router := server.Group("/api")

	router.POST("/paymentmethods", paymentHandler.AddPaymentMode)
	router.GET("/paymentmethods", paymentHandler.GetPaymentMode)
	router.POST("/pay", paymentHandler.CompletePayment)

	return server
}
