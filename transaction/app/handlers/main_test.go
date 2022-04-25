package handlers

import (
	"github.com/gin-gonic/gin"
	"os"
	"qwik.in/transaction/mocks"
	"testing"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

func NewServer(transactionService *mocks.MockTransactionService) *gin.Engine {

	transactionHandler := NewTransactionHandler(transactionService)

	server := gin.Default()
	router := server.Group("api")

	router.POST("/transaction/:userId", transactionHandler.AddTransactionPoints)
	router.GET("/transaction/:userId", transactionHandler.GetTransactionPointsByUserID)
	router.POST("/transaction/use-transaction-points/:userId", transactionHandler.UseTransactionPoints)

	return server
}
