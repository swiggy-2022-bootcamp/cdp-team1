package services

import (
	"github.com/ashwin2125/qwk/shipping/domain/repository"
	"github.com/ashwin2125/qwk/shipping/domain/tools/errs"
)

//ShippingAddrService ..
type ShippingAddrService interface {
	CreateNewShippingAddress(int, string, string, string, string, string, string, string, string, int, string, bool, int) (string, *errs.AppError)
	GetShippingAddressById(string) (*repository.ShippingAddress, *errs.AppError)
	GetDefaultShippingAddr(bool) (*repository.ShippingAddress, *errs.AppError)
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
func (s ShippingAddrServiceImpl) CreateNewShippingAddress(userID int, shippingAddressId, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool, shippingCost int) (string, *errs.AppError) {
	shippingAddress := repository.ShippingAddrFunc(userID, shippingAddressId, firstName, lastName, addressLine1, addressLine2, city, state, phone, pincode, addressType, defaultAddress, shippingCost)
	data, err := s.shippingAddrRepo.CreateNewShippingAddrImpl(*shippingAddress)
	if err != nil {
		return "", err
	}
	return data, nil
}

//GetShippingAddressById ..
func (s ShippingAddrServiceImpl) GetShippingAddressById(shippingAddressId string) (*repository.ShippingAddress, *errs.AppError) {
	res, err := s.shippingAddrRepo.FindShippingAddressByIdImpl(shippingAddressId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//GetDefaultShippingAddr ..
func (s ShippingAddrServiceImpl) GetDefaultShippingAddr(isDefaultAddress bool) (*repository.ShippingAddress, *errs.AppError) {
	data, err := s.shippingAddrRepo.FindDefaultShippingAddressImpl(isDefaultAddress)
	if err != nil {
		return nil, err
	}
	return data, nil
}
