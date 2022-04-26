package repository

import "qwik.in/productsAdmin/entity"

type ProductRepository interface {
	Connect() error
	FindOne(productId string) (entity.Product, error)
	FindAll() ([]entity.Product, error)
	SaveProduct(product entity.Product) error
	DeleteProduct(productId string) error
	FindWithLimit(limit int64) ([]entity.Product, error)
}
