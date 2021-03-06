package db

import (
	"authService/config"
	"authService/errs"
	"authService/log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func NewDynamoDBClient() *dynamodb.DynamoDB {

	var err error
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(config.EnvVars.AWSRegion),
		Credentials: credentials.NewStaticCredentials(config.EnvVars.AWSAccessKeyID, config.EnvVars.AWSSecretAccessKey, ""),
	})
	if err != nil {
		log.Error("Session creation in dynamoDB failed : " + err.Error())
		return nil
	}

	// Create DynamoDB client
	client := dynamodb.New(sess)

	//ping the database
	_, err = client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Connection to dynamoDB failed : " + err.Error())
		return nil
	}
	log.Info("Connected to DynamoDB")
	return client
}

func PingDatabase(client *dynamodb.DynamoDB) *errs.AppError {

	_, err := client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("error while pinging database: ", err)
		return errs.NewUnexpectedError(err.Error())
	}
	return nil
}
