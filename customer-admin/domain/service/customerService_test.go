package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository/mocks"
	"qwik.in/customers-admin/internal/errors"
	"testing"
)

func TestShouldThrowNewEmailAlreadyRegisteredError(t *testing.T) {
	customer := model.Customer{}
	customer.CustomerId = "894425b4-9141-41d0-9590-177336c0ca76"
	customer.Email = "repeatedEmail@gmail.com"

	customerRepo := &mocks.CustomerRepositoryInterface{}
	customerService := InitCustomerService(customerRepo)

	customerRepo.On("GetByEmail", "repeatedEmail@gmail.com").Return(&customer, nil)
	createdCustomer, err := customerService.CreateCustomer(customer)

	assert.Nil(t, createdCustomer)
	assert.EqualError(t, err, errors.NewEmailAlreadyRegisteredError().ErrorMessage)
}

func TestShouldCreateNewCustomer(t *testing.T) {
	customer := model.Customer{}
	customer.CustomerId = "894425b4-9141-41d0-9590-177336c0ca76"
	customer.Email = "nonRepeatedEmail@gmail.com"

	customerRepo := &mocks.CustomerRepositoryInterface{}
	customerService := InitCustomerService(customerRepo)

	customerRepo.On("GetByEmail", "nonRepeatedEmail@gmail.com").Return(nil, errors.NewUserNotFoundError())
	customerRepo.On("Create", mock.Anything).Return(&customer, nil)
	createdCustomer, err := customerService.CreateCustomer(customer)

	assert.Nil(t, err)
	assert.Equal(t, createdCustomer.CustomerId, customer.CustomerId)
}

func TestShouldThrowUserNotFoundError(t *testing.T) {
	customer := model.Customer{}
	customer.CustomerId = "894425b4-9141-41d0-9590-177336c0ca76"
	customer.Email = "repeatedEmail@gmail.com"

	customerRepo := &mocks.CustomerRepositoryInterface{}
	customerService := InitCustomerService(customerRepo)

	customerRepo.On("GetByEmail", "repeatedEmail@gmail.com").Return(nil, errors.NewUserNotFoundError())
	customerRepo.On("GetById", "894425b4-9141-41d0-9590-177336c0ca76").Return(nil, errors.NewUserNotFoundError())
	customerRepo.On("Update", customer).Return(nil, errors.NewUserNotFoundError())
	customerRepo.On("Delete", "894425b4-9141-41d0-9590-177336c0ca76").Return(nil, errors.NewUserNotFoundError())

	fetchedCustomer, err := customerService.GetCustomerById(customer.CustomerId)
	assert.Nil(t, fetchedCustomer)
	assert.EqualError(t, err, errors.NewUserNotFoundError().ErrorMessage)

	fetchedCustomer, err = customerService.GetCustomerByEmail(customer.Email)
	assert.Nil(t, fetchedCustomer)
	assert.EqualError(t, err, errors.NewUserNotFoundError().ErrorMessage)

	fetchedCustomer, err = customerService.UpdateCustomer(customer.CustomerId, customer)
	assert.Nil(t, fetchedCustomer)
	assert.EqualError(t, err, errors.NewUserNotFoundError().ErrorMessage)

	successMessage, err := customerService.DeleteCustomer(customer.CustomerId)
	assert.Nil(t, successMessage)
	assert.EqualError(t, err, errors.NewUserNotFoundError().ErrorMessage)
}
