package app

import (
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"qwik.in/rewards/app/handlers"
	"qwik.in/rewards/app/routes"
	"qwik.in/rewards/config"
	"qwik.in/rewards/log"
	"qwik.in/rewards/repository"
	"qwik.in/rewards/service"
)

func Start() {

	RewardRepository := repository.NewDynamoRepository()
	err := RewardRepository.Connect()
	if err != nil {
		return
	}
	fmt.Println("Connection successful")
	rewardService := service.NewRewardService(RewardRepository)
	fmt.Println("Service created")
	rewardController := handlers.NewRewardHandler(rewardService)
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

	routes.InitRoutes(router, rewardController)

	log.Info(config.SERVICE_NAME, " starting on port ", config.SERVER_PORT)
	err = router.Run(":" + config.SERVER_PORT)

	if err != nil {
		log.Error(config.SERVICE_NAME, " startup failed")
		return
	}
}
