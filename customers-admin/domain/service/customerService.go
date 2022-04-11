package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
	"qwik.in/customers-admin/internal/errors"
)

type CustomerServiceInterface interface {
	CreateCustomer(customer model.Customer) (*model.Customer, error)
	GetCustomerById(customerId string) (*model.Customer, error)
	GetCustomerByEmail(customerEmail string) (*model.Customer, error)
	UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error)
	DeleteCustomer(customerId string) (*string, error)
}

type CustomerService struct {
	customerRepository repository.CustomerRepository
}

func (customerService *CustomerService) CreateCustomer(customer model.Customer) (*model.Customer, error) {
	createdCustomer, err := customerService.customerRepository.Create(customer)
	if err != nil {
		return nil, err
	}
	return createdCustomer, nil
}

func (customerService *CustomerService) GetCustomerById(customerId string) (*model.Customer, error) {
	objectId, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return nil, errors.NewMalformedIdError()
	}

	fetchedCustomer, err := customerService.customerRepository.GetById(objectId)
	if err != nil {
		return nil, err
	}
	return fetchedCustomer, nil
}

func (customerService *CustomerService) GetCustomerByEmail(customerEmail string) (*model.Customer, error) {
	fetchedCustomer, err := customerService.customerRepository.GetByEmail(customerEmail)
	if err != nil {
		return nil, err
	}
	return fetchedCustomer, nil
}

func (customerService *CustomerService) UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error) {
	objectId, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return nil, errors.NewMalformedIdError()
	}

	customer.CustomerId = objectId
	updatedCustomer, err := customerService.customerRepository.Update(customer)
	if err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}

func (customerService *CustomerService) DeleteCustomer(customerId string) (*string, error) {
	objectId, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return nil, errors.NewMalformedIdError()
	}

	successMessage, err := customerService.customerRepository.Delete(objectId)
	if err != nil {
		return nil, err
	}
	return successMessage, nil
}


