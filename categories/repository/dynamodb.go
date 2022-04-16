package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"qwik.in/categories/config"
	"qwik.in/categories/entity"
	"qwik.in/categories/log"
)

type dynamoRepository struct{}

var db *dynamodb.DynamoDB

func NewDynamoRepository() CategoryRepository {
	return &dynamoRepository{}
}

func (r dynamoRepository) Connect() error {
	// create an aws session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String(config.DYNAMO_DB_REGION),
		Endpoint: aws.String(config.DYNAMO_DB_URL),
	}))

	// create a dynamodb instance
	db = dynamodb.New(sess)
	return nil
}
func (r dynamoRepository) FindAll() ([]entity.Category, error) {

	// create the api params
	params := &dynamodb.ScanInput{
		TableName: aws.String("Categories"),
	}

	var CategoryList []entity.Category

	// scan and filter for the items
	err := db.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		// Unmarshal the slice of dynamodb attribute values into a slice of custom structs
		var categories []entity.Category
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &categories)
		if err != nil {
			fmt.Printf("\nCould not unmarshal AWS data: err = %v\n", err)
			return true
		}

		CategoryList = append(CategoryList, categories...)

		return true
	})

	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		log.Error(err.Error())
		return nil, err
	}

	return CategoryList, nil
}

func (r dynamoRepository) FindOne(categoryId string) (entity.Category, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Categories"),
		Key: map[string]*dynamodb.AttributeValue{
			"category_id": {
				S: aws.String(categoryId),
			},
		},
	}

	resp, err := db.GetItem(params)
	if err != nil {
		log.Error(err.Error())
		return entity.Category{}, err
	}

	var category entity.Category
	err = dynamodbattribute.UnmarshalMap(resp.Item, &category)

	if err != nil {

		return entity.Category{}, err
	}

	return category, nil
}
func (r dynamoRepository) SaveCategory(category entity.Category) error {

	productAVMap, err := dynamodbattribute.MarshalMap(category)
	if err != nil {
		return err
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("Categories"),
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
