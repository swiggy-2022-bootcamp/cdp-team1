package app

import (
	"authService/app/routes"
	"authService/config"
	"authService/db"
	"authService/domain"
	"authService/log"
	"authService/protos"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func Start() {

	if fErr := config.LoadEnvConfig(); fErr != nil {
		log.Error(fErr.Message)
	}

	dbClient := db.NewDynamoDBClient()
	authRepo := db.NewAuthRepositoryDB(dbClient, 100)
	adminRepo := db.NewAdminRepositoryDB(dbClient, 100)
	custRepo := db.NewCustomerRepositoryDB(dbClient, 100)
	authSvc := domain.NewAuthService(authRepo, adminRepo, custRepo)
	authProtoSvc := domain.NewAuthProtoService(authSvc)

	//Opening file for log collection
	file, err := os.OpenFile("./server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	} else {
		log.Error(err)
	}

	// gracefulStop logic to allow go routines to finish
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		log.Info("caught sig: %+v", sig)
		log.Info("Waiting for 5 seconds to finish processing")
		time.Sleep(5 * time.Second)
		os.Exit(0)
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	go StartRESTServer(&wg, authSvc)
	go StartGRPCServer(&wg, authProtoSvc)
	wg.Wait()
}

func StartRESTServer(wg *sync.WaitGroup, authSvc domain.AuthService) {
	defer wg.Done()
	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	routes.RegisterAuthRoutes(router, authSvc)

	//Starting Gin Server
	port := config.EnvVars.GinPort
	err := router.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error(err.Error() + " - Failed to start Gin Server")
	} else {
		log.Info("Server started successfully.")
	}
}

func StartGRPCServer(wg *sync.WaitGroup, authProtoSvc *domain.AuthProtoService) {
	defer wg.Done()

	port := config.EnvVars.GrpcPort
	//Starting GRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error("Failed to listen on port %s with error %v", port, err)
	}

	grpcServer := grpc.NewServer()
	protos.RegisterAuthServer(grpcServer, authProtoSvc)
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Error("Failed to start the grpc server : %v", err)
	}
	log.Info("gRPC Server started successfully")
}
