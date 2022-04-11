package domain

import (
	"cartService/domain/db"
	"cartService/domain/repository"
	error "cartService/internal"
)

type CartService interface {
	AddToCart(db.Cart) (*db.Cart, *error.AppError)
	GetAllCart() (*[]db.Cart, *error.AppError)
	UpdateCart(db.Cart) (*db.Cart, *error.AppError)
	DeleteCartById(string) (*db.Cart, *error.AppError)
	DeleteAllCart() *error.AppError
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

func (csvc DefaultCartService) GetAllCart() (*[]db.Cart, *error.AppError) {

	u, err := csvc.CartDB.ReadAll()

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc DefaultCartService) UpdateCart(id string, Cart db.Cart) (*db.Cart, *error.AppError) {

	u, err := csvc.CartDB.Update(db.Cart{})

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc DefaultCartService) DeleteCartById(id string) (*db.Cart, *error.AppError) {

	u, err := csvc.CartDB.Delete(db.Cart{})

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc DefaultCartService) DeleteAllCart() *error.AppError {

	err := csvc.CartDB.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
