package app

import (
	"context"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/shipping/app/handlers"
	"qwik.in/shipping/app/routes"
	"qwik.in/shipping/internal/db"
	repository2 "qwik.in/shipping/internal/repository"
	"qwik.in/shipping/internal/tools/log"
)

var (
	server             *gin.Engine
	shippingAddrRepo   repository2.ShippingAddrRepo
	shippingAddrRoutes routes.ShippingAddrRoutes
	ctx                context.Context
	shippingAddressDB  *dynamodb.DynamoDB
	healthCheckHandler handlers.HealthCheckHandler
)

//Start ..
func Start() {

	ctx = context.TODO()

	//DynamoDB
	shippingAddressDB = db.ConnectDB()
	err := db.CreateTable(shippingAddressDB)

	//Dependency Injectors
	shippingAddrRepo = repository2.ShippingAddrRepoImplFunc(ctx, shippingAddressDB)
	healthCheckHandler = handlers.HealthCheckHandlerFunc(shippingAddrRepo)
	shippingAddrRoutes = routes.ShippingAddrRoutesFunc(healthCheckHandler)

	//Custom Logger - Logs actions to 'shippingAddressService.log' file
	file, err := os.OpenFile("./pkg/logs/shippingAddressServiceLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configure - ShippingAddress Server and Router
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.SetTrustedProxies(nil)

	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("/shipping/api")
	shippingAddrRoutes.InitRoutes(router)

	//Kickstart - ShippingAddress Server on PORT 9003
	err = server.Run(":9003")
	if err != nil {
		log.Error(err.Error() + " - Server Failed to Start Up ! 🛌💤💤")
	} else {
		log.Info("Server Up and Running Successfully ! 🏃💨💨")
	}
}
