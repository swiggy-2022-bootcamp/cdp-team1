package service

import (
	"orderService/domain/model"
	"orderService/domain/repository"
	"orderService/internal/error"
)

type OrderService interface {
	GetAllOrder() (*[]model.Order, *error.AppError)
	GetOrderByStatus(string) (*[]model.Order, *error.AppError)
	UpdateOrder(model.Order) (*model.Order, *error.AppError)
	DeleteOrderById(string) (*model.Order, *error.AppError)
	DeleteAllOrder() *error.AppError
}

type DefaultOrderService struct {
	OrderDB repository.OrderRepositoryDB
}

func (odb DefaultOrderService) GetAllOrder() (*[]model.Order, *error.AppError) {

	u, err := odb.OrderDB.ReadAll()

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb DefaultOrderService) GetOrderByStatus(status string) (*[]model.Order, *error.AppError) {

	u, err := odb.OrderDB.ReadStatus(status)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb DefaultOrderService) UpdateOrder(order model.Order) (*model.Order, *error.AppError) {

	u, err := odb.OrderDB.Update(order)

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb DefaultOrderService) DeleteOrderById(id string) (*model.Order, *error.AppError) {

	u, err := odb.OrderDB.Delete(model.Order{Id: id})

	if err != nil {
		return nil, err
	}

	return u, nil
}

func (odb DefaultOrderService) DeleteAllOrder() *error.AppError {

	err := odb.OrderDB.DeleteAll()

	if err != nil {
		return err
	}

	return nil
}
