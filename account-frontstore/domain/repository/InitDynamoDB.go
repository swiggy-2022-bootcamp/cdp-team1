package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	config2 "qwik.in/account-frontstore/config"
)

var dbInitialized = false
var svc *dynamodb.DynamoDB

func GetDynamoDBInstance() *dynamodb.DynamoDB {
	if !dbInitialized {
		config := &aws.Config{
			Region:      aws.String("us-west-2"),
			Credentials: credentials.NewStaticCredentials(config2.AccessKeyID, config2.SecretAccessKey, ""),
		}

		sess := session.Must(session.NewSession(config))

		svc = dynamodb.New(sess)
		dbInitialized = true
	}

	//ping the database
	_, err := svc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Println("Connection to dynamoDB failed.")
		return nil
	}
	fmt.Println("Connected to DB")
	return svc
}
