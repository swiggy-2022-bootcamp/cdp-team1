package service

import (
	"context"
	"qwik.in/productsAdmin/log"
	"qwik.in/productsAdmin/proto/productPrice"
	"qwik.in/productsAdmin/repository"
	"strconv"
)

type ProductPriceService struct {
	productPrice.UnimplementedProductPriceServiceServer
}

var (
	Repository repository.ProductRepository
)

func NewProductPriceService() *ProductPriceService {
	Repository = repository.NewDynamoRepository()
	err := Repository.Connect()
	if err != nil {
		log.Error("Error while Connecting to DB: ", err)
		return nil
	}
	return &ProductPriceService{}
}

func (p *ProductPriceService) GetTotalPriceForProducts(ctx context.Context, products *productPrice.ProductsPriceRequests) (*productPrice.ResponsePrice, error) {
	log.Info("gRPC received message: ", products)

	totalPrice := 0.0

	for _, product := range products.GetProducts() {
		productObject, err := Repository.FindOne(product.Id)
		if err != nil {
			log.Error(err.Error())
			continue
		}
		quantity, err := strconv.Atoi(product.Quantity)
		price, err := strconv.ParseFloat(productObject.Price, 32)
		if err != nil {
			log.Error(err.Error())
			continue
		}

		totalPrice += price * float64(quantity)
	}
	return &productPrice.ResponsePrice{Price: totalPrice}, nil
}
