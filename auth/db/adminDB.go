package db

import (
	"authService/domain"
	"authService/errs"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"time"
)

const (
	adminTableName string = "team-1-admin-users"
	adminTableGSI  string = "username-index"
)

type adminRepositoryDB struct {
	timeout time.Duration
	client  *dynamodb.DynamoDB
}

func NewAdminRepositoryDB(dbClient *dynamodb.DynamoDB, timeout time.Duration) domain.UserRepositoryDB {
	return &adminRepositoryDB{
		timeout: timeout,
		client:  dbClient,
	}
}

func (adb adminRepositoryDB) RepoHealthCheck() *errs.AppError {

	if adb.client == nil {
		fmt.Println("(adb adminRepositoryDB) RepoHealthCheck()")
		fmt.Println(adb.client)
	}
	err := PingDatabase(adb.client)
	return err
}

func (adb adminRepositoryDB) FetchByID(userID string) (*domain.User, *errs.AppError) {

	input := &dynamodb.GetItemInput{
		TableName: aws.String(adminTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {S: aws.String(userID)},
		},
	}

	res, err := adb.client.GetItem(input)
	if err != nil {
		log.Println(err)
		return nil, errs.NewUnexpectedError(err.Error())
	}

	if res.Item == nil {
		return nil, nil
	}

	var user User
	if err := dynamodbattribute.UnmarshalMap(res.Item, &user); err != nil {
		log.Println(err)

		return nil, errs.NewUnexpectedError(err.Error())
	}

	return user.DomainUser(), nil
}

func (adb adminRepositoryDB) FetchByUsername(username string) (*domain.User, *errs.AppError) {

	input := &dynamodb.QueryInput{
		TableName: aws.String(adminTableName),
		IndexName: aws.String(adminTableGSI),
		ExpressionAttributeNames: map[string]*string{
			"#username": aws.String("username"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":username": {S: aws.String(username)},
		},
		KeyConditionExpression: aws.String("#username = :username"),
	}
	res, err := adb.client.Query(input)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}
	if *res.Count != 1 {
		return nil, errs.NewUnexpectedError("records mismatch in database")
	}

	var user User
	if err := dynamodbattribute.UnmarshalMap(res.Items[0], &user); err != nil {
		log.Println(err)

		return nil, errs.NewUnexpectedError(err.Error())
	}
	return user.DomainUser(), nil
}

func (adb adminRepositoryDB) FetchByEmail(email string) (*domain.User, *errs.AppError) {

	return nil, errs.NewUnexpectedError("admin has no email")
}
