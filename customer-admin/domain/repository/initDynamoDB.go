package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	config2 "qwik.in/customers-admin/config"
	"qwik.in/customers-admin/log"
)

var dbInitialized = false
var svc *dynamodb.DynamoDB

func GetDynamoDBInstance() *dynamodb.DynamoDB {
	if !dbInitialized {
		config := &aws.Config{
			Region: aws.String("us-west-2"),
			//Endpoint:    aws.String("dynamodb.us-west-2.amazonaws.com"),
			Credentials: credentials.NewStaticCredentials(config2.AccessKeyID, config2.SecretAccessKey, ""),
		}

		sess := session.Must(session.NewSession(config))

		svc = dynamodb.New(sess)
		dbInitialized = true
	}

	//ping the database
	_, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Connection to dynamoDB failed.")
		fmt.Println("Connection to dynamoDB failed.")
		return nil
	}
	log.Info("Connected to dynamoDB")
	fmt.Println("Connected to dynamoDB")
	return svc
}
