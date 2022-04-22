package services

import (
	"qwik.in/checkout/domain/repository"
)

//CheckoutService ..
type CheckoutService interface {
}

//CheckoutServiceImpl ..
type CheckoutServiceImpl struct {
	checkoutRepo repository.CheckoutRepo
}

//CheckoutServiceFunc ..
func CheckoutServiceFunc(checkoutRepo repository.CheckoutRepo) CheckoutService {
	return &CheckoutServiceImpl{
		checkoutRepo: checkoutRepo,
	}
}
