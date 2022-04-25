package service

import (
	"context"
	"fmt"
	"orderService/domain/model"
	"orderService/domain/repository"
	"orderService/log"
	"orderService/protos"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var orderRepository repository.OrderRepositoryDB

type OrderProtoServer struct {
	protos.UnimplementedOrderServer
}

func NewOrderProtoService(pr repository.OrderRepositoryDB) OrderProtoServer {
	orderRepository = pr
	return OrderProtoServer{}
}

func (o OrderProtoServer) GetAmountFromProduct(products []*protos.ProductPriceRequest) (*protos.ResponsePrice, error) {

	serverAddress := ":19191"
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.ErrorLogger.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := protos.NewProductPriceServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	r, err := c.GetTotalPriceForProducts(ctx, &protos.ProductsPriceRequests{
		Products: products,
	})
	if err != nil {
		log.ErrorLogger.Fatalf("could not echo: %v", err)
	}
	fmt.Println(r)
	time.Sleep(time.Second)
	return r, nil
}

func (o OrderProtoServer) GetCartFromCartService(customer_id string) (*protos.GetCartResponse, error) {
	log.Info("Connecting with gRPC server")
	// Set up a connection to the server.
	serverAddress := "localhost:5004"
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: ", err)
		return nil, err
	}

	//close
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Error("Connection closed with error", err.Error())
		}
	}(conn)
	c := protos.NewCartClient(conn)

	// Disconnect gRPC call upon
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Send gRPC request to server the
	r, err := c.GetCart(ctx, &protos.GetCartRequest{CustomerId: customer_id})
	if err != nil {
		log.Error("could not get response: ", err)
		return nil, err
	}

	// log.Info("gRPC received id: ", r.GetId(), " and quantity: ", r.GetQuantity())
	return r, nil
}

func (o OrderProtoServer) CreateOrder(ctx context.Context, req *protos.CreateOrderRequest) (*protos.CreateOrderResponse, error) {

	fmt.Println("Create Order check")

	// Get cart from gRPC server
	customer_id := req.GetCustomerId()

	cart, err := o.GetCartFromCartService(customer_id)
	if err != nil {
		log.Error("could not get cart: ", err)
		return nil, err
	}

	fmt.Println("Cart: ", cart)
	// return nil, nil

	products := make([]model.Product, 0)

	for _, product := range cart.GetProducts() {
		products = append(products, model.Product{
			ProductID: product.GetProductId(),
			Quantity:  int(product.GetQuantity()),
		})
	}

	fmt.Println("Products: ", products)

	products_amount := make([]*protos.ProductPriceRequest, 0)
	for _, product := range products {
		products_amount = append(products_amount, &protos.ProductPriceRequest{
			Id:       product.ProductID,
			Quantity: string(rune(product.Quantity)),
		})
	}

	fmt.Println("Products amount: ", products_amount)

	// Get amount from gRPC server
	amount, err := o.GetAmountFromProduct(products_amount)
	if err != nil {
		log.Error("could not get amount: ", err)
		return nil, err
	}

	fmt.Println("Amount received: ", amount)

	order := model.Order{
		CustomerId: customer_id,
		Amount:     int(amount.Price),
		Status:     "pending",
		Products:   products,
	}

	fmt.Println("Order: ", order)

	// Create order
	err2 := orderRepository.Create(&order)

	if err2 != nil {
		log.Error("could not create order: ", err2)
		return nil, err2.Error()
	}

	fmt.Println("Order created")

	// Create response
	// new_products := make([]protos.Product, 0)
	// response := &protos.CreateOrderResponse{
	// 	OrderId:    order.OrderId,
	// 	CustomerId: order.CustomerId,
	// 	Status:     order.Status,
	// 	Datetime:   order.Datetime,
	// 	Amount:     int32(order.Amount),
	// 	// Products:  	,
	// }

	return nil, nil

	// var productProto *protos.Product

	// for _, product := range cart.Products {
	// 	productProto = &protos.Product{
	// 		ProductId: product.ProductId,
	// 		Quantity:  int32(product.Quantity),
	// 	}
	// }

	// products = append(products, productProto)

	// response.Products = new_products
	// return response, nil
}
