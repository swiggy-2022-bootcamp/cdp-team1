package service

import (
	"cartService/domain/repository"
	"cartService/protos"
	"context"
	"fmt"
)

var cartRepository repository.CartRepositoryDB

type CartProtoServer struct {
	protos.UnimplementedCartServer
}

func NewCartProtoService(pr repository.CartRepositoryDB) CartProtoServer {
	cartRepository = pr
	return CartProtoServer{}
}

func (c CartProtoServer) GetCart(ctx context.Context, req *protos.GetCartRequest) (*protos.GetCartResponse, error) {

	fmt.Println("GetCart check")

	products := make([]*protos.Product, 0)
	response := &protos.GetCartResponse{
		CustomerId: req.GetCustomerId(),
		Products:   products,
	}

	customer_id := req.GetCustomerId()

	cart, err := cartRepository.Read(customer_id)
	if err != nil {
		return nil, err.Error()
	}

	var productProto *protos.Product
	for _, product := range cart.Products {
		productProto = &protos.Product{
			ProductId: product.ProductId,
			Quantity:  int32(product.Quantity),
		}
		products = append(products, productProto)
	}

	response.Products = products
	return response, nil
}
