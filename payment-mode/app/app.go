package app

import (
	"context"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/payment-mode/app/handlers"
	"qwik.in/payment-mode/app/routes"
	"qwik.in/payment-mode/config"
	"qwik.in/payment-mode/domain/repository"
	"qwik.in/payment-mode/domain/services"
	"qwik.in/payment-mode/log"
)

var (
	server             *gin.Engine
	paymentRepository  repository.PaymentRepository
	paymentService     services.PaymentService
	paymentHandler     handlers.PaymentHandler
	paymentRoutes      routes.PaymentRoutes
	ctx                context.Context
	paymentDB          *dynamodb.DynamoDB
	healthCheckHandler handlers.HealthCheckHandler
)

func Start() {
	ctx = context.TODO()

	//Variable initializations for DynamoDB
	paymentDB = config.ConnectDB()
	config.CreateTable(paymentDB)

	//Variable initializations to be used as dependency injectors
	paymentRepository = repository.NewPaymentRepositoryImpl(paymentDB, ctx)
	paymentService = services.NewPaymentServiceImpl(paymentRepository)
	paymentHandler = handlers.NewPaymentHandler(paymentService)
	healthCheckHandler = handlers.NewHealthCheckHandler(paymentRepository)
	paymentRoutes = routes.NewPaymentRoutes(paymentHandler, healthCheckHandler)

	//Opening file for log collection
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configuring gin server and router
	server = gin.New()
	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("payment-mode/api")
	paymentRoutes.InitRoutes(router)

	//Starting server on port 9000
	err = server.Run(":9000")
	if err != nil {
		log.Error(err.Error() + " - Failed to start server")
	} else {
		log.Info("Server started successfully.")
	}
}
