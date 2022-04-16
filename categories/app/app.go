package app

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"qwik.in/categories/app/handlers"
	"qwik.in/categories/app/routes"
	"qwik.in/categories/config"
	"qwik.in/categories/log"
	"qwik.in/categories/repository"
	"qwik.in/categories/service"
)

func Start() {

	CategoryRepository := repository.NewDynamoRepository()
	err := CategoryRepository.Connect()
	if err != nil {
		return
	}
	fmt.Println("Connection successful")
	categoryService := service.NewCategoryService(CategoryRepository)
	fmt.Println("Service created")
	categoryController := handlers.NewCategoryHandler(categoryService)
	fmt.Println("Controller created")
	file, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil && config.LOG_FILE_MODE {
		log.Info("Opened log file successfully")
		gin.DefaultWriter = io.MultiWriter(file)
	} else {
		log.Warn("Could not open log file, switching to IO mode")
	}

	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())

	routes.InitRoutes(router, categoryController)

	log.Info(config.SERVICE_NAME, " starting on port ", config.SERVER_PORT)
	err = router.Run(":" + config.SERVER_PORT)

	if err != nil {
		log.Error(config.SERVICE_NAME, " startup failed")
		return
	}
}
