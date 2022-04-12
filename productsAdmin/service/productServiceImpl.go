package service

import (
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/repository"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (p ProductServiceImpl) CreateProduct(product entity.Product) error {
	if err := p.productRepository.SaveProduct(product); err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) GetAll() ([]entity.Product, error) {
	all, err := p.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (p ProductServiceImpl) UpdateProduct(product entity.Product) error {
	err := p.productRepository.FindAndUpdate(product)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) DeleteProduct(productId string) error {
	err := p.productRepository.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) SearchProduct(query string) ([]entity.Product, error) {
	all, err := p.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}
