package app

import (
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"qwik.in/rewards/app/handlers"
	"qwik.in/rewards/app/routes"
	"qwik.in/rewards/config"
	"qwik.in/rewards/log"
	"qwik.in/rewards/proto"
	"qwik.in/rewards/repository"
	"qwik.in/rewards/service"
)

var (
	Repository repository.RewardRepository
	Service    service.RewardService
	Handler    handlers.RewardHandler
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
	// proto.RegisterRewardServiceServer(gs, &RewardServiceImpl{})
	proto.RegisterRewardPointServiceServer(gs, service.NewRewardPointService())
	log.Info("gRPC Server: Listening on port ", lis.Addr())
	if err := gs.Serve(lis); err != nil {
		log.Error("gRPC Server: Failed to serve : ", err.Error())
	}
}

func initRestServer() {

	RewardRepository := repository.NewDynamoRepository()
	err := RewardRepository.Connect()
	if err != nil {
		return
	}
	fmt.Println("Connection successful")
	rewardService := service.NewRewardService(RewardRepository)
	fmt.Println("Service created")
	rewardController := handlers.NewRewardHandler(rewardService)
	fmt.Println("Controller created")
	file, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil && config.LOG_FILE_MODE {
		log.Info("Opened log file successfully")
		gin.DefaultWriter = io.MultiWriter(file)
	} else {
		log.Warn("Could not open log file, switching to IO mode")
	}

	router := gin.New()
	router.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())

	routes.InitRoutes(router, rewardController)

	log.Info(config.SERVICE_NAME, " starting on port ", config.SERVER_PORT)
	err = router.Run(":" + config.SERVER_PORT)

	if err != nil {
		log.Error(config.SERVICE_NAME, " startup failed")
		return
	}
}
