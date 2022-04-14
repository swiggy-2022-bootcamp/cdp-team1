package service

import (
	"fmt"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
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
	fetchedCustomer, err := customerService.customerRepository.GetById(customerId)
	if err != nil {
		return nil, err
	}
	return fetchedCustomer, nil
}

func (customerService *CustomerService) GetCustomerByEmail(customerEmail string) (*model.Customer, error) {
	fetchedCustomer, err := customerService.customerRepository.GetByEmail(customerEmail)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return fetchedCustomer, nil
}

func (customerService *CustomerService) UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error) {
	customer.CustomerId = customerId
	updatedCustomer, err := customerService.customerRepository.Update(customer)
	if err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}

func (customerService *CustomerService) DeleteCustomer(customerId string) (*string, error) {
	successMessage, err := customerService.customerRepository.Delete(customerId)
	if err != nil {
		return nil, err
	}
	return successMessage, nil
}
