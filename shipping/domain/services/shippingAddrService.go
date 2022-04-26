package services

import (
	"qwik.in/shipping/domain/models"
	"qwik.in/shipping/domain/repository"
	"qwik.in/shipping/domain/tools/errs"
)

//ShippingAddrService ..
type ShippingAddrService interface {
	CreateNewShippingAddress(string, string, string, string, string, string, string, string, string, int, string, bool, int) (string, *errs.AppError)
	GetShippingAddressById(string) (*repository.ShippingAddress, *errs.AppError)
	GetDefaultShippingAddr(string) (*repository.ShippingAddress, *errs.AppError)
	GetAllShippingAddressOfUser(string) (*[]models.ShippingAddrModel, *errs.AppError)
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
func (s ShippingAddrServiceImpl) CreateNewShippingAddress(userID string, shippingAddressId, firstName, lastName, addressLine1, addressLine2, city, state, phone string, pincode int, addressType string, defaultAddress bool, shippingCost int) (string, *errs.AppError) {
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
func (s ShippingAddrServiceImpl) GetDefaultShippingAddr(userId string) (*repository.ShippingAddress, *errs.AppError) {
	data, err := s.shippingAddrRepo.FindDefaultShippingAddressImpl(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s ShippingAddrServiceImpl) GetAllShippingAddressOfUser(userId string) (*[]models.ShippingAddrModel, *errs.AppError) {
	data, err := s.shippingAddrRepo.FindAllShippingAddressOfUser(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
