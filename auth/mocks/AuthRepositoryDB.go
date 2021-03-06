// Code generated by mockery v2.10.2. DO NOT EDIT.

package mocks

import (
	domain "authService/domain"
	errs "authService/errs"

	mock "github.com/stretchr/testify/mock"
)

// AuthRepositoryDB is an autogenerated mock type for the AuthRepositoryDB type
type AuthRepositoryDB struct {
	mock.Mock
}

// FetchSessionByTokenSign provides a mock function with given fields: _a0
func (_m *AuthRepositoryDB) FetchSessionByTokenSign(_a0 string) (*domain.Session, *errs.AppError) {
	ret := _m.Called(_a0)

	var r0 *domain.Session
	if rf, ok := ret.Get(0).(func(string) *domain.Session); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Session)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// RepoHealthCheck provides a mock function with given fields:
func (_m *AuthRepositoryDB) RepoHealthCheck() *errs.AppError {
	ret := _m.Called()

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func() *errs.AppError); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// SaveSession provides a mock function with given fields: _a0
func (_m *AuthRepositoryDB) SaveSession(_a0 domain.Session) *errs.AppError {
	ret := _m.Called(_a0)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(domain.Session) *errs.AppError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// UpdateSessionStatusByTokenSign provides a mock function with given fields: _a0
func (_m *AuthRepositoryDB) UpdateSessionStatusByTokenSign(_a0 string) *errs.AppError {
	ret := _m.Called(_a0)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(string) *errs.AppError); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}
