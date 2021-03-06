package repository

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"qwik.in/productsAdmin/config"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
)

type dynamoRepository struct{}

var db *dynamodb.DynamoDB

func NewDynamoRepository() ProductRepository {
	return &dynamoRepository{}
}

func (r dynamoRepository) Connect() error {
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.DYNAMO_DB_REGION),
	}))

	// create a dynamodb instance
	db = dynamodb.New(sess)

	// create Products table
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Products"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("id"), KeyType: aws.String("HASH")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("id"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(5),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := db.CreateTable(params)

	if err != nil {
		target := &dynamodb.ResourceInUseException{}
		if errors.As(err, &target) {
			fmt.Println("Products table already present")
		} else {
			log.Error("Error while creating table Products", err.Error())
			return err
		}
	} else {
		log.Info("Table created Successfully")
	}

	return nil
}

func (r dynamoRepository) FindOne(productId string) (entity.Product, error) {

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

func (r dynamoRepository) FindAll() ([]entity.Product, error) {

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
			fmt.Printf("\nCould not unmarshal AWS data: err = %v\n", err)
			return true
		}

		productList = append(productList, products...)

		return true
	})

	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		log.Error(err.Error())
		return nil, err
	}

	return productList, nil
}

func (r dynamoRepository) SaveProduct(product entity.Product) error {

	productAVMap, err := dynamodbattribute.MarshalMap(product)
	if err != nil {
		return err
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("Products"),
		Item:      productAVMap,
	}

	resp, err := db.PutItem(params)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Info(resp)
		return nil
	}
}

func (r dynamoRepository) DeleteProduct(productId string) error {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Products"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(productId),
			},
		},
	}

	resp, err := db.DeleteItem(params)
	if err != nil {
		log.Error(err.Error())
		return err
	} else {
		log.Info("Success")
		log.Info(resp)
		return nil
	}
}

func (r dynamoRepository) FindAndUpdate(productId string, product entity.Product) error {
	// TODO implement me
	panic("Implement me")
}

func (r dynamoRepository) FindWithLimit(limit int64) ([]entity.Product, error) {

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
			fmt.Printf("\nCould not unmarshal AWS data: err = %v\n", err)
			return true
		}

		if limit-int64(len(products)) >= 0 {
			productList = append(productList, products...)
			limit -= int64(len(products))
		} else {
			productList = append(productList, products[0:limit]...)
			limit -= limit
		}

		return true
	})

	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		log.Error(err.Error())
		return nil, err
	}

	return productList, nil
}
