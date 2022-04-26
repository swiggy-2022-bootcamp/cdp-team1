package handlers

import (
	mocks "cartService/mock"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

// Creating test server with PaymentService mock.
func NewServer(paymentService *mocks.MockCartService) *gin.Engine {

	cartHandler := NewCartHandler(paymentService)

	server := gin.Default()
	newRouter := server.Group("/api")

	newRouter.POST("/cart", cartHandler.CreateCart)
	newRouter.PUT("/cart", cartHandler.UpdateCart)
	newRouter.GET("/cart", cartHandler.GetCart)
	newRouter.DELETE("/cart/:id", cartHandler.DeleteCartItem)
	newRouter.DELETE("/cart/empty", cartHandler.DeleteCartAll)

	return server
}
