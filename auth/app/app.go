package app

import (
	"authService/app/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Start() {

	authRouter := gin.Default()
	routes.RegisterAuthRoutes(authRouter)

	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
		return
	}

	PORT := os.Getenv("PORT")

	err = authRouter.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Starting auth server on %s:%s ...\n", "127.0.0.1", PORT)
}
