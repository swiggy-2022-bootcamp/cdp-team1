package app

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"qwik.in/checkout/app/routes"
	"qwik.in/checkout/log"
)

//Start ..
func Start() {
	file, err := os.OpenFile("checkout-server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}
	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	routes.Router(router)
	router.Run(":9002")
}
