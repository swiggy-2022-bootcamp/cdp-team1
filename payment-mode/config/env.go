package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvMonogoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("MONGOURI")
}

func EnvJWTSecretKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}

	return os.Getenv("SECRET_KEY")
}
