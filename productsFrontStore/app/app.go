package app

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"os/signal"
	"qwik.in/productsFrontStore/app/handlers"
	"qwik.in/productsFrontStore/app/routes"
	"qwik.in/productsFrontStore/config"
	"qwik.in/productsFrontStore/log"
	"qwik.in/productsFrontStore/proto"
	"qwik.in/productsFrontStore/repository"
	"qwik.in/productsFrontStore/service"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func Start() {

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
	go initRestServer()
	go initGrpcServer()
	wg.Wait()
}

func initGrpcServer() {
	defer wg.Done()
	lis, err := net.Listen("tcp", ":"+config.GRPC_SERVER_PORT)
	if err != nil {
		log.Error("failed to GRPC listen: ", err)
	}
	gs := grpc.NewServer()
	proto.RegisterQuantityServiceServer(gs, service.NewProductQualityService())
	log.Info("gRPC Server: Listening on port ", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		log.Error("gRPC Server: Failed to serve : ", err.Error())
	}
}

var (
	Repository repository.ProductRepository
	Service    service.ProductService
	Handler    handlers.ProductHandler
)

func initRestServer() {
	defer wg.Done()
	Repository = repository.NewDynamoRepository()
	err := Repository.Connect()
	if err != nil {
		log.Error("Error while Connecting to DB: ", err)
		return
	}
	Service = service.NewProductService(Repository)
	Handler = handlers.NewProductHandler(Service)

	file, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil && config.LOG_FILE_MODE {
		log.Info("Opened log file successfully")
		gin.DefaultWriter = io.MultiWriter(file)
	} else {
		log.Warn("Could not open log file, switching to IO mode")
	}

	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())

	routes.InitRoutes(router, Handler)

	log.Info(config.SERVICE_NAME, " starting on port ", config.REST_SERVER_PORT)
	err = router.Run(":" + config.REST_SERVER_PORT)

	if err != nil {
		log.Error(config.SERVICE_NAME, " startup failed")
		return
	}
}
