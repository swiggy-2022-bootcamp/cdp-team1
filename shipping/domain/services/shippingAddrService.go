package services

import (
	"github.com/ashwin2125/qwk/shipping/domain/repository"
	"github.com/ashwin2125/qwk/shipping/domain/tools/errs"
)

//ShippingAddrService ..
type ShippingAddrService interface {
	CreateNewShippingAddress(int, string, string, string, string, string, string, string, int, string, bool) (string, *errs.AppError)
}

//ShippingAddrServiceImpl ..
type ShippingAddrServiceImpl struct {
	shippingAddrRepo repository.ShippingAddrRepo
}

//ShippingAddressServiceFunc ..
func ShippingAddressServiceFunc(shippingAddressRepository repository.ShippingAddrRepo) ShippingAddrService {
	return &ShippingAddrServiceImpl{
		shippingAddrRepo: shippingAddressRepository,
	}
}

//CreateNewShippingAddress ..
func (s ShippingAddrServiceImpl) CreateNewShippingAddress(userID int, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool) (string, *errs.AppError) {
	shippingAddress := repository.ShippingAddrFunc(userID, firstName, lastName, addressLine1, addressLine2, city, state, phone, pincode, addressType, defaultAddress)
	data, err := s.shippingAddrRepo.CreateNewShippingAddrImpl(*shippingAddress)
	if err != nil {
		return "", err
	}
	return data, nil
}
