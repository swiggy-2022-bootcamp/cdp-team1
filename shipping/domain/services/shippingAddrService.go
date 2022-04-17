package services

import (
	"github.com/ashwin2125/qwk/shipping/domain/repository"
	"github.com/ashwin2125/qwk/shipping/domain/tools/errs"
)

//ShippingAddrService ..
type ShippingAddrService interface {
	CreateNewShippingAddress(int, string, string, string, string, string, string, string, string, int, string, bool) (string, *errs.AppError)
	GetShippingAddressById(string) (*repository.ShippingAddress, *errs.AppError)
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
func (s ShippingAddrServiceImpl) CreateNewShippingAddress(userID int, shippingAddressId, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool) (string, *errs.AppError) {
	shippingAddress := repository.ShippingAddrFunc(userID, shippingAddressId, firstName, lastName, addressLine1, addressLine2, city, state, phone, pincode, addressType, defaultAddress)
	data, err := s.shippingAddrRepo.CreateNewShippingAddrImpl(*shippingAddress)
	if err != nil {
		return "", err
	}
	return data, nil
}

func (s ShippingAddrServiceImpl) GetShippingAddressById(shippingAddressId string) (*repository.ShippingAddress, *errs.AppError) {
	res, err := s.shippingAddrRepo.FindShippingAddressByIdImpl(shippingAddressId)
	if err != nil {
		return nil, err
	}
	return res, nil
}

//
//func (s ShippingAddrServiceImpl) DeleteShippingAddressById(shippingAddressId string) (bool, *errs.AppError) {
//	_, err := s.shippingAddrRepo.DeleteShippingAddressById(shippingAddressId)
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}
//
//func (s ShippingAddrServiceImpl) UpdateShippingAddressById(shippingAddressId string, newShippingAddress repository.ShippingAddress) (bool, *errs.AppError) {
//	_, err := s.shippingAddrRepo.UpdateShippingAddressById(shippingAddressId, newShippingAddress)
//	if err != nil {
//		return false, err
//	}
//	return true, nil
//}
