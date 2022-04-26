package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"qwik.in/checkout/domain/database"
	"qwik.in/checkout/domain/tools/errs"
	"qwik.in/checkout/domain/tools/logger"
)

//PaymentUserPaymentMode ..
type PaymentUserPaymentMode struct {
	OrderAmount      float64 `json:"order_amount" dynamodbav:"order_amount,omitempty"`
	UserID           string  `json:"user_id" dynamodbav:"user_id,omitempty"`
	PaymentModeModel []struct {
		Mode       string `json:"mode_of_payment" dynamodbav:"mode_of_payment"`
		Credential string `json:"credential" dynamodbav:"credential"`
	} `json:"payment_mode" dynamodbav:"payment_mode"`
}

//PaymentRepo ..
type PaymentRepo interface {
	GetPaymentDetailsImpl() (*PaymentUserPaymentMode, *errs.AppError)
}

//PaymentRepoImpl ..
type PaymentRepoImpl struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}

//PaymentRepoFunc ..
func PaymentRepoFunc() PaymentRepoImpl {
	svc := database.PaymentConnectDB()
	return PaymentRepoImpl{Session: svc, Tablename: "team-1-paymentCollection"}
}

//GetPaymentDetailsImpl ..
func (pr PaymentRepoImpl) GetPaymentDetailsImpl() (*PaymentUserPaymentMode, *errs.AppError) {
	data, err := pr.Session.Scan(&dynamodb.ScanInput{
		TableName: aws.String(pr.Tablename),
	})
	if err != nil {
		fmt.Println("Error while Fetching Data from DynamoDB Table")
		logger.Error("Error while Fetching Data from DynamoDB Table", err)
	}
	if data.Items == nil {
		fmt.Println("Table Items - Empty")
		logger.Error("Table Items are either empty (or) some error while fetching the data")
	}
	PaymentModelData := PaymentUserPaymentMode{}
	err = dynamodbattribute.UnmarshalMap(data.Items[0], &PaymentModelData)
	if err != nil {
		fmt.Printf("Error while unmarshalling the MAP to PaymentModelData %s", err)
		logger.Error("Error while unmarshalling the MAP to PaymentModelData")
	}
	return &PaymentModelData, nil
}
