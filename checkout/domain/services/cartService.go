package services

import (
	"qwik.in/checkout/domain/repository"
	"qwik.in/checkout/domain/tools/errs"
)

//CartService ..
type CartService interface {
	GetCartDetails() (*repository.Cart, *errs.AppError)
}

//CartServiceImpl ..
type CartServiceImpl struct {
	cartRepo repository.CartRepo
}

//CartServiceFunc ..
func CartServiceFunc(cartRepo repository.CartRepo) CartService {
	return &CartServiceImpl{
		cartRepo: cartRepo,
	}
}

//GetCartDetails ..
func (c CartServiceImpl) GetCartDetails() (*repository.Cart, *errs.AppError) {
	data, err := c.cartRepo.GetCartDetailsImpl()
	if err != nil {
		return nil, nil
	}
	return data, nil
}
