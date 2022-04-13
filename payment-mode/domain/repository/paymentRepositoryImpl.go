package repository

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	apperrors "qwik.in/payment-mode/app-errors"
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

func (p PaymentRepositoryImpl) AddPaymentModeToDB(userPaymentMode *models.UserPaymentMode) *apperrors.AppError {

	data, err := dynamodbattribute.MarshalMap(userPaymentMode)
	if err != nil {
		log.Error("Marshalling of userPaymentMode failed - " + err.Error())
		return apperrors.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(paymentCollection),
	}

	_, err = p.paymentDB.PutItem(query)
	if err != nil {
		log.Error("Failed to insert item into database - " + err.Error())
		return apperrors.NewUnexpectedError(err.Error())
	}
	return nil
}

func (p PaymentRepositoryImpl) GetPaymentModeFromDB(userId string) (*models.UserPaymentMode, *apperrors.AppError) {
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
		return nil, apperrors.NewUnexpectedError(err.Error())
	}

	if result.Item == nil {
		log.Error("Payment mode for user doesn't exists. - ")
		err_ := apperrors.NewNotFoundError("Payment mode for user doesn't exists")
		return nil, err_
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, userPaymentMode)
	if err != nil {
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, apperrors.NewUnexpectedError(err.Error())
	}
	return userPaymentMode, nil
}

func (p PaymentRepositoryImpl) UpdatePaymentModeToDB(userPaymentMode *models.UserPaymentMode) *apperrors.AppError {
	type updateQuery struct {
		PaymentMode []models.PaymentMode `json:":payment_mode"`
	}

	update, err := dynamodbattribute.MarshalMap(updateQuery{
		PaymentMode: userPaymentMode.PaymentModes,
	})
	if err != nil {
		log.Error(err)
		return apperrors.NewUnexpectedError(err.Error())
	}

	input := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userPaymentMode.UserId),
			},
		},
		TableName:                 aws.String(paymentCollection),
		UpdateExpression:          aws.String("set payment_modes = :payment_mode"),
		ExpressionAttributeValues: update,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err = p.paymentDB.UpdateItem(input)
	if err != nil {
		log.Error(err)
		return apperrors.NewUnexpectedError(err.Error())
	}

	return nil
}
