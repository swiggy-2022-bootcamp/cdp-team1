package repository

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	_ "qwik.in/checkout/docs"
	"qwik.in/checkout/domain/database"
	"qwik.in/checkout/domain/tools/errs"
	"qwik.in/checkout/domain/tools/logger"
)

//Cart - STRUCT ..
type Cart struct {
	Id          string `json:"id,omitempty" dynamodbav:"id"`
	CustomerId  string `json:"customer_id,omitempty" dynamodbav:"customer_id"`
	CartProduct []struct {
		ProductId string `json:"product_id,omitempty" dynamodbav:"product_id"`
		Quantity  int    `json:"quantity" dynamodbav:"quantity"`
	} `json:"products" dynamodbav:"products"`
}

//CartRepo ..
type CartRepo interface {
	GetCartDetailsImpl() (*Cart, *errs.AppError)
}

//CartRepoImpl ..
type CartRepoImpl struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

//CartRepoFunc ..
func CartRepoFunc() CartRepoImpl {
	svc := database.CartConnectDB()
	return CartRepoImpl{Session: svc, Tablename: "team-1-cartCollection"}
}

// GetCartDetailsImpl  ..
func (cr CartRepoImpl) GetCartDetailsImpl() (*Cart, *errs.AppError) {
	data, err := cr.Session.Scan(&dynamodb.ScanInput{
		TableName: aws.String(cr.Tablename),
	})
	//fmt.Println(data) - DATA is FETCHED Correctly.
	if err != nil {
		fmt.Println("Error while Fetching Data from DynamoDB Table")
		logger.Error("Error while Fetching Data from DynamoDB Table", err)
	}
	if data.Items == nil {
		fmt.Println("Table Items - Empty")
		logger.Error("Error while Fetching Data from DynamoDB Table - Table Empty", err)
	}
	CartModelData := Cart{}
	err = dynamodbattribute.UnmarshalMap(data.Items[0], &CartModelData)
	if err != nil {
		fmt.Printf("Error unmarshalling the MAP to CartModelData %s", err)
		logger.Error("Error unmarshalling the MAP to CartModelData")
	}
	return &CartModelData, nil
}
