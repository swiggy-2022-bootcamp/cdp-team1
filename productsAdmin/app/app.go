package app

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/productsAdmin/app/routes"
	"qwik.in/productsAdmin/config"
	"qwik.in/productsAdmin/log"
)

func Start() {
	file, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil && config.LOG_FILE_MODE {
		log.Info("Opened log file successfully")
		gin.DefaultWriter = io.MultiWriter(file)
	} else {
		log.Warn("Could not open log file, switching to IO mode")
	}

	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())

	routes.InitRoutes(router)

	log.Info("Service starting on port ", config.SERVER_PORT)
	err = router.Run(":" + config.SERVER_PORT)

	if err != nil {
		log.Error("Service startup failed")
		return
	}
}
