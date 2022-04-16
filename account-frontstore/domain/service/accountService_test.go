package service

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/repository/mocks"
	"qwik.in/account-frontstore/internal/errors"
	"testing"
)

func TestShouldThrowNewEmailAlreadyRegisteredError(t *testing.T) {
	account := model.Account{}
	account.CustomerId = "894425b4-9141-41d0-9590-177336c0ca76"
	account.Email = "repeatedEmail@gmail.com"

	accountRepo := &mocks.AccountRepositoryInterface{}
	accountService := InitAccountService(accountRepo)

	accountRepo.On("GetByEmail", "repeatedEmail@gmail.com").Return(&account, nil)
	createdAccount, err := accountService.CreateAccount(account)

	assert.Nil(t, createdAccount)
	assert.EqualError(t, err, errors.NewEmailAlreadyRegisteredError().ErrorMessage)
}

func TestShouldCreateNewCustomer(t *testing.T) {
	account := model.Account{}
	account.CustomerId = "894425b4-9141-41d0-9590-177336c0ca76"
	account.Email = "nonRepeatedEmail@gmail.com"

	accountRepo := &mocks.AccountRepositoryInterface{}
	accountService := InitAccountService(accountRepo)

	accountRepo.On("GetByEmail", "nonRepeatedEmail@gmail.com").Return(nil, errors.NewUserNotFoundError())
	accountRepo.On("Create", mock.Anything).Return(&account, nil)
	createdAccount, err := accountService.CreateAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, createdAccount.CustomerId, createdAccount.CustomerId)
}

func TestShouldThrowUserNotFoundError(t *testing.T) {
	account := model.Account{}
	account.CustomerId = "894425b4-9141-41d0-9590-177336c0ca76"
	account.Email = "repeatedEmail@gmail.com"

	accountRepo := &mocks.AccountRepositoryInterface{}
	accountService := InitAccountService(accountRepo)

	accountRepo.On("GetByEmail", "repeatedEmail@gmail.com").Return(nil, errors.NewUserNotFoundError())
	accountRepo.On("GetById", "894425b4-9141-41d0-9590-177336c0ca76").Return(nil, errors.NewUserNotFoundError())
	accountRepo.On("Update", account).Return(nil, errors.NewUserNotFoundError())

	fetchedAccount, err := accountService.GetAccountById(account.CustomerId)
	assert.Nil(t, fetchedAccount)
	assert.EqualError(t, err, errors.NewUserNotFoundError().ErrorMessage)

	fetchedAccount, err = accountService.UpdateAccount(account.CustomerId, account)
	assert.Nil(t, fetchedAccount)
	assert.EqualError(t, err, errors.NewUserNotFoundError().ErrorMessage)
}
