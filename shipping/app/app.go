package app

import (
	"io"
	"os"
	"qwik.in/shipping/docs"
	"qwik.in/shipping/domain/repository"
	"qwik.in/shipping/domain/services"
	"qwik.in/shipping/domain/tools/logger"

	"github.com/gin-gonic/gin"
)

var shippingHandler ShippingHandler

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

	//Custom Logger - Logs actions to 'shippingAddressService.logger' file
	file, err := os.OpenFile("shippingServer.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configure - ShippingAddress Server and Router

	PORT := os.Getenv("SHIPPING_SERVICE_PORT")
	router := setupRouter()

	configureSwaggerDoc()
	err = router.Run(":" + PORT)
	if err != nil {
		return
	}
}
