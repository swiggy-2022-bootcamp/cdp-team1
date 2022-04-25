package services

import (
	"context"
	"qwik.in/shipping/domain/repository"
	"qwik.in/shipping/domain/tools/logger"
	"qwik.in/shipping/protos"
)

var shippingAddrRepo repository.ShippingAddrRepo

type ShippingProtoServer struct {
	protos.UnsafeShippingAddressProtoFuncServer
}

func NewShippingRepoService(pr repository.ShippingAddrRepo) ShippingProtoServer {
	shippingAddrRepo = pr
	return ShippingProtoServer{}
}

func (s ShippingProtoServer) GetDefaultShippingAddress(ctx context.Context, shippingAddressRequest *protos.ShippingAddressRequest) (*protos.ShippingAddressResponse, error) {
	shippingAddress := make([]*protos.ShippingAddress, 0, 0)
	response := &protos.ShippingAddressResponse{
		ShippingAddress: shippingAddress[0],
	}

	//Fetch Default Shipping Address
	userShippingAddress, err := shippingAddrRepo.FindDefaultShippingAddressImpl(true)
	if err != nil {
		logger.Error("Error in Proto GetDefaultShippingAddress", err)
		return nil, err.Error()
	}

	//Convert DynamoDB Object into Proto Message Object
	var shippingAddressProto *protos.ShippingAddress
	userShippingAddress1 := userShippingAddress
	shippingAddressProto = &protos.ShippingAddress{
		UserId:            float64(userShippingAddress1.UserID),
		ShippingAddressId: userShippingAddress1.ShippingAddressID,
		FirstName:         userShippingAddress1.FirstName,
		LastName:          userShippingAddress1.LastName,
		AddressLine_1:     userShippingAddress1.AddressLine1,
		AddressLine_2:     userShippingAddress1.AddressLine2,
		City:              userShippingAddress1.City,
		State:             userShippingAddress1.State,
		Phone:             userShippingAddress1.Phone,
		Pincode:           float64(userShippingAddress1.Pincode),
		AddressType:       userShippingAddress1.AddressType,
		DefaultAddress:    userShippingAddress1.DefaultAddress,
		ShippingCost:      float64(userShippingAddress1.ShippingCost),
	}
	shippingAddress[0] = shippingAddressProto
	response.ShippingAddress = shippingAddress[0]
	return response, nil
}
