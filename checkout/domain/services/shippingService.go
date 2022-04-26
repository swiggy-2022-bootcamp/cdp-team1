package services

import (
	_ "qwik.in/checkout/docs"
	"qwik.in/checkout/domain/repository"
	"qwik.in/checkout/domain/tools/errs"
)

//ShippingService ..
type ShippingService interface {
	GetDefaultShippingAddr(bool) (*repository.ShippingAddress, *errs.AppError)
}

//ShippingServiceImpl ..
type ShippingServiceImpl struct {
	shippingRepo repository.ShippingRepo
}

//ShippingAddressServiceFunc ..
func ShippingAddressServiceFunc(shippingAddressRepository repository.ShippingRepo) ShippingService {
	return &ShippingServiceImpl{
		shippingRepo: shippingAddressRepository,
	}
}

//GetDefaultShippingAddr ..
func (s ShippingServiceImpl) GetDefaultShippingAddr(isDefaultAddress bool) (*repository.ShippingAddress, *errs.AppError) {
	data, err := s.shippingRepo.FindDefaultShippingAddressImpl(isDefaultAddress)
	if err != nil {
		return nil, err
	}
	return data, nil
}
