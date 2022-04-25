package db

import (
	"log"
	"time"

	"authService/domain"
	"authService/errs"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const authTableName string = "team-1-session"

type authRepositoryDB struct {
	timeout time.Duration
	client  *dynamodb.DynamoDB
}

func NewAuthRepositoryDB(dbClient *dynamodb.DynamoDB, timeout time.Duration) domain.AuthRepositoryDB {
	return &authRepositoryDB{
		timeout: timeout,
		client:  dbClient,
	}
}

func (adb authRepositoryDB) RepoHealthCheck() *errs.AppError {

	err := PingDatabase(adb.client)
	return err
}

func (adb authRepositoryDB) SaveSession(session domain.Session) *errs.AppError {

	item, err := dynamodbattribute.MarshalMap(NewDBSession(session))
	if err != nil {
		log.Println(err)
		return errs.NewUnexpectedError(err.Error())
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(authTableName),
		Item:      item,
		ExpressionAttributeNames: map[string]*string{
			"#uuid": aws.String("uuid"),
		},
		ConditionExpression: aws.String("attribute_not_exists(#uuid)"),
	}

	if _, err := adb.client.PutItem(input); err != nil {
		log.Println(err)

		if _, ok := err.(*dynamodb.ConditionalCheckFailedException); ok {
			return errs.NewUnexpectedError("conditional check failed")
		}

		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}

func (adb authRepositoryDB) FetchSessionByTokenSign(tokenSign string) (*domain.Session, *errs.AppError) {

	input := &dynamodb.GetItemInput{
		TableName: aws.String(authTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"token_signature": {S: aws.String(tokenSign)},
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

	var session Session
	if err := dynamodbattribute.UnmarshalMap(res.Item, &session); err != nil {
		log.Println(err)

		return nil, errs.NewUnexpectedError(err.Error())
	}

	return session.DomainSession(), nil
}

func (adb authRepositoryDB) UpdateSessionStatusByTokenSign(tokenSign string) *errs.AppError {

	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(authTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"token_signature": {S: aws.String(tokenSign)},
		},
		ExpressionAttributeNames: map[string]*string{
			"#is_active":  aws.String("is_active"),
			"#updated_at": aws.String("updated_at"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":is_active": {BOOL: aws.Bool(false)},
			":updated_at": {
				S: aws.String(time.Now().Format("2006-01-02T15:04:05.0000000-07:00")),
			},
		},
		UpdateExpression: aws.String("set #is_active = :is_active, #updated_at = :updated_at"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	if _, err := adb.client.UpdateItem(input); err != nil {
		log.Println(err)

		return errs.NewUnexpectedError(err.Error())
	}

	return nil
}
