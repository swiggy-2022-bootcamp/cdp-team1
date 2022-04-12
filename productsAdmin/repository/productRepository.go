package repository

import "qwik.in/productsAdmin/entity"

type ProductRepository interface {
	Connect() error
	FindAll() ([]entity.Product, error)
	SaveProduct(product entity.Product) error
	DeleteProduct(productId string) error
	FindAndUpdate(product entity.Product) error
}
