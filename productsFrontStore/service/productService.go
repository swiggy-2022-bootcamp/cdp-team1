package service

import "qwik.in/productsFrontStore/entity"

type ProductService interface {
	GetAll() ([]entity.Product, error)
	GetProductById(productId string) (entity.Product, error)
	GetProductsByCategory(categoryId string) ([]entity.Product, error)
}
