package repository

import (
	"encoding/json"
	"os"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
)

type dummyRepository struct{}

func (r dummyRepository) Connect() error {
	log.Info("Dummy Repo Connected")
	return nil
}

func (r dummyRepository) SaveProduct(product entity.Product) error {
	log.Info("Product Saved")
	return nil
}

func (r dummyRepository) FindOne(productId string) (entity.Product, error) {
	productData, _ := os.ReadFile("/Users/aky/Dev/swiggy/cdp-team1/productsAdmin/entity/sampleProductSingle.json")
	var product entity.Product
	if err := json.Unmarshal(productData, &product); err != nil {
		log.Error(err)
		return entity.Product{}, err
	}
	log.Info("Found these products: ", product)
	return product, nil
}

func (r dummyRepository) FindAll() ([]entity.Product, error) {
	productData, _ := os.ReadFile("/Users/aky/Dev/swiggy/cdp-team1/productsAdmin/entity/sampleProducts.json")
	var products []entity.Product
	if err := json.Unmarshal(productData, &products); err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("Found these products: ", products)
	return products, nil
}

func (r dummyRepository) DeleteProduct(productId string) error {
	log.Info("Product Deleted: ", productId)
	return nil
}

func (r dummyRepository) FindAndUpdate(product entity.Product) error {
	log.Info("Product Updated: ", product)
	return nil
}

func NewDummyRepository() ProductRepository {
	return &dummyRepository{}
}
