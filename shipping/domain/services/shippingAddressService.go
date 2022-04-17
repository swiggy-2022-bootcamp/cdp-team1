package services

import (
	"github.com/ashwin2125/qwk/shipping/domain/repository"
	"github.com/ashwin2125/qwk/shipping/domain/tools/errs"
)

//ShippingAddressService ..
type ShippingAddressService interface {
	CreateShippingAddress(int, string, string, string, string, string, string, string, int, string, bool) (string, *errs.AppError)
}

type service1 struct {
	shippingAddressRepository repository.ShippingAddrRepo
}

func (s service1) CreateShippingAddress(userID int, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool) (string, *errs.AppError) {
	shippingAddress := repository.ShippingAddressFunc(userID, firstName, lastName, addressLine1, addressLine2, city, state, phone, pincode, addressType, defaultAddress)
	data, err := s.shippingAddressRepository.InsertShippingAddress(*shippingAddress)
	if err != nil {
		return "", err
	}
	return data, nil
}

//ShippingAddressServiceFunc ..
func ShippingAddressServiceFunc(shippingAddressRepository repository.ShippingAddrRepo) ShippingAddressService {
	return &service1{
		shippingAddressRepository: shippingAddressRepository,
	}
}
