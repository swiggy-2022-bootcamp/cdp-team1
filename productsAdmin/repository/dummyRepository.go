package repository

import (
	"encoding/json"
	"os"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
)

type repo struct{}

func (r repo) Connect() error {
	log.Info("Dummy Repo Connected")
	return nil
}

func (r repo) SaveProduct(product entity.Product) error {
	log.Info("Product Saved")
	return nil
}

func (r repo) FindAll() ([]entity.Product, error) {
	productData, _ := os.ReadFile("/Users/aky/Dev/swiggy/cdp-team1/productsAdmin/entity/sampleProducts.json")
	var products []entity.Product
	if err := json.Unmarshal(productData, &products); err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Found these products: ", products)
	return products, nil
}

func (r repo) DeleteProduct(productId string) error {
	log.Info("Product Deleted: ", productId)
	return nil
}

func (r repo) FindAndUpdate(product entity.Product) error {
	log.Info("Product Updated: ", product)
	return nil
}

func NewDummyRepository() ProductRepository {
	return &repo{}
}
