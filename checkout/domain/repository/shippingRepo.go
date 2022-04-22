package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	_ "qwik.in/checkout/docs"
	"qwik.in/checkout/domain/database"
	"qwik.in/checkout/domain/models"
	"qwik.in/checkout/domain/tools/errs"
	"qwik.in/checkout/domain/tools/logger"
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

//ShippingRepo ..
type ShippingRepo interface {
	FindDefaultShippingAddressImpl(bool) (*ShippingAddress, *errs.AppError)
}

//ShippingRepoImpl ..
type ShippingRepoImpl struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

//ShippingFunc ..
func ShippingFunc(userID int, shippingAddressId, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool) *ShippingAddress {
	return &ShippingAddress{
		UserID:            userID,
		ShippingAddressID: shippingAddressId,
		FirstName:         firstName,
		LastName:          lastName,
		AddressLine1:      addressLine1,
		AddressLine2:      addressLine2,
		City:              city,
		State:             state,
		Phone:             phone,
		Pincode:           pincode,
		AddressType:       addressType,
		DefaultAddress:    defaultAddress,
	}
}

//ShippingAddrRepoFunc ..
func ShippingAddrRepoFunc() ShippingRepoImpl {
	svc := database.ShippingConnectDB()
	return ShippingRepoImpl{Session: svc, Tablename: "team-1-shipping"}
}

// FindDefaultShippingAddressImpl  ..
func (sar ShippingRepoImpl) FindDefaultShippingAddressImpl(isDefaultAddress bool) (*ShippingAddress, *errs.AppError) {
	item := models.ShippingAddrModel{}
	filt := expression.Name("default_address").Equal(expression.Value(isDefaultAddress))
	expr, err := expression.NewBuilder().WithFilter(filt).Build()
	if err != nil {
		fmt.Println("Error in Expression Builder")
		logger.Error("Got Error building expression: %s", err)
	}
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("team-1-shipping"),
	}
	result, err := sar.Session.Scan(params)
	if err != nil {
		fmt.Println("Error in API Query")
		logger.Error("Query API call failed - %s", err)
	}
	for _, i := range result.Items {
		err = dynamodbattribute.UnmarshalMap(i, &item)
		if err != nil {
			fmt.Println(item)
			fmt.Println("Error in Unmarshalling the values")
			logger.Error("Got error while unmarshalling - %s", err)
		}
	}
	return (*ShippingAddress)(&item), nil
}
