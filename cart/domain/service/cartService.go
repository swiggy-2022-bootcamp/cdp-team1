package domain

import (
	"cartService/domain/db"
	"cartService/domain/repository"
	error "cartService/internal"
)

type CartService interface {
	AddToCart(db.Cart) (*db.Cart, *error.AppError)
	UpdateCart(db.Cart) (*db.Cart, *error.AppError)
}

type DefaultCartService struct {
	CartDB repository.CartRepositoryDB
}

func (csvc DefaultCartService) AddToCart(cart db.Cart) (*db.Cart, *error.AppError) {

	u, err := csvc.CartDB.Create(cart)

	if err != nil {
		return nil, err
	}
	return u, err
}

func (csvc DefaultCartService) UpdateCart(id string, Cart db.Cart) (*db.Cart, *error.AppError) {
	u, err := csvc.CartDB.Update(db.Cart{})
	return u, err
}
