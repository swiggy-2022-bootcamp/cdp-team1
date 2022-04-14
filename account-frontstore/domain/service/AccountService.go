package service

import (
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/repository"
	"time"
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
	account.DateAdded = time.Now()
	createdAccount, err := accountService.accountRepository.Create(account)
	if err != nil {
		return nil, err
	}
	return createdAccount, nil
}

func (accountService *AccountService) GetAccountById(customerId string) (*model.Account, error) {
	fetchedAccount, err := accountService.accountRepository.GetById(customerId)
	if err != nil {
		return nil, err
	}
	return fetchedAccount, nil
}

func (accountService *AccountService) UpdateAccount(customerId string, account model.Account) (*model.Account, error) {
	account.CustomerId = customerId
	updatedAccount, err := accountService.accountRepository.Update(account)
	if err != nil {
		return nil, err
	}
	return updatedAccount, nil
}
