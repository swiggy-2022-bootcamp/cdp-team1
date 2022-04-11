package domain

import (
	"cartService/domain/repository"
	"cartService/internal"
)

type CartService interface {
	Add(Cart) (*Cart, *error.AppError)
	Update(string) (*Cart, *error.AppError)
	GetAll(string, Cart) (*Cart, *error.AppError)
	DeleteById(string) (*Cart, *error.AppError)
	DeleteAll(string) (*Cart, *error.AppError)
}

type DefaultCartService struct {
	CartDB repository.CartRepositoryDB
}

//func NewCartService(CartDB CartRepositoryDB) CartService {
//	return &DefaultCartService{
//		CartDB: CartDB,
//	}
//}

func (csvc DefaultCartService) Add(cart Cart) (*Cart, *error.AppError) {

	u, err := csvc.CartDB.Create(cart)

	if err != nil {
		return nil, err
	}
	return u, err
}

func (csvc DefaultCartService) Update(id string, Cart Cart) (*Cart, *error.AppError) {

	u, err := csvc.CartDB.UpdateCart(id, Cart)
	return u, err
}
