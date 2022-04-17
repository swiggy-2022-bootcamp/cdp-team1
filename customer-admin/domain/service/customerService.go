package service

import (
	"golang.org/x/crypto/bcrypt"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
	"qwik.in/customers-admin/internal/errors"
	"time"
)

type CustomerServiceInterface interface {
	CreateCustomer(customer model.Customer) (*model.Customer, error)
	GetCustomerById(customerId string) (*model.Customer, error)
	GetCustomerByEmail(customerEmail string) (*model.Customer, error)
	UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error)
	DeleteCustomer(customerId string) (*string, error)
}

type CustomerService struct {
	customerRepository repository.CustomerRepositoryInterface
}

func InitCustomerService(repositoryToInject repository.CustomerRepositoryInterface) CustomerServiceInterface {
	customerService := new(CustomerService)
	customerService.customerRepository = repositoryToInject
	return customerService
}

func (customerService *CustomerService) CreateCustomer(customer model.Customer) (*model.Customer, error) {
	//customer with email id already exists
	fetchedCustomer, _ := customerService.GetCustomerByEmail(customer.Email)
	if fetchedCustomer != nil {
		return nil, errors.NewEmailAlreadyRegisteredError()
	}

	//encrypt password
	customerPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	customer.Password = string(customerPassword)
	customer.DateAdded = time.Now()

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
	if err != nil {
		return nil, err
	}
	return fetchedCustomer, nil
}

func (customerService *CustomerService) UpdateCustomer(customerId string, customer model.Customer) (*model.Customer, error) {
	customer.CustomerId = customerId
	
	//encrypt password
	customerPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	customer.Password = string(customerPassword)

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
