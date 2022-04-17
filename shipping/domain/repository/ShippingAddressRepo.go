package repository

import (
	"context"
	"fmt"
	"github.com/ashwin2125/qwk/shipping/domain/database"
	"github.com/ashwin2125/qwk/shipping/domain/models"
	"github.com/ashwin2125/qwk/shipping/domain/tools/errs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/google/uuid"
	"time"
)

//ShippingAddress ..
type ShippingAddress struct {
	UserID            int    `json:"user_id" dynamodb:"user_id"`
	ShippingAddressID string `json:"shipping_address_id" dynamodbav:"shipping_address_id"`
	FirstName         string `json:"first_name" dynamodbav:"first_name"`
	LastName          string `json:"last_name" dynamodbav:"lastName"`
	AddressLine1      string `json:"address_line_1" dynamodbav:"address_line_1" `
	AddressLine2      string `json:"address_line_2" dynamodbav:"address_line_2"`
	City              string `json:"city" dynamodbav:"city"`
	State             string `json:"state" dynamodbav:"state"`
	Phone             string `json:"phone" dynamodbav:"phone"`
	Pincode           int    `json:"pincode" dynamodbav:"pincode"`
	AddressType       string `json:"address_type" dynamodbav:"address_type"`
	DefaultAddress    bool   `json:"default_address" dynamodbav:"default_address"`
}

//ShippingAddressFunc ..
func ShippingAddressFunc(userID int, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool) *ShippingAddress {
	return &ShippingAddress{
		UserID:         userID,
		FirstName:      firstName,
		LastName:       lastName,
		AddressLine1:   addressLine1,
		AddressLine2:   addressLine2,
		City:           city,
		State:          state,
		Phone:          phone,
		Pincode:        pincode,
		AddressType:    addressType,
		DefaultAddress: defaultAddress,
	}
}

//ShippingAddrRepo ..
type ShippingAddrRepo interface {
	InsertShippingAddress(ShippingAddress) (string, *errs.AppError)
}

//ShippingAddressRepoImpl ..
type ShippingAddressRepoImpl struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

//InsertShippingAddress ..
func (sdr ShippingAddressRepoImpl) InsertShippingAddress(p ShippingAddress) (string, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ShippingAddressRecord := toPersistedDynamodbEntitySA(p)
	av, err := dynamodbattribute.MarshalMap(ShippingAddressRecord)
	if err != nil {
		return "", &errs.AppError{Message: fmt.Sprintf("unable to marshal - %s", err.Error())}
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("team-1-shipping"),
	}

	_, err = sdr.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return "", &errs.AppError{Message: fmt.Sprintf("unable to put the item - %s", err.Error())}
	}

	return ShippingAddressRecord.ShippingAddressID, nil
}

func toPersistedDynamodbEntitySA(o ShippingAddress) *models.ShippingAddressModel {
	return &models.ShippingAddressModel{
		UserID:            o.UserID,
		ShippingAddressID: uuid.New().String(),
		FirstName:         o.FirstName,
		LastName:          o.LastName,
		AddressLine1:      o.AddressLine1,
		AddressLine2:      o.AddressLine2,
		City:              o.City,
		State:             o.State,
		Phone:             o.Phone,
		Pincode:           o.Pincode,
		AddressType:       o.AddressType,
		DefaultAddress:    o.DefaultAddress,
	}
}

//ShippingAddressRepositoryFunc ..
func ShippingAddressRepositoryFunc() ShippingAddressRepoImpl {
	svc := database.Connect()
	return ShippingAddressRepoImpl{Session: svc, Tablename: "team-1-shipping"}
}
