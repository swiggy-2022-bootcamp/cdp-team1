package service

import "qwik.in/productsAdmin/entity"

type ProductService interface {
	CreateProduct(product entity.Product) error
	GetAll() ([]entity.Product, error)
	UpdateProduct(product entity.Product) error
	DeleteProduct(productId string) error
	SearchProduct(query string) ([]entity.Product, error)
}
