package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func EnvAccessKey() string {
	loadEnvFile()
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

func EnvSecretKey() string {
	loadEnvFile()
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

func EnvRegion() string {
	loadEnvFile()
	return os.Getenv("REGION")
}

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file.")
	}
}
