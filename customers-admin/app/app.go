package app

import (
	"io"
	"os"
	"qwik.in/customers-admin/app/log"
	"qwik.in/customers-admin/app/routes"

	"github.com/gin-gonic/gin"
)

func Start() {
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}
	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	routes.Router(router)
	router.Run(":7000")
}
