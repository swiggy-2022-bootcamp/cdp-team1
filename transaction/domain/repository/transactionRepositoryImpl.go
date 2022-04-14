package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	apperros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
	"qwik.in/transaction/log"
	"strconv"
)

const transactionCollection = "transactionCollection_Team1"

type TransactionRepositoryImpl struct {
	transactionDB *dynamodb.DynamoDB
}

func NewTransactionRepositoryImpl(transactionDB *dynamodb.DynamoDB) TransactionRepository {
	return &TransactionRepositoryImpl{
		transactionDB: transactionDB,
	}
}

func (t TransactionRepositoryImpl) AddTransactionPointsFromDB(transaction *models.Transaction) *apperros.AppError {
	data, err := dynamodbattribute.MarshalMap(transaction)
	if err != nil {
		log.Error("Marshalling of transaction failed - " + err.Error())
		return apperros.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(transactionCollection),
	}

	_, err = t.transactionDB.PutItem(query)
	if err != nil {
		log.Error("Failed to insert item into database - " + err.Error())
		return apperros.NewUnexpectedError(err.Error())
	}
	return nil
}

func (t TransactionRepositoryImpl) GetTransactionPointsByUserIdFromDB(userId string) (int, *apperros.AppError) {
	transaction := &models.Transaction{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(transactionCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(userId),
			},
		},
	}

	result, err := t.transactionDB.GetItem(query)
	if err != nil {
		log.Error("Failed to get item from database - " + err.Error())
		return -1, apperros.NewUnexpectedError(err.Error())
	}

	if result.Item == nil {
		log.Error("User with given userId doesn't exists")
		return -1, apperros.NewNotFoundError("User with given userId doesn't exists")
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, transaction)
	if err != nil {
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return -1, apperros.NewUnexpectedError(err.Error())
	}
	return transaction.TransactionPoints, nil
}

func (t TransactionRepositoryImpl) DBHealthCheck() bool {
	_, err := t.transactionDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Database connection is down.")
		return false
	}
	return true
}

func (t TransactionRepositoryImpl) UpdateTransactionPointsToDB(transaction *models.Transaction) *apperros.AppError {
	query := &dynamodb.UpdateItemInput{
		TableName: aws.String(transactionCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"user_id": {
				S: aws.String(transaction.UserId),
			},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":points": {
				N: aws.String(strconv.Itoa(transaction.TransactionPoints)),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set transaction_points = :points"),
	}

	_, err := t.transactionDB.UpdateItem(query)
	if err != nil {
		log.Error("Failed to update transaction points" + err.Error())
		return apperros.NewUnexpectedError(err.Error())
	}

	return nil
}
