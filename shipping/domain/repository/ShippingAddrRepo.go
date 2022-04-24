package repository

import (
	"context"
	"fmt"
	"github.com/ashwin2125/qwk/shipping/domain/database"
	"github.com/ashwin2125/qwk/shipping/domain/models"
	"github.com/ashwin2125/qwk/shipping/domain/tools/errs"
	"github.com/ashwin2125/qwk/shipping/domain/tools/logger"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
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
	ShippingCost      int    `json:"shipping_cost" dynamodb:"shipping_cost"`
}

//ShippingAddrRepo ..
type ShippingAddrRepo interface {
	CreateNewShippingAddrImpl(ShippingAddress) (string, *errs.AppError)
	FindShippingAddressByIdImpl(string) (*ShippingAddress, *errs.AppError)
	FindDefaultShippingAddressImpl(bool) (*ShippingAddress, *errs.AppError)
}

//ShippingAddressRepoImpl ..
type ShippingAddressRepoImpl struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

//ShippingAddrFunc ..
func ShippingAddrFunc(userID int, shippingAddressId, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool, shippingCost int) *ShippingAddress {
	shippingCost = GetShippingCost(pincode)
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
		ShippingCost:      shippingCost,
	}
}

func toPersistedDynamodbEntitySA(o ShippingAddress) *models.ShippingAddrModel {

	return &models.ShippingAddrModel{
		UserID: o.UserID,
		//ShippingAddressID: uuid.New().String(),
		ShippingAddressID: o.ShippingAddressID,
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
		ShippingCost:      o.ShippingCost,
	}
}

//ShippingAddrRepoFunc ..
func ShippingAddrRepoFunc() ShippingAddressRepoImpl {
	svc := database.ConnectDB()
	return ShippingAddressRepoImpl{Session: svc, Tablename: "team-1-shipping"}
}

//CreateNewShippingAddrImpl ..
func (sar ShippingAddressRepoImpl) CreateNewShippingAddrImpl(p ShippingAddress) (string, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ShippingAddressRecord := toPersistedDynamodbEntitySA(p)
	av, err := dynamodbattribute.MarshalMap(ShippingAddressRecord)
	if err != nil {
		return "", &errs.AppError{Message: fmt.Sprintf("Unable to Marshal ! - %s", err.Error())}
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("team-1-shipping"),
	}

	_, err = sar.Session.PutItemWithContext(ctx, input)

	if err != nil {
		return "", &errs.AppError{Message: fmt.Sprintf("Unable to Put the Item in DB - %s", err.Error())}
	}

	return ShippingAddressRecord.ShippingAddressID, nil
}

//FindShippingAddressByIdImpl ..
func (sar ShippingAddressRepoImpl) FindShippingAddressByIdImpl(shippingAddressId string) (*ShippingAddress, *errs.AppError) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	input := &dynamodb.GetItemInput{
		TableName: aws.String("team-1-shipping"),
		Key: map[string]*dynamodb.AttributeValue{
			"shipping_address_id": {
				S: aws.String(shippingAddressId),
			},
		},
	}

	result, err := sar.Session.GetItemWithContext(ctx, input)
	if err != nil {
		return nil, &errs.AppError{Message: fmt.Sprintf("Unable to Get the Item in DB - %s", err.Error())}
	}

	if result.Item == nil {
		return nil, &errs.AppError{Message: fmt.Sprintf("Unable to Get the Item - %s", err.Error())}
	}

	ShippingAddressModel := models.ShippingAddrModel{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &ShippingAddressModel)

	if err != nil {
		return nil, &errs.AppError{Message: fmt.Sprintf("Unable to Unmarshal map !  - %s", err.Error())}
	}

	return (*ShippingAddress)(&ShippingAddressModel), nil
}

// FindDefaultShippingAddressImpl  ..
func (sar ShippingAddressRepoImpl) FindDefaultShippingAddressImpl(isDefaultAddress bool) (*ShippingAddress, *errs.AppError) {
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

//GetShippingCost ..
func GetShippingCost(pincode int) int {
	division := pincode / 100000
	var cost int
	switch division {
	case 1:
		// Northern Zone : Delhi, Haryana, Punjab, Himachal Pradesh and Jammu & Kashmir
		cost = 75
	case 2:
		// Northern Zone 2 : Uttar Pradesh and Uttarakhand
		cost = 60
	case 3:
		// Western Zone 1 : Rajasthan and Gujarat
		cost = 50
	case 4:
		// Western Zone 2 : Maharashtra, Madhya Pradesh and Chattisgarh
		cost = 40
	case 5:
		// Southern Zone 1 : Andhra Pradesh and Karnataka
		cost = 40
	case 6:
		// Southern Zone 2 : Kerala and Tamil Nadu
		cost = 50
	case 7:
		// Eastern Zone 1 : West Bengal, Orissa and North Eastern States
		cost = 75
	case 8:
		// Eastern Zone 2 : Bihar and Jharkhand
		cost = 60
	case 9:
		// APS Zone - Army Prohibited Service Zone
		cost = 0
	default:
		// Default Shipping Cost
		cost = 50
	}
	return cost
}
