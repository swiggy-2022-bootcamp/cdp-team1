package service

import "qwik.in/productsAdmin/entity"

type ProductService interface {
	CreateProduct(product entity.Product) error
	GetAll() ([]entity.Product, error)
	UpdateProduct(productId string, product entity.Product) error
	DeleteProduct(productId string) error
	SearchProduct(limit int64) ([]entity.Product, error)
}
