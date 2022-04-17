package app

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/productsFrontStore/app/handlers"
	"qwik.in/productsFrontStore/app/routes"
	"qwik.in/productsFrontStore/config"
	"qwik.in/productsFrontStore/log"
	"qwik.in/productsFrontStore/repository"
	"qwik.in/productsFrontStore/service"
)

func Start() {

	productRepository := repository.NewDynamoRepository()
	err := productRepository.Connect()
	if err != nil {
		log.Error("Error while Connecting to DB: ", err)
		return
	}
	productService := service.NewProductService(productRepository)
	productHandler := handlers.NewProductHandler(productService)

	file, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil && config.LOG_FILE_MODE {
		log.Info("Opened log file successfully")
		gin.DefaultWriter = io.MultiWriter(file)
	} else {
		log.Warn("Could not open log file, switching to IO mode")
	}

	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())

	routes.InitRoutes(router, productHandler)

	log.Info(config.SERVICE_NAME, " starting on port ", config.SERVER_PORT)
	err = router.Run(":" + config.SERVER_PORT)

	if err != nil {
		log.Error(config.SERVICE_NAME, " startup failed")
		return
	}
}
