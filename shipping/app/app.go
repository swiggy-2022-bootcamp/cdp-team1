package app

import (
	"google.golang.org/grpc"
	"io"
	"net"
	"os"
	"os/signal"
	"qwik.in/shipping/docs"
	"qwik.in/shipping/domain/repository"
	"qwik.in/shipping/domain/services"
	"qwik.in/shipping/domain/tools/logger"
	"qwik.in/shipping/protos"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

var shippingHandler ShippingHandler
var wg sync.WaitGroup
var shippingServiceProto services.ShippingProtoServer

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}
	router.Use(logger.UseLogger(logger.DefaultLoggerFormatter), gin.Recovery())
	// health check route
	HealthCheckRouter(router)
	ShippingRouter(router)
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Shipping API"
}

//Start ..
func Start() {
	shippingRepo := repository.ShippingAddrRepoFunc()
	shippingHandler = ShippingHandler{
		ShippingAddrService: services.ShippingAddressServiceFunc(shippingRepo),
	}
	shippingServiceProto = services.NewShippingRepoService(shippingRepo)

	//Custom Logger - Logs actions to 'shippingAddressService.logger' file
	file, err := os.OpenFile("shippingServer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	// gracefulStop logic to allow go routines to finish
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		logger.Info("caught sig: %+v", sig)
		logger.Info("Wait for 2 second to finish processing")
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
	//Configure - ShippingAddress Server and Router

	//PORT := os.Getenv("SHIPPING_SERVICE_PORT")
	router := setupRouter()

	configureSwaggerDoc()
	err := router.Run(":" + "9005")
	if err != nil {
		return
	}
}

func StartGRPCServer() {
	defer wg.Done()
	//Opening PORT 9005 for GRPC server
	lis, err := net.Listen("tcp", ":9006")
	if err != nil {
		logger.Error("Failed to listen on port %s with error %v", "9004", err)
	}

	//Creating and registering the GRPC server
	grpcServer := grpc.NewServer()
	//protos.RegisterPaymentServer(grpcServer, paymentServiceProto)
	protos.RegisterShippingAddressProtoFuncServer(grpcServer, shippingServiceProto)
	err = grpcServer.Serve(lis)
	if err != nil {
		logger.Error("Failed to start the grpc server : %v", err)
	}
	logger.Info("GRPC server started successfully")
}
