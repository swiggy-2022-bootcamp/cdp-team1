package db

import (
	"authService/domain"
	"authService/errs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"time"
)

const (
	custTableName string = "team-1-customers"
	custTableGSI  string = "email-index"
)

type customerRepositoryDB struct {
	timeout time.Duration
	client  *dynamodb.DynamoDB
}

func NewCustomerRepositoryDB(dbClient *dynamodb.DynamoDB, timeout time.Duration) domain.UserRepositoryDB {
	return &customerRepositoryDB{
		timeout: timeout,
		client:  dbClient,
	}
}

func (cdb customerRepositoryDB) RepoHealthCheck() *errs.AppError {

	err := PingDatabase(cdb.client)
	return err
}

func (cdb customerRepositoryDB) FetchByID(userID string) (*domain.User, *errs.AppError) {

	input := &dynamodb.GetItemInput{
		TableName: aws.String(custTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {S: aws.String(userID)},
		},
	}

	res, err := cdb.client.GetItem(input)
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

func (cdb customerRepositoryDB) FetchByUsername(username string) (*domain.User, *errs.AppError) {

	return nil, errs.NewUnexpectedError("customer has no username")
}

func (cdb customerRepositoryDB) FetchByEmail(email string) (*domain.User, *errs.AppError) {

	input := &dynamodb.QueryInput{
		TableName: aws.String(custTableName),
		IndexName: aws.String(custTableGSI),
		ExpressionAttributeNames: map[string]*string{
			"#email": aws.String("email"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":email": {S: aws.String(email)},
		},
		KeyConditionExpression: aws.String("#email = :email"),
	}
	res, err := cdb.client.Query(input)
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
