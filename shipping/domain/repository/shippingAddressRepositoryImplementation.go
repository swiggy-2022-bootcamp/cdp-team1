package repository

import (
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"qwik.in/shipping/internal/log"
)

const shippingAddressCollection = "shippingAddressCollection6"

//ShippingAddressRepositoryImplementation ...
type ShippingAddressRepositoryImplementation struct {
	shippingAddressDB *dynamodb.DynamoDB
	ctx               context.Context
}

//NewShippingAddressRepositoryImplementation ..
func NewShippingAddressRepositoryImplementation(shippingAddressDB *dynamodb.DynamoDB, ctx context.Context) ShippingAddressRepository {
	return &ShippingAddressRepositoryImplementation{
		shippingAddressDB: shippingAddressDB,
		ctx:               ctx,
	}
}

//DBHealthCheck ..
func (s ShippingAddressRepositoryImplementation) DBHealthCheck() bool {
	_, err := s.shippingAddressDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Database connection is down.")
		return false
	}
	return true
}
