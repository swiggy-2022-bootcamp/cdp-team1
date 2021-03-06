package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"qwik.in/customers-admin/env"
	"qwik.in/customers-admin/log"
)

var dbInitialized = false
var svc *dynamodb.DynamoDB

func GetDynamoDBInstance() *dynamodb.DynamoDB {
	if !dbInitialized {
		config := &aws.Config{
			//Region: aws.String("us-west-2"),
			//Endpoint:    aws.String("dynamodb.us-west-2.amazonaws.com"),
			Region:      aws.String(env.GetRegion()),
			Credentials: credentials.NewStaticCredentials(env.GetAccessKey(), env.GetSecretKey(), ""),
		}

		sess := session.Must(session.NewSession(config))

		svc = dynamodb.New(sess)
		dbInitialized = true
	}

	//ping the database
	_, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println(err)
		log.Error("Connection to dynamoDB failed.")
		fmt.Println("Connection to dynamoDB failed.")
		return nil
	}
	log.Info("Connected to dynamoDB")
	fmt.Println("Connected to dynamoDB")
	return svc
}
