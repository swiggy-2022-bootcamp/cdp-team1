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

/*
1.1 gRPC - Cart : checkout/confirm : Get the items from cart
1.2 Random OrderID creation for Cart Details
1.3 Return User - the orderID, cartDetails

---
2.1 get Shipping Address through gRPC call to ShippingAddress
2.2
*/
