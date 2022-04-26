package service

import (
	"qwik.in/productsFrontStore/entity"
	"qwik.in/productsFrontStore/repository"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (p ProductServiceImpl) GetAll() ([]entity.Product, error) {
	products, err := p.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (p ProductServiceImpl) GetProductById(productId string) (entity.Product, error) {
	product, err := p.productRepository.FindById(productId)
	if err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func (p ProductServiceImpl) GetProductsByCategory(categoryId string) ([]entity.Product, error) {
	products, err := p.productRepository.FindByCategory(categoryId)
	if err != nil {
		return []entity.Product{}, err
	}
	return products, nil
}
