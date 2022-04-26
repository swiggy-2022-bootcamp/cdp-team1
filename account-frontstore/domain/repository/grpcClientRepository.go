package repository

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"qwik.in/account-frontstore/config"
	log2 "qwik.in/account-frontstore/log"
	"qwik.in/account-frontstore/proto"
	"qwik.in/account-frontstore/protos"
	"time"
)

type GrpcClientRepositoryInterface interface {
	GetTransactionRewardPointsByCustomerId(customerId string) int32
	GetPaymentMethodsByCustomerId(customerId string) []*protos.PaymentMode
	GetCartByCustomerId(customerId string) ([]*protos.Product, error)
	GetRewardPointsByCustomerId(customerId string) (*int32, error)
}

type GrpcClientRepository struct {
}

func (grpcClientRepository *GrpcClientRepository) GetTransactionRewardPointsByCustomerId(customerId string) int32 {
	conn, err := grpc.Dial("localhost:9003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection failed : %v", err)
	}
	defer conn.Close()
	c := protos.NewTransactionPointsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	getPointsRequest := &protos.GetPointsRequest{
		UserId: "bb912edc-50d9-42d7-b7a1-9ce66d459tuf",
		//UserId: customerId,
	}

	result, err := c.GetTransactionPoints(ctx, getPointsRequest)
	if err != nil {
		log.Printf("Error getting transaction points : %v", err)
	} else {
		log.Printf("Transaction points are %d", result.Points)
	}
	return result.Points
}

func (grpcClientRepository *GrpcClientRepository) GetPaymentMethodsByCustomerId(customerId string) []*protos.PaymentMode {
	conn, err := grpc.Dial("localhost:9004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Connection failed : %v", err)
	}
	defer conn.Close()
	c := protos.NewPaymentClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	paymentModeRequest := &protos.PaymentModeRequest{
		UserId: "bb912edc-50d9-42d7-b7a1-9ce66d457ecc",
	}

	result, err := c.GetPaymentModes(ctx, paymentModeRequest)
	if err != nil {
		log.Printf("Failed to fetch payment modes for the given user : %v", err)
	} else {
		log.Printf("%s - Payment Successful", result.PaymentModes)
	}
	return result.PaymentModes
}

func (grpcClientRepository *GrpcClientRepository) GetCartByCustomerId(customerId string) ([]*protos.Product, error) {
	serverAddress := "localhost:5004"
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("did not connect: : %v", err)
		log2.Error("did not connect: ", err)
		return nil, err
	}

	//close
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Connection closed with error: : %v", err)
			log2.Error("Connection closed with error", err.Error())
		}
	}(conn)
	c := protos.NewCartClient(conn)

	// Disconnect gRPC call upon
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Send gRPC request to server the
	r, err := c.GetCart(ctx, &protos.GetCartRequest{CustomerId: customerId})
	if err != nil {
		log.Printf("could not get response:  #{err}")
		log2.Error("could not get response: ", err)
		return nil, err
	}

	// log.Info("gRPC received id: ", r.GetId(), " and quantity: ", r.GetQuantity())
	log.Println(r.Products)
	return r.Products, nil
}

func (grpcClientRepository *GrpcClientRepository) GetRewardPointsByCustomerId(customerId string) (*int32, error) {
	serverAddress := fmt.Sprintf("%s:%s", config.GRPC_SERVER_IP, config.GRPC_SERVER_PORT)
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log2.Error("did not connect: ", err)
		return nil, err
	}

	//close
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log2.Error("Connection closed with error", err.Error())
		}
	}(conn)
	c := proto.NewRewardPointServiceClient(conn)

	// Disconnect gRPC call upon
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Send gRPC request to server the
	r2, err := c.GetRewardPoints(ctx, &proto.Request{Id: customerId})
	if err != nil {
		log2.Error("could not get response: ", err)
		return nil, err
	}
	log2.Info("Response from gRPC server", r2)
	return &r2.RewardPoints, nil
}
