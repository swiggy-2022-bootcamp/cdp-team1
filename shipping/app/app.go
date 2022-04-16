package app

import (
	"context"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"qwik.in/shipping/app/handlers"
	"qwik.in/shipping/app/routes"
	db "qwik.in/shipping/config/db"
	"qwik.in/shipping/domain/repository"
	"qwik.in/shipping/internal/log"
)

var (
	shippingAddressServer     *gin.Engine
	shippingAddressRepository repository.ShippingAddressRepository
	shippingAddressRoutes     routes.ShippingAddressRoutes
	ctx                       context.Context
	shippingAddressDB         *dynamodb.DynamoDB
	healthCheckHandler        handlers.HealthCheckHandler
)

//Start ..
func Start() {

	ctx = context.TODO()

	//DynamoDB
	shippingAddressDB = db.ConnectDB()
	db.CreateTable(shippingAddressDB)

	//Dependency Injectors
	shippingAddressRepository = repository.NewShippingAddressRepositoryImplementation(shippingAddressDB, ctx)
	healthCheckHandler = handlers.NewHealthCheckHandler(shippingAddressRepository)
	shippingAddressRoutes = routes.NewShippingAddressRoutes(healthCheckHandler)

	//Custom Logger - Logs actions to 'shippingAddressService.log' file
	file, err := os.OpenFile("shippingAddressService.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configure - ShippingAddress Server and Router
	shippingAddressServer := gin.New()
	shippingAddressServer.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	shippingAddressRouter := shippingAddressServer.Group("/shipping/api")
	shippingAddressRoutes.InitRoutes(shippingAddressRouter)

	//Kickstart - ShippingAddress Server on PORT 9003
	err = shippingAddressServer.Run(":9003")
	if err != nil {
		log.Error(err.Error() + " - Server Failed to Start Up ! 🛏️💤")
	} else {
		log.Info("Server Up and Running Successfully ! 🏃🏼💨")
	}
}
