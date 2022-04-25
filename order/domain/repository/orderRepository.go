package repository

import (
	"context"
	"fmt"
	"math/rand"
	"orderService/domain/model"
	"orderService/internal/error"
	"orderService/log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type OrderRepositoryDB interface {
	Create(*model.Order) *error.AppError
	ReadStatus(string) (*[]model.Order, *error.AppError)
	ReadOrderID(string) (*model.Order, *error.AppError)
	ReadCustomerID(string) (*[]model.Order, *error.AppError)
	ReadAll() (*[]model.Order, *error.AppError)
	Update(string, string) *error.AppError
	Delete(model.Order) *error.AppError
	CreateOrderInvoice(string) *error.AppError
	DBHealthCheck() bool
	// DeleteAll() *error.AppError
}

const orderCollection = "orderCollection"

type OrderRepository struct {
	orderDB *dynamodb.DynamoDB
	ctx     context.Context
}

func NewOrderRepository(orderDB *dynamodb.DynamoDB, ctx context.Context) OrderRepositoryDB {
	return &OrderRepository{
		orderDB: orderDB,
		ctx:     ctx,
	}
}

func (or OrderRepository) DBHealthCheck() bool {

	_, err := or.orderDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Database connection is down.")
		return false
	}
	return true
}

func (odb OrderRepository) Create(order *model.Order) *error.AppError {

	order.Datetime = time.Now().String()
	order.OrderId = strconv.Itoa(rand.Intn(1000))

	data, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		log.Error("Marshalling of order failed - " + err.Error())
		return error.NewUnexpectedError(err.Error())
	}

	query := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(orderCollection),
	}

	_, err = odb.orderDB.PutItem(query)
	if err != nil {
		log.Error("Failed to insert item into database - " + err.Error())
		return error.NewUnexpectedError(err.Error())
	}

	return nil
}

func (odb OrderRepository) ReadStatus(status string) (*[]model.Order, *error.AppError) {

	fmt.Println("repository status: ", status)

	order := &[]model.Order{}

	query := &dynamodb.ScanInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":order_status": {
				S: aws.String(status),
			},
		},
		FilterExpression: aws.String("order_status = :order_status"),
		TableName:        aws.String(orderCollection),
	}

	result, err := odb.orderDB.Scan(query)
	if err != nil {
		log.Info(result)
		log.Error("Failed to get item from database - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	if result.Items == nil {
		log.Error("Order for user with that status doesn't exist - ")
		notFoundError := error.NewNotFoundError("Order with that status for user doesn't exist")
		return nil, notFoundError
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, order)
	if err != nil {
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	return order, nil
}

func (odb OrderRepository) ReadOrderID(order_id string) (*model.Order, *error.AppError) {

	order := &model.Order{}

	query := &dynamodb.GetItemInput{
		TableName: aws.String(orderCollection),
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(order_id),
			},
		},
	}

	result, err := odb.orderDB.GetItem(query)
	if err != nil {
		log.Info(result)
		log.Error("Failed to get item from database - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	if result.Item == nil {
		log.Error("This order id doesn't exist - ")
		notFoundError := error.NewNotFoundError(" This order id for user doesn't exist")
		return nil, notFoundError
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, order)
	if err != nil {
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	return order, nil
}

func (odb OrderRepository) ReadCustomerID(customer_id string) (*[]model.Order, *error.AppError) {

	order := &[]model.Order{}

	query := &dynamodb.ScanInput{
		TableName:        aws.String(orderCollection),
		FilterExpression: aws.String("customer_id = :customer_id"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(customer_id),
			},
		},
	}

	result, err := odb.orderDB.Scan(query)
	if err != nil {
		log.Info(result)
		log.Error("Failed to get item from database - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	if result.Items == nil {
		log.Error("Orders for user doesn't exist - ")
		notFoundError := error.NewNotFoundError("Orders for user doesn't exist")
		return nil, notFoundError
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, order)
	if err != nil {
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	return order, nil
}

func (odb OrderRepository) ReadAll() (*[]model.Order, *error.AppError) {

	order := &[]model.Order{}

	query := &dynamodb.ScanInput{
		TableName: aws.String(orderCollection),
	}

	result, err := odb.orderDB.Scan(query)
	if err != nil {
		log.Info(result)
		log.Error("Failed to get item from database - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	if result.Items == nil {
		log.Error("Orders don't exist. - ")
		notFoundError := error.NewNotFoundError("Order don't exists")
		return nil, notFoundError
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, order)
	if err != nil {
		log.Error("Failed to unmarshal document fetched from DB - " + err.Error())
		return nil, error.NewUnexpectedError(err.Error())
	}

	return order, nil
}

func (odb OrderRepository) Update(order_id string, updated_status string) *error.AppError {

	fmt.Println("update repository orderid: ", order_id)
	fmt.Println("update repository status: ", updated_status)

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":order_status": {
				S: aws.String(updated_status),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(order_id),
			},
		},
		TableName:        aws.String(orderCollection),
		UpdateExpression: aws.String("set order_status = :order_status"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	_, err := odb.orderDB.UpdateItem(input)
	if err != nil {
		log.Error(err)
		return error.NewUnexpectedError(err.Error())
	}

	return nil
}

func (odb OrderRepository) Delete(order model.Order) *error.AppError {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(order.OrderId),
			},
		},
		TableName: aws.String(orderCollection),
	}

	_, err := odb.orderDB.DeleteItem(input)
	if err != nil {
		log.Error(err)
		return error.NewUnexpectedError(err.Error())
	}

	return nil
}

func (odb OrderRepository) CreateOrderInvoice(order_id string) *error.AppError {

	invoice_number := rand.Intn(10000)
	invoice_number_str := strconv.Itoa(invoice_number)

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":invoice": {
				S: aws.String(invoice_number_str),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"order_id": {
				S: aws.String(order_id),
			},
		},
		TableName:        aws.String(orderCollection),
		UpdateExpression: aws.String("set invoice = :invoice"),
		ReturnValues:     aws.String("UPDATED_NEW"),
	}

	_, err := odb.orderDB.UpdateItem(input)
	if err != nil {
		log.Error(err)
		return error.NewUnexpectedError(err.Error())
	}

	return nil
}

// func (odb OrderRepository) DeleteAll() *error.AppError {

// 	input := &dynamodb.ScanInput{
// 		TableName: aws.String(orderCollection),
// 	}

// 	result, err := odb.orderDB.Scan(input)
// 	if err != nil {
// 		log.Error(err)
// 		return error.NewUnexpectedError(err.Error())
// 	}

// 	if result.Items == nil {
// 		log.Error("Order for user doesn't exist. - ")
// 		notFoundError := error.NewNotFoundError("Order for user doesn't exists")
// 		return notFoundError
// 	}

// 	for _, item := range result.Items {
// 		input := &dynamodb.DeleteItemInput{
// 			Key: map[string]*dynamodb.AttributeValue{
// 				"id": {
// 					S: aws.String(*item["id"].S),
// 				},
// 			},
// 			TableName: aws.String(orderCollection),
// 		}

// 		_, err := odb.orderDB.DeleteItem(input)
// 		if err != nil {
// 			log.Error(err)
// 			return error.NewUnexpectedError(err.Error())
// 		}
// 	}

// 	return nil
// }
