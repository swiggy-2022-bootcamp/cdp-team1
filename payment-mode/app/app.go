package app

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/payment-mode/app/routes"
	"qwik.in/payment-mode/services"
)

func Start() {
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	router := gin.New()
	router.Use(services.UseLogger(services.DefaultLoggerFormatter), gin.Recovery())

	routes.InitRoutes(router)

	router.Run(":9000")
}
