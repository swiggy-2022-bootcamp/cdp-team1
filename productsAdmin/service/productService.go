package service

import (
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/proto"
)

type ProductService interface {
	CreateProduct(product entity.Product) error
	GetAll() ([]entity.Product, error)
	UpdateProduct(productId string, product entity.Product) error
	DeleteProduct(productId string) error
	SearchProduct(limit int64) ([]entity.Product, error)
	GetQuantityForProductId(productId string) (*proto.Response, error)
}
