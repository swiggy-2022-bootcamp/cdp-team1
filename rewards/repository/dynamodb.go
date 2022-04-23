package repository

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"qwik.in/rewards/config"
	"qwik.in/rewards/entity"
	"qwik.in/rewards/log"
)

type dynamoRepository struct{}

var db *dynamodb.DynamoDB

func NewDynamoRepository() RewardRepository {
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

func (r dynamoRepository) FindOne(rewardId string) (entity.Reward, error) {
	params := &dynamodb.GetItemInput{
		TableName: aws.String("Rewards"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(rewardId),
			},
		},
	}

	resp, err := db.GetItem(params)
	if err != nil {
		log.Error(err.Error())
		return entity.Reward{}, err
	}

	var reward entity.Reward
	err = dynamodbattribute.UnmarshalMap(resp.Item, &reward)

	if err != nil {

		return entity.Reward{}, err
	}

	return reward, nil
}
func (r dynamoRepository) FindAll() ([]entity.Reward, error) {

	// create the api params
	params := &dynamodb.ScanInput{
		TableName: aws.String("Rewards"),
	}

	var rewardList []entity.Reward

	// scan and filter for the items
	err := db.ScanPages(params, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		// Unmarshal the slice of dynamodb attribute values into a slice of custom structs
		var rewards []entity.Reward
		err := dynamodbattribute.UnmarshalListOfMaps(page.Items, &rewards)
		if err != nil {
			fmt.Printf("\nCould not unmarshal AWS data: err = %v\n", err)
			return true
		}

		rewardList = append(rewardList, rewards...)
		fmt.Printf("\n%v\n", rewardList)
		return true
	})

	if err != nil {
		fmt.Printf("ERROR: %v\n", err.Error())
		log.Error(err.Error())
		return nil, err
	}

	return rewardList, nil
}
func (r dynamoRepository) DeleteReward(rewardId string) error {
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("Rewards"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(rewardId),
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
func (r dynamoRepository) SaveReward(Reward entity.Reward) error {

	RewardAVMap, err := dynamodbattribute.MarshalMap(Reward)
	if err != nil {
		return err
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String("Rewards"),
		Item:      RewardAVMap,
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
