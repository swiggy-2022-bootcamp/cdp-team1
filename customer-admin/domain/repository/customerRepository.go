package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/internal/errors"
)

type CustomerRepositoryInterface interface {
	Create(customer model.Customer) (*model.Customer, error)
	GetById(customerId string) (*model.Customer, error)
	GetByEmail(customerEmail string) (*model.Customer, error)
	Update(customer model.Customer) (*model.Customer, error)
	Delete(customerId string) (*string, error)
	GetCustomerAddress(customerId string) ([]model.Address, error)
	AddCustomerAddress(address model.Address) (bool, error)
}

type CustomerRepository struct {
}

var db *dynamodb.DynamoDB

func init() {
	db = GetDynamoDBInstance()
}

func (customerRepository *CustomerRepository) Create(customer model.Customer) (*model.Customer, error) {
	fetchedCustomer, _ := customerRepository.GetByEmail(customer.Email)
	if fetchedCustomer != nil {
		return nil, errors.NewEmailAlreadyRegisteredError()
	}

	customer.CustomerId = uuid.New().String()

	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("team-1-customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, &errors.CustomerError{Status: 400, ErrorMessage: err.Error()}
	}
	return &customer, nil
}

func (customerRepository *CustomerRepository) GetById(customerId string) (*model.Customer, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("team-1-customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(customerId),
			},
		},
	}

	resp, err := db.GetItem(params)
	if err != nil {
		return nil, &errors.CustomerError{Status: 400, ErrorMessage: err.Error()}
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
		TableName:              aws.String("team-1-customers"),
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
		return nil, &errors.CustomerError{Status: 400, ErrorMessage: err.Error()}
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

	//if email is updated, check if the changed email is not already being used my other users
	if fetchedCustomer.Email != customer.Email {
		fetchedCustomerWithEmail, _ := customerRepository.GetByEmail(customer.Email)
		if fetchedCustomerWithEmail != nil {
			return nil, errors.NewEmailAlreadyRegisteredError()
		}
	}

	customer.DateAdded = fetchedCustomer.DateAdded

	info, err := dynamodbattribute.MarshalMap(customer)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("team-1-customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, &errors.CustomerError{Status: 400, ErrorMessage: err.Error()}
	}
	return &customer, nil
}

func (customerRepository *CustomerRepository) Delete(customerId string) (*string, error) {
	allOld := "ALL_OLD"
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("team-1-customers"),
		Key: map[string]*dynamodb.AttributeValue{
			"customer_id": {
				S: aws.String(customerId),
			},
		},
		ReturnValues: &allOld,
	}

	deletedItem, err := db.DeleteItem(params)
	if err != nil {
		return nil, &errors.CustomerError{Status: 400, ErrorMessage: err.Error()}
	}

	if len(deletedItem.Attributes) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	str := "deletion successful"
	return &str, nil
}

func (customerRepository *CustomerRepository) GetCustomerAddress(customerId string) ([]model.Address, error) {
	url := "http://localhost:9005/api/shipping/allAddressOfCustomer/" + customerId

	// Create a new request using http
	req, _ := http.NewRequest("GET", url, nil)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	if resp.StatusCode != http.StatusAccepted {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return nil, &errors.CustomerError{Status: resp.StatusCode, ErrorMessage: bodyString}
	}
	var fetchedCustomerAddresses []model.Address
	json.NewDecoder(resp.Body).Decode(&fetchedCustomerAddresses)
	return fetchedCustomerAddresses, nil
}

func (customerRepository *CustomerRepository) AddCustomerAddress(address model.Address) (bool, error) {
	url := "http://localhost:9005/api/shipping/newAddress"

	// Marshal it into JSON prior to requesting
	addressJSON, err := json.Marshal(address)

	// Create a new request using http
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(addressJSON))

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp)
	if resp.StatusCode != http.StatusAccepted {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return false, &errors.CustomerError{Status: resp.StatusCode, ErrorMessage: bodyString}
	}

	return true, nil
}
