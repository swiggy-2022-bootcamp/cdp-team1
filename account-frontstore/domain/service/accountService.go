package service

import (
	"golang.org/x/crypto/bcrypt"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/repository"
	"qwik.in/account-frontstore/internal/errors"
	"time"
)

type AccountServiceInterface interface {
	CreateAccount(account model.Account) (*model.Account, error)
	GetAccountById(customerId string) (*model.Account, error)
	UpdateAccount(customerId string, account model.Account) (*model.Account, error)
}

type AccountService struct {
	accountRepository    repository.AccountRepositoryInterface
	grpcClientRepository repository.GrpcClientRepositoryInterface
}

func InitAccountService(accountRepositoryToInject repository.AccountRepositoryInterface, grpcClientRepositoryToInject repository.GrpcClientRepositoryInterface) AccountServiceInterface {
	accountService := new(AccountService)
	accountService.accountRepository = accountRepositoryToInject
	accountService.grpcClientRepository = grpcClientRepositoryToInject
	return accountService
}

func (accountService *AccountService) CreateAccount(account model.Account) (*model.Account, error) {
	fetchedAccount, _ := accountService.accountRepository.GetByEmail(account.Email)
	if fetchedAccount != nil {
		return nil, errors.NewEmailAlreadyRegisteredError()
	}

	//encrypt password
	accountPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	account.Password = string(accountPassword)
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
	fetchedAccount.RewardsTotal = accountService.grpcClientRepository.GetTransactionRewardPointsByCustomerId(customerId)
	fetchedAccount.UserBalance = accountService.grpcClientRepository.GetPaymentMethodsByCustomerId(customerId)

	fetchedAccount.Cart, _ = accountService.grpcClientRepository.GetCartByCustomerId("2")

	rewardsTotal, _ := accountService.grpcClientRepository.GetRewardPointsByCustomerId(customerId)
	fetchedAccount.RewardsTotal = *rewardsTotal

	return fetchedAccount, nil
}

func (accountService *AccountService) UpdateAccount(customerId string, account model.Account) (*model.Account, error) {
	account.CustomerId = customerId

	//encrypt password
	accountPassword, err := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	account.Password = string(accountPassword)

	updatedAccount, err := accountService.accountRepository.Update(account)
	if err != nil {
		return nil, err
	}
	return updatedAccount, nil
}
