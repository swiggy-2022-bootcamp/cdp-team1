package app

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"os/signal"
	"qwik.in/transaction/app/handlers"
	"qwik.in/transaction/app/routes"
	"qwik.in/transaction/config"
	"qwik.in/transaction/domain/repository"
	"qwik.in/transaction/domain/services"
	"qwik.in/transaction/log"
	"qwik.in/transaction/protos"
	"sync"
	"syscall"
	"time"
)

var (
	restServer              *gin.Engine
	transactionRepository   repository.TransactionRepository
	transactionService      services.TransactionService
	transactionHandler      handlers.TransactionHandler
	transactionRoutes       routes.TransactionRoutes
	transactionDB           *dynamodb.DynamoDB
	healthCheckHandler      handlers.HealthCheckHandler
	transactionServiceProto services.TransactionProtoServer
	wg                      sync.WaitGroup
)

func Start() {

	//Variable initializations for DynamoDB
	transactionDB = config.ConnectDB()
	config.CreateTable(transactionDB)

	//Variable initializations to be used as dependency injections
	transactionRepository = repository.NewTransactionRepositoryImpl(transactionDB)
	transactionService = services.NewTransactionServiceImpl(transactionRepository)
	transactionHandler = handlers.NewTransactionHandler(transactionService)
	healthCheckHandler = handlers.NewHealthCheckHandler(transactionRepository)
	transactionRoutes = routes.NewTransactionRoutes(transactionHandler, healthCheckHandler)
	transactionServiceProto = services.NewTransactionProtoServer(transactionRepository)

	//Opening file for log collection
	file, err := os.OpenFile("./server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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

	wg.Add(2)
	go StartRESTServer()
	go StartGRPCServer()
	wg.Wait()
}

func StartRESTServer() {
	defer wg.Done()
	restServer = gin.New()
	restServer.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := restServer.Group("api")
	transactionRoutes.InitRoutes(router)

	//Starting restServer on port 9001
	err := restServer.Run(":9001")
	if err != nil {
		log.Error(err.Error() + " - Failed to start restServer")
	} else {
		log.Info("Server started successfully.")
	}
}

func StartGRPCServer() {
	defer wg.Done()
	//Starting GRPC server on PORT 9003
	lis, err := net.Listen("tcp", ":9003")
	if err != nil {
		log.Error("Failed to listen on port %s with error %v", "9003", err)
	}

	grpcServer := grpc.NewServer()
	protos.RegisterTransactionPointsServer(grpcServer, transactionServiceProto)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("Failed to start the grpc server : %v", err)
	}
	log.Info("GRPC Server started successfully")
}
