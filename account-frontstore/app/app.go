package app

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/account-frontstore/app/routes"
	"qwik.in/account-frontstore/log"
)

func Start() {
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}
	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	routes.Router(router)
	router.Run(":7001")
}
