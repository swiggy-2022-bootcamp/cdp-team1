package service

import (
	"orderService/domain/repository"
	"orderService/protos"
)

var orderRepository repository.OrderRepository

type OrderProtoServer struct {
	protos.UnimplementedOrderServer
}

func NewOrderProtoService(pr repository.OrderRepository) OrderProtoServer {
	orderRepository = pr
	return OrderProtoServer{}
}

func (o OrderProtoServer) CreateOrder(ctx context.Context, req *protos.CreateOrderRequest) (*protos.CreateOrderResponse, error) {

	