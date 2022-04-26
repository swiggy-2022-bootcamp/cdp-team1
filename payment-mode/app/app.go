package app

import (
	"context"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/gin-gonic/gin"
	prometheusUtility "github.com/swiggy-2022-bootcamp/cdp-team1/common-utilities/prometheus-utility"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"os/signal"
	"qwik.in/payment-mode/app/handlers"
	"qwik.in/payment-mode/app/routes"
	"qwik.in/payment-mode/config"
	"qwik.in/payment-mode/domain/repository"
	"qwik.in/payment-mode/domain/services"
	"qwik.in/payment-mode/log"
	"qwik.in/payment-mode/protos"
	"sync"
	"syscall"
	"time"
)

var (
	server              *gin.Engine
	paymentRepository   repository.PaymentRepository
	paymentService      services.PaymentService
	paymentHandler      handlers.PaymentHandler
	paymentRoutes       routes.PaymentRoutes
	ctx                 context.Context
	paymentDB           *dynamodb.DynamoDB
	healthCheckHandler  handlers.HealthCheckHandler
	paymentServiceProto services.PaymentProtoServer
	wg                  sync.WaitGroup
)

func Start() {
	ctx = context.TODO()

	//Variable initializations for DynamoDB
	paymentDB = config.ConnectDB()
	config.CreateTable(paymentDB)

	//Variable initializations to be used as dependency injectors
	paymentRepository = repository.NewPaymentRepositoryImpl(paymentDB, ctx)
	paymentService = services.NewPaymentServiceImpl(paymentRepository)
	paymentHandler = handlers.NewPaymentHandler(paymentService)
	healthCheckHandler = handlers.NewHealthCheckHandler(paymentRepository)
	paymentRoutes = routes.NewPaymentRoutes(paymentHandler, healthCheckHandler)
	paymentServiceProto = services.NewPaymentProtoService(paymentRepository)

	//Opening file for log collection
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
	wg.Wait()
}

func StartRESTServer() {
	defer wg.Done()

	//Configuring gin server and router
	server = gin.New()
	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	server.Use(prometheusUtility.PrometheusMiddleware())
	router := server.Group("/api")
	paymentRoutes.InitRoutes(router)

	//Starting server on port 9000
	err := server.Run(":9000")
	if err != nil {
		log.Error(err.Error() + " - Failed to start server")
	} else {
		log.Info("Server started successfully.")
	}
}

func StartGRPCServer() {
	defer wg.Done()
	//Opening PORT 9004 for GRPC server
	lis, err := net.Listen("tcp", ":9004")
	if err != nil {
		log.Error("Failed to listen on port %s with error %v", "9004", err)
	}

	//Creating and registering the GRPC server
	grpcServer := grpc.NewServer()
	protos.RegisterPaymentServer(grpcServer, paymentServiceProto)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("Failed to start the grpc server : %v", err)
	}
	log.Info("GRPC server started successfully")
}
