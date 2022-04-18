package service

import (
	"github.com/stretchr/testify/mock"
	"qwik.in/productsAdmin/app/proto"
	"qwik.in/productsAdmin/entity"
)

type MockProductService struct {
	mock.Mock
}

func NewMockProductService() ProductService {
	return &MockProductService{}
}

func (m MockProductService) CreateProduct(product entity.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m MockProductService) GetAll() ([]entity.Product, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]entity.Product), args.Error(1)
}

func (m MockProductService) UpdateProduct(productId string, product entity.Product) error {
	args := m.Called(productId, product)
	return args.Error(0)
}

func (m MockProductService) DeleteProduct(productId string) error {
	args := m.Called(productId)
	return args.Error(0)
}

func (m MockProductService) SearchProduct(limit int64) ([]entity.Product, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]entity.Product), args.Error(1)
}

func (m MockProductService) GetQuantityForProductId(productId string) (*proto.Response, error) {
	args := m.Called(productId)
	result := args.Get(0)
	return result.(*proto.Response), args.Error(1)
}
