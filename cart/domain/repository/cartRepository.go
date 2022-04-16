package repository

import (
	"cartService/domain/model"
	"cartService/internal/error"
	"cartService/log"
	"context"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartRepositoryDB interface {
	Create(model.Cart) (*model.Cart, *error.AppError)
	Read(id string) (*model.Cart, *error.AppError)
	ReadAll() (*[]model.Cart, *error.AppError)
	Update(model.Cart) (*model.Cart, *error.AppError)
	Delete(model.Cart) (*model.Cart, *error.AppError)
	DeleteAll() *error.AppError
	DBHealthCheck() bool
}

type CartRepository struct {
	cartDB *dynamodb.DynamoDB
	ctx    context.Context
}

func NewCartRepository(cartDB *dynamodb.DynamoDB, ctx context.Context) CartRepositoryDB {
	return &CartRepository{
		cartDB: cartDB,
		ctx:    ctx,
	}
}

func (cr CartRepository) DBHealthCheck() bool {

	_, err := cr.cartDB.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Error("Database connection is down.")
		return false
	}
	return true
}

func (cdb CartRepository) Create(cart model.Cart) (*model.Cart, *error.AppError) {

	newCart := model.Cart{
		CustomerId: cart.CustomerId,
		Products:   cart.Products,
	}

	newCart.Id = primitive.NewObjectID().Hex()

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "cart")
	// _, err := cartCollection.InsertOne(ctx, newCart)

	//if err != nil {
	//	return nil, error.NewUnexpectedError("Unexpected error from DB")
	//}

	cart.Id = newCart.Id

	return &cart, nil
}

func (cdb CartRepository) Read(id string) (*model.Cart, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "cart")
	// order := model.Cart{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &model.Cart{}, nil
}

func (cdb CartRepository) ReadAll() (*[]model.Cart, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "cart")
	// order := model.Cart{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &[]model.Cart{}, nil
}

func (cdb CartRepository) Update(cart model.Cart) (*model.Cart, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "cart")
	// _, err := orderCollection.UpdateOne(ctx, bson.M{"_id": cart.Id}, bson.M{"$set": cart})

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &model.Cart{}, nil
}

func (cdb CartRepository) Delete(cart model.Cart) (*model.Cart, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "cart")
	// _, err := orderCollection.DeleteOne(ctx, bson.M{"_id": cart.Id})

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &model.Cart{}, nil
}

func (cdb CartRepository) DeleteAll() *error.AppError {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "cart")
	// _, err := orderCollection.DeleteMany(ctx, bson.M{})

	// if err != nil {
	// 	return error.NewUnexpectedError("Unexpected error from DB")
	// }

	return nil
}
