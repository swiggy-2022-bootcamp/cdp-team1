package service

import (
	"context"
	"qwik.in/productsAdmin/config"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
	"qwik.in/productsAdmin/proto"
	"qwik.in/productsAdmin/repository"
)

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &ProductServiceImpl{productRepository: productRepository}
}

func (p ProductServiceImpl) CreateProduct(product entity.Product) error {
	product.SetId()
	if err := p.productRepository.SaveProduct(product); err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) GetAll() ([]entity.Product, error) {
	all, err := p.productRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (p ProductServiceImpl) UpdateProduct(productId string, product entity.Product) error {
	product.ID = productId
	if err := p.productRepository.SaveProduct(product); err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) DeleteProduct(productId string) error {
	err := p.productRepository.DeleteProduct(productId)
	if err != nil {
		return err
	}
	return nil
}

func (p ProductServiceImpl) SearchProduct(limit int64) ([]entity.Product, error) {
	all, err := p.productRepository.FindWithLimit(limit)
	if err != nil {
		return nil, err
	}
	return all, nil
}

func (p ProductServiceImpl) GetQuantityForProductId(productId string) (*proto.Response, error) {
	log.Info("Connecting with gRPC server")
	// Set up a connection to the server.
	conn, err := grpc.Dial(":"+config.GRPC_SERVER_PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: ", err)
		return nil, err
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	c := proto.NewQuantityServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	r, err := c.GetQuantity(ctx, &proto.Request{Id: productId})
	if err != nil {
		log.Error("could not get response: ", err)
		return nil, err
	}
	log.Info("gRPC received id: ", r.GetId(), " and quantity: ", r.GetQuantity())
	return r, nil
}
