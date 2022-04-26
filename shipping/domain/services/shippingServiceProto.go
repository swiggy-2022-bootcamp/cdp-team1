package services

import (
	"context"
	"fmt"
	"qwik.in/shipping/domain/repository"
	"qwik.in/shipping/domain/tools/logger"
	"qwik.in/shipping/protos"
)

var shippingAddrRepo repository.ShippingAddrRepo

type ShippingProtoServer struct {
	protos.UnimplementedShippingAddressProtoFuncServer
}

func NewShippingRepoService(pr repository.ShippingAddrRepo) ShippingProtoServer {
	shippingAddrRepo = pr
	return ShippingProtoServer{}
}

func (s ShippingProtoServer) GetDefaultShippingAddress(ctx context.Context, shippingAddressRequest *protos.ShippingAddressRequest) (*protos.ShippingAddressResponse, error) {
	response := &protos.ShippingAddressResponse{}
	//Fetch Default Shipping Address
	userShippingAddress, err := shippingAddrRepo.FindDefaultShippingAddressImpl(int(shippingAddressRequest.GetUserId()))
	if err != nil {
		logger.Error("Error in Proto GetDefaultShippingAddress", err)
		return nil, err.Error()
	}
	//Convert DynamoDB Object into Proto Message Object
	var shippingAddressProto *protos.ShippingAddress
	userShippingAddress1 := userShippingAddress
	shippingAddressProto = &protos.ShippingAddress{
		UserId:            int32(int(userShippingAddress1.UserID)),
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
	response.ShippingAddress = shippingAddressProto
	return response, nil
}

func (s ShippingProtoServer) GetAllShippingAddress(ctx context.Context, shippingAddressRequest *protos.ShippingAddressRequest) (*protos.ShippingAddressResponse, error) {
	//shippingAddr := make([]*protos.ShippingAddress, 0, 15)
	//response := &protos.AllShippingAddressResponse{
	//	ShippingAddress: shippingAddr,
	//}
	//Fetch Default Shipping Address
	userShippingAddress, err := shippingAddrRepo.FindAllShippingAddressOfUser(int(shippingAddressRequest.GetUserId()))
	if err != nil {
		logger.Error("Error in Proto GetDefaultShippingAddress", err)
		return nil, err.Error()
	}

	fmt.Println(userShippingAddress)
	////Convert DynamoDB Object into Proto Message Object
	//var shippingAddressProto *protos.ShippingAddress
	//for _, data := range response.ShippingAddress {
	//	shippingAddressProto = &protos.ShippingAddress{
	//		UserId:            data.UserId,
	//		ShippingAddressId: data.ShippingAddressId,
	//		FirstName:         data.FirstName,
	//		LastName:          data.LastName,
	//		AddressLine_1:     data.AddressLine_1,
	//		AddressLine_2:     data.AddressLine_2,
	//		City:              data.City,
	//		State:             data.State,
	//		Phone:             data.Phone,
	//		Pincode:           data.Pincode,
	//		AddressType:       data.AddressType,
	//		DefaultAddress:    data.DefaultAddress,
	//		ShippingCost:      data.ShippingCost,
	//	}
	//	response.ShippingAddress = append(shippingAddr, shippingAddressProto)
	//}
	//fmt.Println(response.ShippingAddress)
	return nil, nil
	//shippingAddressProto = &protos.ShippingAddress{
	//	UserId:            int32(int(userShippingAddress1[0].UserID)),
	//	ShippingAddressId: userShippingAddress1.ShippingAddressID,
	//	FirstName:         userShippingAddress1.FirstName,
	//	LastName:          userShippingAddress1.LastName,
	//	AddressLine_1:     userShippingAddress1.AddressLine1,
	//	AddressLine_2:     userShippingAddress1.AddressLine2,
	//	City:              userShippingAddress1.City,
	//	State:             userShippingAddress1.State,
	//	Phone:             userShippingAddress1.Phone,
	//	Pincode:           float64(userShippingAddress1.Pincode),
	//	AddressType:       userShippingAddress1.AddressType,
	//	DefaultAddress:    userShippingAddress1.DefaultAddress,
	//	ShippingCost:      float64(userShippingAddress1.ShippingCost),
	//}
	//response.ShippingAddress = shippingAddressProto
	//return response, nil
}
