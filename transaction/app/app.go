package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/transaction/app/handlers"
	"qwik.in/transaction/app/routes"
	"qwik.in/transaction/config"
	"qwik.in/transaction/domain/repository"
	"qwik.in/transaction/domain/services"
	"qwik.in/transaction/log"
)

var (
	server                *gin.Engine
	transactionRepository repository.TransactionRepository
	transactionService    services.TransactionService
	transactionHandler    handlers.TransactionHandler
	transactionRoutes     routes.TransactionRoutes
	transactionDB         *dynamodb.DynamoDB
	healthCheckHandler    handlers.HealthCheckHandler
)

func Start() {
	//Variable initializations for DynamoDB
	transactionDB = config.ConnectDB()
	config.CreateTable(transactionDB)

	//Variable initializations to be used as dependency injections
	transactionRepository = repository.NewTransactionRepositoryImpl(transactionDB)
	transactionService = services.NewTransactionServiceImpl(transactionRepository)
	transactionHandler = handlers.NewTransactionHandler(transactionService)
	healthCheckHandler = handlers.NewHealthCheckHandler(transactionRepository)
	transactionRoutes = routes.NewTransactionRoutes(transactionHandler, healthCheckHandler)

	//Opening file for log collection
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	server = gin.New()
	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("transaction/api")
	transactionRoutes.InitRoutes(router)

	//Starting server on port 9001
	err = server.Run(":9001")
	if err != nil {
		log.Error(err.Error() + " - Failed to start server")
	} else {
		log.Info("Server started successfully.")
	}
}
