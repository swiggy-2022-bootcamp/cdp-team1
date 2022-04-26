package repository

import "qwik.in/productsFrontStore/entity"

type ProductRepository interface {
	Connect() error
	FindAll() ([]entity.Product, error)
	FindById(productId string) (entity.Product, error)
	FindByCategory(categoryId string) ([]entity.Product, error)
}
