package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

//GetAccessKey - AWS_ACCESS_KEY_ID
func GetAccessKey() string {
	LoadEnvConfig()
	return os.Getenv("AWS_ACCESS_KEY_ID")
}

//GetSecretKey - AWS_SECRET_ACCESS_KEY
func GetSecretKey() string {
	LoadEnvConfig()
	return os.Getenv("AWS_SECRET_ACCESS_KEY")
}

//GetRegion ..
func GetRegion() string {
	LoadEnvConfig()
	return os.Getenv("REGION")
}

//LoadEnvConfig - Reads .env file using OS
func LoadEnvConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file.")
	}
}
