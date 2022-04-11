package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qwik.in/customers-admin/domain/model"
)

type CustomerRepositoryInterface interface {
	Create(customer model.Customer) (*model.Customer, error)
	GetById(customerId primitive.ObjectID) (*model.Customer, error)
	GetByEmail(customerEmail string) (*model.Customer, error)
	Update(customer model.Customer) (*model.Customer, error)
	Delete(customerId primitive.ObjectID) (*string, error)
}

type CustomerRepository struct {
}

func (customerRepository *CustomerRepository) Create(customer model.Customer) (*model.Customer, error) {
	return nil, nil
}

func (customerRepository *CustomerRepository) GetById(customerId primitive.ObjectID) (*model.Customer, error) {
	return nil, nil
}

func (customerRepository *CustomerRepository) GetByEmail(customerEmail string) (*model.Customer, error) {
	return nil, nil
}

func (customerRepository *CustomerRepository) Update(customer model.Customer) (*model.Customer, error) {
	return nil, nil
}

func (customerRepository *CustomerRepository) Delete(customerId primitive.ObjectID) (*string, error) {
	return nil, nil
}
