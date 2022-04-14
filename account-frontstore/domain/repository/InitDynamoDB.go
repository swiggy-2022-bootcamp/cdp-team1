package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dbInitialized = false
var svc *dynamodb.DynamoDB

func GetDynamoDBInstance() *dynamodb.DynamoDB {
	if !dbInitialized {
		config := &aws.Config{
			Region:   aws.String("us-west-2"),
			Endpoint: aws.String("dynamodb.us-west-2.amazonaws.com"),
		}

		sess := session.Must(session.NewSession(config))

		svc = dynamodb.New(sess)
		dbInitialized = true
	}
	return svc
}
