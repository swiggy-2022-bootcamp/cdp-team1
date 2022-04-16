package db

import (
	"cartService/log"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *dynamodb.DynamoDB {
	//initialize client
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String(EnvRegion()),
		Credentials: credentials.NewStaticCredentials(EnvAccessKey(), EnvSecretKey(), ""),
	})
	client := dynamodb.New(sess)

	//ping the database
	_, err := client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Connection to dynamoDB failed.")
	}
	fmt.Println("Connected to DynamoDB")
	return client
}

func CreateTable(DB *dynamodb.DynamoDB) error {

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("user_id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("user_id"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("paymentCollection"),
	}
	response, err := DB.CreateTable(input)
	if err != nil {
		log.Error("Got error calling CreateTable: %s", err)
		return err
	}

	log.Info("Created the table" + response.String())
	return nil
}

func NewDbClient() *mongo.Client {

	clientOptions := options.Client().ApplyURI("")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		// logger.Fatal("Error while connecting to DB : " + err.Error())
		panic(err)
	} else {
		// logger.Info("Database Connected Successfully......")

	}
	return client
}
