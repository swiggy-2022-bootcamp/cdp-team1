package app

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
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
	server            *gin.Engine
	paymentRepository repository.PaymentRepository
	paymentService    services.PaymentService
	paymentHandler    handlers.PaymentHandler
	paymentRoutes     routes.PaymentRoutes
	ctx               context.Context
	mongoCollection   *mongo.Collection
	mongoClient       *mongo.Client
	err               error
)

func Start() {
	ctx = context.TODO()

	//Variable initializations for DB
	mongoClient = config.ConnectDB()
	mongoCollection = config.GetCollection(mongoClient, "paymentMode")

	//Variable initializations to be used as dependency injectors
	paymentRepository = repository.NewPaymentRepositoryImpl(mongoCollection, ctx)
	paymentService = services.NewPaymentServiceImpl(paymentRepository)
	paymentHandler = handlers.NewPaymentHandler(paymentService)
	paymentRoutes = routes.NewPaymentRoutes(paymentHandler)

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
