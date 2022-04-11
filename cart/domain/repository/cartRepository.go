package repository

import (
	"cartService/domain"
	"cartService/internal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//type CartRepositoryDB interface {
//	//Create(cart *domain.Cart) (*domain.Cart, *error.AppError)
//	SaveItem(userId string, item domain.Item) (domain.Cart, error)
//	RemoveItem(userId string, itemId string) (domain.Cart, error)
//	UpdateItem(userId string, item domain.Item) (domain.Cart, error)
//	GetItems(userId string) ([]domain.Item, error)
//}

type CartRepository struct{}

func (cdb CartRepository) Create(cart domain.Cart) (*domain.Cart, *error.AppError) {

	newCart := domain.Cart{
		CustomerId: cart.CustomerId,
		Products:   cart.Products,
	}

	newCart.Id = primitive.NewObjectID().Hex()

	//ctx, cxl := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cxl()

	//orderCollection := Collection(cdb.dbClient, "cart")
	//_, err := cartCollection.InsertOne(ctx, newCart)

	//if err != nil {
	//	return nil, error.NewUnexpectedError("Unexpected error from DB")
	//}

	cart.Id = newCart.Id

	return &cart, nil
}
