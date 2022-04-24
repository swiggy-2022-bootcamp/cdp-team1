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

//CartConnectDB ..
func CartConnectDB() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(env.GetRegion()),
		Credentials: credentials.NewStaticCredentials(env.GetAccessKey(), env.GetSecretKey(), ""),
	})
	cartClient := dynamodb.New(sess)

	//If cartCollectionTable is missing. Creates the table and stores the data.
	response, err1 := cartClient.CreateTable(&dynamodb.CreateTableInput{
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
		TableName: aws.String("team-1-cartCollection"),
	})
	if err1 != nil {
		logger.Error("Got err calling Cart - CreateTable: %s", err1)
	}
	logger.Info("Created the table" + response.String())
	//*******************************

	//Lists the Table Input.
	_, err := cartClient.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.Error("Connection to Cart - DynamoDB - FAILED ! ")
	} else {
		fmt.Println("Connection to Cart - DynamoDB - SUCCESSFUL ! ")
		//fmt.Println(result)
	}
	return cartClient
}
