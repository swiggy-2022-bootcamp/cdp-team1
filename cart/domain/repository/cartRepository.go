package repository

import (
	"cartService/domain/model"
	"cartService/internal/error"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartRepositoryDB interface {
	Create(model.Cart) (*model.Cart, *error.AppError)
	Read(id string) (*model.Cart, *error.AppError)
	ReadAll() (*[]model.Cart, *error.AppError)
	Update(model.Cart) (*model.Cart, *error.AppError)
	Delete(model.Cart) (*model.Cart, *error.AppError)
	DeleteAll() *error.AppError
}

type CartRepository struct {
	// dbClient *mongo.Client
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
