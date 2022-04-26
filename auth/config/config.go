package config

import (
	"authService/errs"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type EnvConfig struct {
	GinPort            string
	GrpcPort           string
	SecretBytes        []byte
	TokenDuration      int
	AWSRegion          string
	AWSAccessKeyID     string
	AWSSecretAccessKey string
}

var EnvVars = &EnvConfig{}

func LoadEnvConfig() *errs.AppError {

	if err := godotenv.Load("./.env", "./local.env"); err != nil {
		return errs.NewUnexpectedError(err.Error())
	}

	EnvVars.GinPort = os.Getenv("GIN_PORT")
	EnvVars.GrpcPort = os.Getenv("GRPC_PORT")
	EnvVars.SecretBytes = []byte(os.Getenv("SECRET"))

	var err error
	if EnvVars.TokenDuration, err = strconv.Atoi(os.Getenv("TOKEN_EXP_TIME_IN_MINS")); err != nil {
		EnvVars.TokenDuration = 60
	}

	EnvVars.AWSRegion = os.Getenv("REGION")
	EnvVars.AWSAccessKeyID = os.Getenv("AWS_ACCESS_KEY_ID")
	EnvVars.AWSSecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	return nil
}
