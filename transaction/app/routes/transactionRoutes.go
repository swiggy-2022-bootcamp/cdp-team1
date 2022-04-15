package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/transaction/app/handlers"
	_ "qwik.in/transaction/docs"
)

type TransactionRoutes struct {
	transactionHandler handlers.TransactionHandler
	healthCheckhandler handlers.HealthCheckHandler
}

func NewTransactionRoutes(transactionHandler handlers.TransactionHandler, healthCheckhandler handlers.HealthCheckHandler) TransactionRoutes {
	return TransactionRoutes{transactionHandler: transactionHandler, healthCheckhandler: healthCheckhandler}
}

func (tr TransactionRoutes) InitRoutes(newRouter *gin.RouterGroup) {

	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", tr.healthCheckhandler.HealthCheck)

	newRouter.POST("/transaction/:userId", tr.transactionHandler.AddTransactionPoints)
	newRouter.GET("/transaction/:userId", tr.transactionHandler.GetTransactionPointsByUserID)
	newRouter.POST("/transaction/use-transaction-points/:userId", tr.transactionHandler.UseTransactionPoints)
}
