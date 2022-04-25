package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/repository"
	"testing"
)

func TestCreateProductSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("SaveProduct").Return(nil)
	productService := NewProductService(mockRepo)

	product := entity.Product{}
	result := productService.CreateProduct(product)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, result)
}

func TestCreateProductFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("SaveProduct").Return(errors.New("Cannot save product"))
	productService := NewProductService(mockRepo)

	product := entity.Product{}
	result := productService.CreateProduct(product)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, result)
}

func TestGetAllSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	product := entity.Product{}
	mockRepo.On("FindAll").Return([]entity.Product{product}, nil)
	productService := NewProductService(mockRepo)

	products, err := productService.GetAll()

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, len(products), 1)
}

func TestGetAllFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("FindAll").Return([]entity.Product{}, errors.New("Cannot find products"))
	productService := NewProductService(mockRepo)

	products, err := productService.GetAll()

	mockRepo.AssertExpectations(t)
	assert.Nil(t, products)
	assert.NotNil(t, err)
}

func TestUpdateProductSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	product := entity.Product{}
	mockRepo.On("SaveProduct").Return(nil)
	productService := NewProductService(mockRepo)

	err := productService.UpdateProduct(entity.NewID(), product)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestUpdateProductFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	product := entity.Product{}
	mockRepo.On("SaveProduct").Return(errors.New("Cannot save product"))
	productService := NewProductService(mockRepo)

	err := productService.UpdateProduct(entity.NewID(), product)

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, err)
}

func TestDeleteProductSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("DeleteProduct").Return(nil)
	productService := NewProductService(mockRepo)

	err := productService.DeleteProduct(entity.NewID())

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
}

func TestDeleteProductFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("DeleteProduct").Return(errors.New("Cannot delete product"))
	productService := NewProductService(mockRepo)

	err := productService.DeleteProduct(entity.NewID())

	mockRepo.AssertExpectations(t)
	assert.NotNil(t, err)
}

func TestSearchProductSuccess(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	product := entity.Product{}
	mockRepo.On("FindWithLimit").Return([]entity.Product{product}, nil)
	productService := NewProductService(mockRepo)

	products, err := productService.SearchProduct(1)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, err)
	assert.NotNil(t, products)
	assert.Equal(t, len(products), 1)
}

func TestSearchProductFailure(t *testing.T) {
	mockRepo := new(repository.MockRepository)
	mockRepo.On("FindWithLimit").Return([]entity.Product{}, errors.New("Cannot find products"))
	productService := NewProductService(mockRepo)

	products, err := productService.SearchProduct(1)

	mockRepo.AssertExpectations(t)
	assert.Nil(t, products)
	assert.NotNil(t, err)
}
