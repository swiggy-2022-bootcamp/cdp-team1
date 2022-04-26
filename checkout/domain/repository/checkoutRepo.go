package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"qwik.in/checkout/domain/tools/errs"
)

//CheckoutRepo ..
type CheckoutRepo interface {
	GetDefaultShippingAddressFlowImpl() (*ShippingAddress, *errs.AppError)
	GetConfirmFlowImpl() (bool, *errs.AppError)
}

//CheckoutRepoImpl ..
type CheckoutRepoImpl struct {
	Session   *dynamodb.DynamoDB
	Tablename string
}
