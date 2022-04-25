package app

import (
	"context"
	"io"
	"orderService/app/handlers"
	"orderService/app/routes"
	"orderService/db"
	"orderService/domain/repository"
	"orderService/domain/service"
	"orderService/log"
	"orderService/protos"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

var (
	server             *gin.Engine
	orderRepository    repository.OrderRepositoryDB
	orderService       service.OrderService
	orderHandler       handlers.OrderHandler
	orderRoutes        routes.OrderRoutes
	ctx                context.Context
	orderDB            *dynamodb.DynamoDB
	healthCheckHandler handlers.HealthCheckHandler
)

func Start() {

	ctx = context.TODO()

	//Variable initializations for DynamoDB
	orderDB = db.ConnectDB()
	db.CreateTable(orderDB)

	//Variable initializations to be used as dependency injectors
	orderRepository = repository.NewOrderRepository(orderDB, ctx)
	orderService = service.NewOrderService(orderRepository)
	orderHandler = handlers.NewOrderHandler(orderService)
	healthCheckHandler = handlers.NewHealthCheckHandler(orderRepository)
	orderRoutes = routes.NewOrderRoutes(orderHandler, healthCheckHandler)

	obj := service.NewOrderProtoService(orderRepository)
	obj.CreateOrder(ctx, &protos.CreateOrderRequest{CustomerId: "1"})

	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configuring gin server and router
	server = gin.New()
	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("order/api")
	orderRoutes.InitRoutes(router)

	//Starting server on port 7000
	err = server.Run(":7000")
	if err != nil {
		log.Error(err.Error() + " - Failed to start server")
	} else {
		log.Info("Server started successfully.")
	}

}
