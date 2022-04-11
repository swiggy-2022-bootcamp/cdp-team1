package repository

import (
	"cartService/domain/db"
	error "cartService/internal"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CartRepositoryDB interface {
	Create(db.Cart) (*db.Cart, *error.AppError)
	Read(id string) (*db.Cart, *error.AppError)
	Update(db.Cart) (*db.Cart, *error.AppError)
	Delete(db.Cart) (*db.Cart, *error.AppError)
	DeleteAll() *error.AppError
}

type CartRepository struct {
	// dbClient *mongo.Client
}

func (cdb CartRepository) Create(cart db.Cart) (*db.Cart, *error.AppError) {

	newCart := db.Cart{
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

func (cdb CartRepository) Read(id string) (*db.Cart, *error.AppError) {

	// ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cxl()

	// orderCollection := Collection(cdb.dbClient, "cart")
	// order := db.Cart{}
	// err := orderCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&order)

	// if err != nil {
	// 	return nil, error.NewUnexpectedError("Unexpected error from DB")
	// }

	return &db.Cart{}, nil
}
