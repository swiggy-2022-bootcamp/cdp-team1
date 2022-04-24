package repository

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"qwik.in/account-frontstore/protos"
	"time"
)

type GrpcClientRepositoryInterface interface {
	GetRewardPointsByCustomerId(customerId string) int32
	GetPaymentMethodsByCustomerId(customerId string) []*protos.PaymentMode
}

type GrpcClientRepository struct {
}

func (grpcClientRepository *GrpcClientRepository) GetRewardPointsByCustomerId(customerId string) int32 {
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
