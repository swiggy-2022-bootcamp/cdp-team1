package app

import (
	"context"
	"io"
	"net"
	"orderService/app/handlers"
	"orderService/app/routes"
	"orderService/db"
	"orderService/domain/repository"
	"orderService/domain/service"
	"orderService/log"
	"orderService/protos"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	server             *gin.Engine
	orderRepository    repository.OrderRepositoryDB
	orderService       service.OrderService
	orderHandler       handlers.OrderHandler
	orderRoutes        routes.OrderRoutes
	ctx                context.Context
	orderDB            *dynamodb.DynamoDB
	healthCheckHandler handlers.HealthCheckHandler
	wg                 sync.WaitGroup
	orderServiceProto  service.OrderProtoServer
)

func Start() {

	ctx = context.TODO()

	//Variable initializations for DynamoDB
	orderDB = db.ConnectDB()
	db.CreateTable(orderDB)

	//Variable initializations to be used as dependency injectors
	orderRepository = repository.NewOrderRepository(orderDB, ctx)
	orderService = service.NewOrderService(orderRepository)
	orderHandler = handlers.NewOrderHandler(orderService)
	healthCheckHandler = handlers.NewHealthCheckHandler(orderRepository)
	orderRoutes = routes.NewOrderRoutes(orderHandler, healthCheckHandler)
	orderServiceProto = service.NewOrderProtoService(orderRepository)

	obj := service.NewOrderProtoService(orderRepository)
	obj.CreateOrder(ctx, &protos.CreateOrderRequest{CustomerId: "2"})

	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}
	// gracefulStop logic to allow go routines to finish
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Info("caught sig: %+v", sig)
		log.Info("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	}()
	//Starting REST server and GRPC server
	wg.Add(2)
	go StartRESTServer()
	go StartGRPCServer()
	//obj := service.NewOrderProtoService(orderRepository)
	//obj.CreateOrder(ctx, &protos.CreateOrderRequest{CustomerId: "2"})
	wg.Wait()
}

//StartRESTServer ..
func StartRESTServer() {
	//Configuring gin server and router
	server = gin.New()
	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("order/api")
	orderRoutes.InitRoutes(router)
	//Starting server on port 7000
	err := server.Run(":7000")
	if err != nil {
		log.Error(err.Error() + " - Failed to start server")
	} else {
		log.Info("Server started successfully.")
	}
}

//StartGRPCServer ..
func StartGRPCServer() {
	defer wg.Done()
	//Opening PORT 5001 for GRPC server
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Error("Failed to listen on port %s with error %v", "9004", err)
	}
	//Creating and registering the GRPC server
	grpcServer := grpc.NewServer()
	protos.RegisterOrderServer(grpcServer, orderServiceProto)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("Failed to start the grpc server : %v", err)
	}
	log.Info("GRPC server started successfully")
}
