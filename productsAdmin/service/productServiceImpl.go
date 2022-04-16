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
	product.SetId()
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

func (p ProductServiceImpl) UpdateProduct(productId string, product entity.Product) error {
	product.ID = productId
	if err := p.productRepository.SaveProduct(product); err != nil {
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

func (p ProductServiceImpl) SearchProduct(limit int64) ([]entity.Product, error) {
	all, err := p.productRepository.FindWithLimit(limit)
	if err != nil {
		return nil, err
	}
	return all, nil
}