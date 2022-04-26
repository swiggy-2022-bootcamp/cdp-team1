package app

import (
	"cartService/app/handlers"
	"cartService/app/routes"
	"cartService/db"
	"cartService/domain/repository"
	"cartService/domain/service"
	"cartService/log"
	"cartService/protos"
	"context"
	"io"
	"net"
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
	cartRepository     repository.CartRepositoryDB
	cartService        service.CartService
	cartHandler        handlers.CartHandler
	cartRoutes         routes.CartRoutes
	ctx                context.Context
	cartDB             *dynamodb.DynamoDB
	healthCheckHandler handlers.HealthCheckHandler
	cartServiceProto   service.CartProtoServer
	wg                 sync.WaitGroup
)

func Start() {

	ctx = context.TODO()

	//Variable initializations for DynamoDB
	cartDB = db.ConnectDB()
	db.CreateTable(cartDB)

	//Variable initializations to be used as dependency injectors
	cartRepository = repository.NewCartRepository(cartDB, ctx)
	cartService = service.NewCartService(cartRepository)
	cartHandler = handlers.NewCartHandler(cartService)
	healthCheckHandler = handlers.NewHealthCheckHandler(cartRepository)
	cartRoutes = routes.NewCartRoutes(cartHandler, healthCheckHandler)
	cartServiceProto = service.NewCartProtoService(cartRepository)

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
	go StartRestServer()
	go StartGRPCServer()
	wg.Wait()
}

func StartRestServer() {

	//Configuring gin server and router
	server = gin.New()
	server.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	router := server.Group("api/")
	cartRoutes.InitRoutes(router)

	//Starting server on port 5000
	err := server.Run(":5000")
	if err != nil {
		log.Error(err.Error() + " - Failed to start server")
	} else {
		log.Info("Server started successfully.")
	}
}

func StartGRPCServer() {

	defer wg.Done()
	//Opening PORT 5004 for GRPC server
	lis, err := net.Listen("tcp", ":5004")
	if err != nil {
		log.Error("Failed to listen on port %s with error %v", "9004", err)
	}

	//Creating and registering the GRPC server
	grpcServer := grpc.NewServer()
	protos.RegisterCartServer(grpcServer, cartServiceProto)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("Failed to start the grpc server : %v", err)
	}
	log.Info("GRPC server started successfully")
}
