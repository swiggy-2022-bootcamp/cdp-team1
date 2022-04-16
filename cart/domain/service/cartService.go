package service

import (
	"cartService/domain/model"
	"cartService/domain/repository"
	"cartService/internal/error"
)

type CartService interface {
	AddToCart(model.Cart) *error.AppError
	GetAllCart() (*[]model.Cart, *error.AppError)
	UpdateCart(model.Cart) (*model.Cart, *error.AppError)
	DeleteCartById(string) (*model.Cart, *error.AppError)
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

func (csvc CartServiceImpl) AddToCart(cart model.Cart) *error.AppError {

	_, err := csvc.cartRepository.Create(cart)

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

func (csvc CartServiceImpl) UpdateCart(Cart model.Cart) (*model.Cart, *error.AppError) {

	u, err := csvc.cartRepository.Update(model.Cart{})

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc CartServiceImpl) DeleteCartById(id string) (*model.Cart, *error.AppError) {

	u, err := csvc.cartRepository.Delete(model.Cart{})

	if err != nil {
		return nil, err
	}

	return u, err
}

func (csvc CartServiceImpl) DeleteAllCart() *error.AppError {

	err := csvc.cartRepository.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
