package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/internal/errors"
)

type CustomerRepositoryInterface interface {
	Create(customer model.Customer) (*model.Customer, error)
	GetById(customerId string) (*model.Customer, error)
	GetByEmail(customerEmail string) (*model.Customer, error)
	Update(customer model.Customer) (*model.Customer, error)
	Delete(customerId string) (*string, error)
}

type CustomerRepository struct {
}

var db *dynamodb.DynamoDB

func init() {
	db = GetDynamoDBInstance()
}

func (customerRepository *CustomerRepository) Create(customer model.Customer) (*model.Customer, error) {
	customer.CustomerId = uuid.New().String()
	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &customer, nil
	return nil, nil
}

func (customerRepository *CustomerRepository) GetById(customerId string) (*model.Customer, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
	}

	resp, err := db.GetItem(params)
	if err != nil {
		return nil, err
	}

	if len(resp.Item) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	var fetchedCustomer model.Customer
	dynamodbattribute.UnmarshalMap(resp.Item, &fetchedCustomer)
	return &fetchedCustomer, nil
}

func (customerRepository *CustomerRepository) GetByEmail(customerEmail string) (*model.Customer, error) {
	emailIndex := "email-index"
	params := &dynamodb.QueryInput{
		TableName:              aws.String("Customers"),
		IndexName:              &emailIndex,
		KeyConditionExpression: aws.String("#email = :customersEmail"),
		ExpressionAttributeNames: map[string]*string{
			"#email": aws.String("email"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":customersEmail": {
				S: aws.String(customerEmail),
			},
		},
	}

	resp, err := db.Query(params)
	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	var fetchedCustomer []model.Customer
	dynamodbattribute.UnmarshalListOfMaps(resp.Items, &fetchedCustomer)
	return &fetchedCustomer[0], nil
}

func (customerRepository *CustomerRepository) Update(customer model.Customer) (*model.Customer, error) {
	fetchedCustomer, err := customerRepository.GetById(customer.CustomerId)
	if err != nil {
		return nil, err
	}

	customer.DateAdded = fetchedCustomer.DateAdded

	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("Customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (customerRepository *CustomerRepository) Delete(customerId string) (*string, error) {
	allOld := "ALL_OLD"
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customerId": {
				S: aws.String(customerId),
			},
		},
		ReturnValues: &allOld,
	}

	deletedItem, err := db.DeleteItem(params)
	if err != nil {
		return nil, err
	}

	if len(deletedItem.Attributes) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	str := "deletion successful"
	return &str, nil
}
