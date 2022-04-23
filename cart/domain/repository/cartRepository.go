package repository

import (
	"cartService/domain/model"
	"cartService/internal/error"
	"cartService/log"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type CartRepositoryDB interface {
	Create(*model.Cart) *error.AppError
	ReadAll() (*[]model.Cart, *error.AppError)
	Update(string, string) *error.AppError
	Delete(string) *error.AppError
	DeleteAll() *error.AppError
	DBHealthCheck() bool
}

const cartCollection = "cartCollection"

type CartRepository struct {
	cartDB *dynamodb.DynamoDB
	ctx    context.Context
}

func NewCartRepository(cartDB *dynamodb.DynamoDB, ctx context.Context) CartRepositoryDB {
	return &CartRepository{
		cartDB: cartDB,
		ctx:    ctx,
	}
}

func (cr CartRepository) DBHealthCheck() bool {

	_, err := cr.cartDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Database connection is down.")
		return false
	}
	return true
}

func (cdb CartRepository) Create(cart *model.Cart) *error.AppError {

	data, err := dynamodbattribute.MarshalMap(cart)
	if err != nil {
		log.Error("Marshalling of cart failed - " + err.Error())
		return error.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(cartCollection),
	}

	_, err = cdb.cartDB.PutItem(query)
	if err != nil {
		log.Error("Failed to insert item into database - " + err.Error())
		return error.NewUnexpectedError(err.Error())
	}

	return nil
}

func (cdb CartRepository) ReadAll() (*[]model.Cart, *error.AppError) {

	cart := &[]model.Cart{}

	query := &dynamodb.ScanInput{
		TableName: aws.String(cartCollection),
	}

	result, err := cdb.cartDB.Scan(query)
	if err != nil {
		log.Info(result)
		log.Error("Failed to get item from database - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	if result.Items == nil {
		log.Error("Cart for user doesn't exist. - ")
		notFoundError := error.NewNotFoundError("Payment mode for user doesn't exists")
		return nil, notFoundError
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, cart)
	if err != nil {
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	return cart, nil
}

func (cdb CartRepository) Update(customer_id string, updated_quantity string) *error.AppError {

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"quantity": {
				N: aws.String(updated_quantity),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"CustomerId": {
				S: aws.String(customer_id),
			},
		},
		TableName:        aws.String(cartCollection),
		UpdateExpression: aws.String("set quantity = :quantity"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	_, err := cdb.cartDB.UpdateItem(input)
	if err != nil {
		log.Error(err)
		return error.NewUnexpectedError(err.Error())
	}

	return nil
}

func (cdb CartRepository) Delete(customer_id string) *error.AppError {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(customer_id),
			},
		},
		TableName: aws.String(cartCollection),
	}

	_, err := cdb.cartDB.DeleteItem(input)
	if err != nil {
		log.Error(err)
		return error.NewUnexpectedError(err.Error())
	}

	return nil
}

func (cdb CartRepository) DeleteAll() *error.AppError {

	input := &dynamodb.ScanInput{
		TableName: aws.String(cartCollection),
	}

	result, err := cdb.cartDB.Scan(input)
	if err != nil {
		log.Error(err)
		return error.NewUnexpectedError(err.Error())
	}

	if result.Items == nil {
		log.Error("Cart for user doesn't exist. - ")
		notFoundError := error.NewNotFoundError("Payment mode for user doesn't exists")
		return notFoundError
	}

	for _, item := range result.Items {
		input := &dynamodb.DeleteItemInput{
			Key: map[string]*dynamodb.AttributeValue{
				"CustomerId": {
					S: aws.String(*item["customer_id"].S),
				},
			},
			TableName: aws.String(cartCollection),
		}

		_, err := cdb.cartDB.DeleteItem(input)
		if err != nil {
			log.Error(err)
			return error.NewUnexpectedError(err.Error())
		}
	}

	return nil
}
