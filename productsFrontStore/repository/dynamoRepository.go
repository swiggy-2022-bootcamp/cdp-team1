package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"qwik.in/productsFrontStore/config"
	"qwik.in/productsFrontStore/entity"
	"qwik.in/productsFrontStore/log"
)

type dynamoRepository struct{}

var db *dynamodb.DynamoDB

func NewDynamoRepository() ProductRepository {
	return &dynamoRepository{}
}

func (d dynamoRepository) Connect() error {
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(config.DYNAMO_DB_REGION),
		Endpoint: aws.String(config.DYNAMO_DB_URL),
	}))

	// create a dynamodb instance
	db = dynamodb.New(sess)
	return nil
}

func (d dynamoRepository) FindAll() ([]entity.Product, error) {
	// create the api params
	params := &dynamodb.ScanInput{
		TableName: aws.String("Products"),
	}

	var productList []entity.Product

	// scan and filter for the items
	err := db.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		// Unmarshal the slice of dynamodb attribute values into a slice of custom structs
		var products []entity.Product
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &products)
		if err != nil {
			log.Error("\nCould not unmarshal AWS data: err = %v\n", err)
			return true
		}

		productList = append(productList, products...)

		return true
	})

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return productList, nil
}

func (d dynamoRepository) FindById(productId string) (entity.Product, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Products"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(productId),
			},
		},
	}

	resp, err := db.GetItem(params)
	if err != nil {
		log.Error(err.Error())
		return entity.Product{}, err
	}

	var product entity.Product
	err = dynamodbattribute.UnmarshalMap(resp.Item, &product)

	if err != nil {
		log.Error(err.Error())
		return entity.Product{}, err
	}

	return product, nil
}

func (d dynamoRepository) FindByCategory(categoryId string) ([]entity.Product, error) {

	filter := expression.Name("product_category").Contains(categoryId)
	expr, _ := expression.NewBuilder().WithFilter(filter).Build()

	// create the api params
	params := &dynamodb.ScanInput{
		TableName:                 aws.String("Products"),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
	}

	var productList []entity.Product

	// scan and filter for the items
	err := db.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		// Unmarshal the slice of dynamodb attribute values into a slice of custom structs
		var products []entity.Product
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &products)
		if err != nil {
			log.Error("\nCould not unmarshal AWS data: err = %v\n", err)
			return true
		}

		productList = append(productList, products...)

		return true
	})

	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	return productList, nil
}
