package repository

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"qwik.in/account-frontstore/domain/model"
)

type AccountRepositoryInterface interface {
	Create(account model.Account) (*model.Account, error)
	GetById(customerId primitive.ObjectID) (*model.Account, error)
	Update(account model.Account) (*model.Account, error)
}

type AccountRepository struct {
}

func (accountRepository *AccountRepository) Create(account model.Account) (*model.Account, error) {
	return nil, nil
}

func (accountRepository *AccountRepository) GetById(customerId primitive.ObjectID) (*model.Account, error) {
	return nil, nil
}

func (accountRepository *AccountRepository) Update(account model.Account) (*model.Account, error) {
	return nil, nil
}
