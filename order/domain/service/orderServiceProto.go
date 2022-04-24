package service

import (
	"orderService/domain/repository"
	"orderService/protos"
)

var orderRepository repository.OrderRepository

type OrderProtoServer struct {
	protos.UnimplementedOrderServer
}

func NewOrderProtoService(pr repository.OrderRepository) OrderProtoServer {
	orderRepository = pr
	return OrderProtoServer{}
}

func () GetAmountFromProduct(customer_id string) (*proto.Cart, error) {
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
	r, err := c.GetCart(ctx, &proto.GetCartRequest{CustomerId: customer_id})
	if err != nil {
		log.Error("could not get response: ", err)
		return nil, err
	}

	// log.Info("gRPC received id: ", r.GetId(), " and quantity: ", r.GetQuantity())
	return r, nil
}

func () GetCartFromCartService(customer_id string) (*proto.Cart, error) {
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
	r, err := c.GetCart(ctx, &proto.GetCartRequest{CustomerId: customer_id})
	if err != nil {
		log.Error("could not get response: ", err)
		return nil, err
	}

	// log.Info("gRPC received id: ", r.GetId(), " and quantity: ", r.GetQuantity())
	return r, nil
}

// func (o OrderProtoServer) CreateOrder(ctx context.Context, req *protos.CreateOrderRequest) (*protos.CreateOrderResponse, error) {

	// Get cart from gRPC server
	customer_id := req.GetCustomerId()
	cart, err := GetCartFromCartService(customer_id)
	if err != nil {
		log.Error("could not get cart: ", err)
		return nil, err
	}

	fmt.Println("Cart: ", cart)
	// Create order
	// order := repository.CreateOrder(cart)

	// // Create response
	// products := make(*[]protos.Product, 0)
	// response := &protos.CreateOrderResponse{
	// 	order_id:    order.ID,
	// 	customer_id: order.CustomerID,
	// 	status:      order.Status,
	// 	datetime:    order.Datetime,
	// 	amount:      order.Amount,
	// 	products:    products,
	// }

	// var productProto *protos.Product
	// for _, product := range cart.Products {
	// 	productProto = &protos.Product{
	// 		ProductId: product.ProductId,
	// 		Quantity:  int32(product.Quantity),
	// 	}
	// 	products = append(products, productProto)
	// }

	// response.Products = products
	// return response, nil
}
