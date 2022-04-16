package repository

import (
	"github.com/stretchr/testify/mock"
	"qwik.in/productsAdmin/entity"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Connect() error {
	args := mock.Called()
	return args.Error(0)

}
func (mock *MockRepository) FindOne(productId string) (entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(entity.Product), args.Error(1)
}
func (mock *MockRepository) FindAll() ([]entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Product), args.Error(1)
}
func (mock *MockRepository) SaveProduct(product entity.Product) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) DeleteProduct(productId string) error {
	args := mock.Called()
	return args.Error(0)
}
func (mock *MockRepository) FindWithLimit(limit int64) ([]entity.Product, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Product), args.Error(1)
}
