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

func NewServer(paymentService *mocks.MockPaymentService) *gin.Engine {
	paymentHandler := NewPaymentHandler(paymentService)

	server := gin.Default()
	router := server.Group("payment-mode/api")

	router.POST("/paymentmethods/:userId", paymentHandler.AddPaymentMode)
	router.GET("/paymentmethods/:userId", paymentHandler.GetPaymentMode)
	router.POST("/setpaymentmethods/:userId", paymentHandler.SetPaymentMode)
	router.POST("/pay", paymentHandler.CompletePayment)

	return server
}
