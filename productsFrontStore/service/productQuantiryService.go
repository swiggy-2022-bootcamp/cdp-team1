package service

import (
	"context"
	"qwik.in/productsFrontStore/log"
	"qwik.in/productsFrontStore/proto"
	"qwik.in/productsFrontStore/repository"
	"strconv"
)

type ProductQualityService struct {
	proto.UnimplementedQuantityServiceServer
}

var (
	Repository repository.ProductRepository
	Service    ProductService
)

func NewProductQualityService() *ProductQualityService {
	Repository = repository.NewDynamoRepository()
	err := Repository.Connect()
	if err != nil {
		log.Error("Error while Connecting to DB: ", err)
		return nil
	}
	Service = NewProductService(Repository)
	return &ProductQualityService{}
}

func (s *ProductQualityService) GetQuantity(ctx context.Context, in *proto.Request) (*proto.Response, error) {
	log.Info("gRPC received message: ", in.GetId())
	product, err1 := Service.GetProductById(in.GetId())
	if err1 != nil {
		return nil, err1
	}
	productQuantity, err2 := strconv.Atoi(product.Quantity)
	if err2 != nil {
		return nil, err2
	}
	return &proto.Response{
		Id:       in.GetId(),
		Quantity: int32(productQuantity),
	}, nil
}
