package service

import (
	"cartService/domain/model"
	"cartService/domain/repository"
	"cartService/internal/error"
)

type CartService interface {
	AddToCart(*model.Cart) *error.AppError
	GetAllCart() (*[]model.Cart, *error.AppError)
	UpdateCart(string, string) *error.AppError
	DeleteCartByCustomerId(string) *error.AppError
	DeleteAllCart() *error.AppError
}

type CartServiceImpl struct {
	cartRepository repository.CartRepositoryDB
}

func NewCartService(cartRepository repository.CartRepositoryDB) CartService {
	return &CartServiceImpl{
		cartRepository: cartRepository,
	}
}

func (csvc CartServiceImpl) AddToCart(cart *model.Cart) *error.AppError {

	err := csvc.cartRepository.Create(cart)

	if err != nil {
		return err
	}

	return err
}

func (csvc CartServiceImpl) GetAllCart() (*[]model.Cart, *error.AppError) {

	u, err := csvc.cartRepository.ReadAll()

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc CartServiceImpl) UpdateCart(customer_id string, quantity string) *error.AppError {

	err := csvc.cartRepository.Update(customer_id, quantity)

	if err != nil {
		return err
	}

	return err
}

func (csvc CartServiceImpl) DeleteCartByCustomerId(customer_id string) *error.AppError {

	err := csvc.cartRepository.Delete(customer_id)

	if err != nil {
		return err
	}

	return err
}

func (csvc CartServiceImpl) DeleteAllCart() *error.AppError {

	err := csvc.cartRepository.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
