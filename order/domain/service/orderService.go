package service

import (
	"fmt"
	"orderService/domain/model"
	"orderService/domain/repository"
	"orderService/internal/error"
)

type OrderService interface {
	CreateOrder(*model.Order) *error.AppError
	GetAllOrders() (*[]model.Order, *error.AppError)
	GetOrderByStatus(string) (*[]model.Order, *error.AppError)
	GetOrderById(string) (*model.Order, *error.AppError)
	GetOrderByCustomerId(string) (*[]model.Order, *error.AppError)
	UpdateOrder(string, string) *error.AppError
	DeleteOrderById(string) *error.AppError
	CreateInvoice(string) *error.AppError
	// DeleteAllOrders() *error.AppError
}

type OrderServiceImpl struct {
	orderRepository repository.OrderRepositoryDB
}

func NewOrderService(orderRepository repository.OrderRepositoryDB) OrderService {
	return &OrderServiceImpl{
		orderRepository: orderRepository,
	}
}

func (odb OrderServiceImpl) CreateOrder(new_order *model.Order) *error.AppError {

	err := odb.orderRepository.Create(new_order)

	if err != nil {
		return err
	}

	return nil
}

func (odb OrderServiceImpl) GetAllOrders() (*[]model.Order, *error.AppError) {

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

func (odb OrderServiceImpl) GetOrderById(order_id string) (*model.Order, *error.AppError) {

	u, err := odb.orderRepository.ReadOrderID(order_id)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb OrderServiceImpl) GetOrderByCustomerId(id string) (*[]model.Order, *error.AppError) {

	u, err := odb.orderRepository.ReadCustomerID(id)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb OrderServiceImpl) UpdateOrder(order_id string, status string) *error.AppError {

	fmt.Println("update service status: ", status)

	err := odb.orderRepository.Update(order_id, status)

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

func (odb OrderServiceImpl) CreateInvoice(order_id string) *error.AppError {

	err := odb.orderRepository.CreateOrderInvoice(order_id)

	if err != nil {
		return err
	}

	return nil
}

// func (odb OrderServiceImpl) DeleteAllOrder() *error.AppError {

// 	err := odb.orderRepository.DeleteAll()

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
