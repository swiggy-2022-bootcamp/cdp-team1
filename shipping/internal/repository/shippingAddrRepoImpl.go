package repository

import (
	"context"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"qwik.in/shipping/internal/tools/log"
)

const shippingAddressCollection = "shippingAddressCollection"

//ShippingAddrRepo ..
type ShippingAddrRepo interface {
	DBHealthCheck() bool
}

//ShippingAddrRepoImpl ...
type ShippingAddrRepoImpl struct {
	shippingAddressDB *dynamodb.DynamoDB
	ctx               context.Context
}

//ShippingAddrRepoImplFunc ..
func ShippingAddrRepoImplFunc(ctx context.Context, shippingAddressDB *dynamodb.DynamoDB) ShippingAddrRepo {
	return &ShippingAddrRepoImpl{
		shippingAddressDB: shippingAddressDB,
		ctx:               ctx,
	}
}

//DBHealthCheck ..
func (s ShippingAddrRepoImpl) DBHealthCheck() bool {
	_, err := s.shippingAddressDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Database connection is down.")
		return false
	}
	return true
}
