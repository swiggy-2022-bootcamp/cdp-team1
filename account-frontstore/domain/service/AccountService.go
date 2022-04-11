package service

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/repository"
	"qwik.in/account-frontstore/internal/errors"
)

type AccountServiceInterface interface {
	CreateAccount(account model.Account) (*model.Account, error)
	GetAccountById(customerId string) (*model.Account, error)
	UpdateAccount(customerId string, account model.Account) (*model.Account, error)
}

type AccountService struct {
	accountRepository repository.AccountRepository
}

func (accountService *AccountService) CreateAccount(account model.Account) (*model.Account, error) {
	createdAccount, err := accountService.accountRepository.Create(account)
	if err != nil {
		return nil, err
	}
	return createdAccount, nil
}

func (accountService *AccountService) GetAccountById(customerId string) (*model.Account, error) {
	objectId, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return nil, errors.NewMalformedIdError()
	}

	fetchedAccount, err := accountService.accountRepository.GetById(objectId)
	if err != nil {
		return nil, err
	}
	return fetchedAccount, nil
}

func (accountService *AccountService) UpdateAccount(customerId string, account model.Account) (*model.Account, error) {
	objectId, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		return nil, errors.NewMalformedIdError()
	}

	account.CustomerId = objectId
	updatedAccount, err := accountService.accountRepository.Update(account)
	if err != nil {
		return nil, err
	}
	return updatedAccount, nil
}
