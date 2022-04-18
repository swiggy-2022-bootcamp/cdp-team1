package service

import (
	"orderService/domain/model"
	"orderService/domain/repository"
	"orderService/internal/error"
)

type OrderService interface {
	GetAllOrder() (*[]model.Order, *error.AppError)
	GetOrderByStatus(string) (*[]model.Order, *error.AppError)
	UpdateOrder(model.Order) *error.AppError
	DeleteOrderById(string) *error.AppError
	DeleteAllOrder() *error.AppError
}

type OrderServiceImpl struct {
	orderRepository repository.OrderRepositoryDB
}

func NewOrderService(orderRepository repository.OrderRepositoryDB) OrderService {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
	}
}

func (odb OrderServiceImpl) GetAllOrder() (*[]model.Order, *error.AppError) {

	u, err := odb.orderRepository.ReadAll()

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb OrderServiceImpl) GetOrderByStatus(status string) (*[]model.Order, *error.AppError) {

	u, err := odb.orderRepository.ReadStatus(status)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb OrderServiceImpl) UpdateOrder(order model.Order) *error.AppError {

	err := odb.orderRepository.Update(order)

	if err != nil {
		return err
	}

	return nil
}

func (odb OrderServiceImpl) DeleteOrderById(id string) *error.AppError {

	err := odb.orderRepository.Delete(model.Order{OrderId: id})

	if err != nil {
		return err
	}

	return nil
}

func (odb OrderServiceImpl) DeleteAllOrder() *error.AppError {

	err := odb.orderRepository.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
