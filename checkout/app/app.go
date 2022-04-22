package app

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"qwik.in/checkout/docs"
	_ "qwik.in/checkout/docs"
	"qwik.in/checkout/domain/repository"
	"qwik.in/checkout/domain/services"
	"qwik.in/checkout/domain/tools/logger"
)

var shippingHandler ShippingHandler
var checkoutHandler CheckoutHandler
var cartHandler CartHandler

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
	CheckoutRouter(router)
	CartRouter(router)
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Shipping API"
}

func Start() {

	shippingRepo := repository.ShippingAddrRepoFunc()
	shippingHandler = ShippingHandler{
		ShippingService: services.ShippingAddressServiceFunc(shippingRepo),
	}
	cartRepo := repository.CartRepoFunc()
	cartHandler = CartHandler{
		CartService: services.CartServiceFunc(cartRepo),
	}

	//Custom Logger - Logs actions to 'shippingAddressService.logger' file
	file, err := os.OpenFile("./logs/shippingAddressServiceLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
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
