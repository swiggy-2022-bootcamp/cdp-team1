// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	model "qwik.in/account-frontstore/domain/model"
)

// AccountRepositoryInterface is an autogenerated mock type for the AccountRepositoryInterface type
type AccountRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: account
func (_m *AccountRepositoryInterface) Create(account model.Account) (*model.Account, error) {
	ret := _m.Called(account)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(model.Account) *model.Account); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Account) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: customerEmail
func (_m *AccountRepositoryInterface) GetByEmail(customerEmail string) (*model.Account, error) {
	ret := _m.Called(customerEmail)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(string) *model.Account); ok {
		r0 = rf(customerEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(customerEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetById provides a mock function with given fields: customerId
func (_m *AccountRepositoryInterface) GetById(customerId string) (*model.Account, error) {
	ret := _m.Called(customerId)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(string) *model.Account); ok {
		r0 = rf(customerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(customerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: account
func (_m *AccountRepositoryInterface) Update(account model.Account) (*model.Account, error) {
	ret := _m.Called(account)

	var r0 *model.Account
	if rf, ok := ret.Get(0).(func(model.Account) *model.Account); ok {
		r0 = rf(account)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Account)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Account) error); ok {
		r1 = rf(account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
