package database

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"qwik.in/checkout/domain/tools/env"
	"qwik.in/checkout/domain/tools/logger"
)

//OrderConnectDB ..
func OrderConnectDB() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(env.GetRegion()),
		Credentials: credentials.NewStaticCredentials(env.GetAccessKey(), env.GetSecretKey(), ""),
	})
	orderClient := dynamodb.New(sess)

	//If cartCollectionTable is missing. Creates the table and stores the data.
	response, err1 := orderClient.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("team-1-orderCollection"),
	})
	if err1 != nil {
		logger.Error("Got err calling Order - CreateTable: %s", err1)
	}
	logger.Info("Created the table" + response.String())
	//*******************************

	//Lists the Table Input.
	result, err := orderClient.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.Error("Connection to Order - DynamoDB - FAILED ! ")
	} else {
		fmt.Println("Connection to Order - DynamoDB - SUCCESSFUL ! ")
		fmt.Println(result)
	}
	return orderClient
}
