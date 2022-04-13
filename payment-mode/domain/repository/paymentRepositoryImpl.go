package repository

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	app_errors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/log"
)

const paymentCollection = "paymentCollection"

type PaymentRepositoryImpl struct {
	paymentDB *dynamodb.DynamoDB
	ctx       context.Context
}

func NewPaymentRepositoryImpl(paymentDB *dynamodb.DynamoDB, ctx context.Context) PaymentRepository {
	return &PaymentRepositoryImpl{
		paymentDB: paymentDB,
		ctx:       ctx,
	}
}

func (p PaymentRepositoryImpl) DBHealthCheck() bool {

	_, err := p.paymentDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Database connection is down.")
		return false
	}
	return true
}

func (p PaymentRepositoryImpl) AddPaymentModeToDB(userPaymentMode *models.UserPaymentMode) error {

	data, err := dynamodbattribute.MarshalMap(userPaymentMode)
	if err != nil {
		log.Error("Marshalling of userPaymentMode failed - " + err.Error())
		return err
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(paymentCollection),
	}

	_, err = p.paymentDB.PutItem(query)
	if err != nil {
		log.Error("Failed to insert item into database - " + err.Error())
		return err
	}
	return nil
}

func (p PaymentRepositoryImpl) GetPaymentModeFromDB(userId string) (*models.UserPaymentMode, error) {
	userPaymentMode := &models.UserPaymentMode{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(paymentCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userId),
			},
		},
	}

	result, err := p.paymentDB.GetItem(query)
	if err != nil {
		log.Info(result)
		log.Error("Failed to get item from database - " + err.Error())
		return nil, err
	}

	if result.Item == nil {
		log.Error("Payment mode for user doesn't exists. - ")
		err_ := app_errors.NewBadRequestError("Payment mode for user doesn't exists")
		return nil, err_.Error()
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, userPaymentMode)
	if err != nil {
		log.Info(result.Item)
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, err
	}
	return userPaymentMode, nil
}
