package database

import (
	"fmt"
	"github.com/ashwin2125/qwk/shipping/domain/tools/env"
	"github.com/ashwin2125/qwk/shipping/domain/tools/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//Connect DynamoDB
func Connect() *dynamodb.DynamoDB {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(env.GetRegion()),
		Credentials: credentials.NewStaticCredentials(env.GetAccessKey(), env.GetSecretKey(), ""),
	})
	client := dynamodb.New(sess)

	//*******************************
	response, err1 := client.CreateTable(&dynamodb.CreateTableInput{
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
		logger.Error("Got err calling CreateTable: %s", err1)
	}
	logger.Info("Created the table" + response.String())
	//*******************************

	_, err := client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		logger.Error("Connection to DynamoDB failed. 🛌💤💤")
	}
	fmt.Println("Connected to DynamoDB ! 🏃💨💨")
	return client

}
