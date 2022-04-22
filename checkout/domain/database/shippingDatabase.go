package database

import (
	"fmt"
	"qwik.in/checkout/domain/tools/env"
	"qwik.in/checkout/domain/tools/logger"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//ShippingConnectDB DynamoDB
func ShippingConnectDB() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(env.GetRegion()),
		Credentials: credentials.NewStaticCredentials(env.GetAccessKey(), env.GetSecretKey(), ""),
	})
	shippingClient := dynamodb.New(sess)

	//*******************************
	//Creates DynamoDB Table.
	response, err1 := shippingClient.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("shipping_address_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("shipping_address_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
		TableName: aws.String("team-1-shipping"),
	})
	if err1 != nil {
		logger.Error("Got err calling Shipping CreateTable: %s", err1)
	}
	logger.Info("Created the table" + response.String())
	//*******************************

	result, err := shippingClient.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.Error("Connection to Shipping Address - DynamoDB - FAILED ! ")
	} else {
		fmt.Println("Connection to Shipping Address - DynamoDB - SUCCESSFUL ! ")
		fmt.Println(result)
	}
	return shippingClient
}
