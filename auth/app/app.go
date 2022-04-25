package app

import (
	"authService/app/routes"
	"authService/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Start() {

	if fErr := config.LoadEnvConfig(); fErr != nil {
		fmt.Println("Start(): Error in loading config")
		log.Fatal(fErr.Message)
	}

	authRouter := gin.Default()
	routes.RegisterAuthRoutes(authRouter)

	PORT := config.EnvVars.GinPort

	err := authRouter.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Starting auth server on %s:%s ...\n", "127.0.0.1", PORT)
}
