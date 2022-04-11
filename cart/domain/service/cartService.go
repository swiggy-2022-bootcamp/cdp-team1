package domain

import (
	"cartService/domain/model"
	"cartService/domain/repository"
	"cartService/internal/error"
)

type CartService interface {
	AddToCart(model.Cart) (*model.Cart, *error.AppError)
	GetAllCart() (*[]model.Cart, *error.AppError)
	UpdateCart(model.Cart) (*model.Cart, *error.AppError)
	DeleteCartById(string) (*model.Cart, *error.AppError)
	DeleteAllCart() *error.AppError
}

type DefaultCartService struct {
	CartDB repository.CartRepositoryDB
}

func (csvc DefaultCartService) AddToCart(cart model.Cart) (*model.Cart, *error.AppError) {

	u, err := csvc.CartDB.Create(cart)

	if err != nil {
		return nil, err
	}
	return u, err
}

func (csvc DefaultCartService) GetAllCart() (*[]model.Cart, *error.AppError) {

	u, err := csvc.CartDB.ReadAll()

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc DefaultCartService) UpdateCart(id string, Cart model.Cart) (*model.Cart, *error.AppError) {

	u, err := csvc.CartDB.Update(model.Cart{})

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc DefaultCartService) DeleteCartById(id string) (*model.Cart, *error.AppError) {

	u, err := csvc.CartDB.Delete(model.Cart{})

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
