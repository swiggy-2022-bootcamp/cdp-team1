package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/internal/errors"
)

type AdminUserRepositoryInterface interface {
	GetAll() []model.AdminUser
	Create(admin model.AdminUser) (*model.AdminUser, error)
}

type AdminUserRepository struct {
}

func init() {
	db = GetDynamoDBInstance()
}

func (adminUser *AdminUserRepository) GetAll() []model.AdminUser {
	input := &dynamodb.ScanInput{
		TableName: aws.String("team-1-admins"),
	}

	scanOutput, _ := db.Scan(input)
	var adminUsers []model.AdminUser
	dynamodbattribute.UnmarshalListOfMaps(scanOutput.Items, &adminUsers)
	return adminUsers
}

func (adminUser *AdminUserRepository) Create(admin model.AdminUser) (*model.AdminUser, error) {
	admin.UserId = uuid.New().String()
	info, err := dynamodbattribute.MarshalMap(admin)
	if err != nil {
		return nil, errors.NewMarshallError()
	}

	input := &dynamodb.PutItemInput{
		Item:      info,
		TableName: aws.String("team-1-admins"),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return nil, &errors.CustomerError{Status: 400, ErrorMessage: err.Error()}
	}
	return &admin, nil
}
