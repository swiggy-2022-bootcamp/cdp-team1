package app

import (
	"cartService/app/handlers"
	"cartService/app/routes"
	"cartService/db"
	"cartService/domain/repository"
	"cartService/domain/service"
	"cartService/log"
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
)

var (
	server             *gin.Engine
	cartRepository     repository.CartRepositoryDB
	cartService        service.CartService
	cartHandler        handlers.CartHandler
	cartRoutes         routes.CartRoutes
	ctx                context.Context
	cartDB             *dynamodb.DynamoDB
	healthCheckHandler handlers.HealthCheckHandler
)

func Start() {

	ctx = context.TODO()

	//Variable initializations for DynamoDB
	cartDB = db.ConnectDB()
	db.CreateTable(cartDB)

	//Variable initializations to be used as dependency injectors
	cartRepository = repository.NewCartRepository(cartDB, ctx)
	cartService = service.NewCartService(cartRepository)
	cartHandler = handlers.NewCartHandler(cartService)
	healthCheckHandler = handlers.NewHealthCheckHandler(cartRepository)
	cartRoutes = routes.NewCartRoutes(cartHandler, healthCheckHandler)

	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configuring gin server and router
	server = gin.New()
	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("cart/api")
	cartRoutes.InitRoutes(router)

	//Starting server on port 9000
	err = server.Run(":9000")
	if err != nil {
		log.Error(err.Error() + " - Failed to start server")
	} else {
		log.Info("Server started successfully.")
	}
}
