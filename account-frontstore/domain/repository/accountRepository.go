package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/internal/errors"
)

type AccountRepositoryInterface interface {
	Create(account model.Account) (*model.Account, error)
	GetById(customerId string) (*model.Account, error)
	GetByEmail(customerEmail string) (*model.Account, error)
	Update(account model.Account) (*model.Account, error)
}

type AccountRepository struct {
}

var db *dynamodb.DynamoDB

func init() {
	db = GetDynamoDBInstance()
}

func (accountRepository *AccountRepository) Create(account model.Account) (*model.Account, error) {
	account.CustomerId = uuid.New().String()
	info, err := dynamodbattribute.MarshalMap(account)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("team-1-customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, &errors.AccountError{Status: 400, ErrorMessage: err.Error()}
	}
	return &account, nil
}

func (accountRepository *AccountRepository) GetById(customerId string) (*model.Account, error) {
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
		return nil, &errors.AccountError{Status: 400, ErrorMessage: err.Error()}
	}

	if len(resp.Item) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	var fetchedAccount model.Account
	dynamodbattribute.UnmarshalMap(resp.Item, &fetchedAccount)
	return &fetchedAccount, nil
}

func (accountRepository *AccountRepository) Update(account model.Account) (*model.Account, error) {
	fetchedAccount, err := accountRepository.GetById(account.CustomerId)
	if err != nil {
		return nil, err
	}

	//if email is updated, check if the changed email is not already being used my other users
	if fetchedAccount.Email != account.Email {
		fetchedAccountWithEmail, _ := accountRepository.GetByEmail(account.Email)
		if fetchedAccountWithEmail != nil {
			return nil, errors.NewEmailAlreadyRegisteredError()
		}
	}

	account.DateAdded = fetchedAccount.DateAdded

	info, err := dynamodbattribute.MarshalMap(account)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("team-1-customers"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, &errors.AccountError{Status: 400, ErrorMessage: err.Error()}
	}
	return &account, nil
}

func (accountRepository *AccountRepository) GetByEmail(customerEmail string) (*model.Account, error) {
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
		return nil, &errors.AccountError{Status: 400, ErrorMessage: err.Error()}
	}

	if len(resp.Items) == 0 {
		return nil, errors.NewUserNotFoundError()
	}

	var fetchedAccount []model.Account
	dynamodbattribute.UnmarshalListOfMaps(resp.Items, &fetchedAccount)
	return &fetchedAccount[0], nil
}
