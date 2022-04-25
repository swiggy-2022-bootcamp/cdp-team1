package app

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
	"qwik.in/checkout/docs"
	_ "qwik.in/checkout/docs" //GoSwagger
	"qwik.in/checkout/domain/repository"
	"qwik.in/checkout/domain/services"
	"qwik.in/checkout/domain/tools/logger"
)

var shippingHandler ShippingHandler
var checkoutHandler CheckoutHandler
var cartHandler CartHandler
var paymentHandler PaymentHandler

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
	//ShippingRouter(router)
	CheckoutRouter(router)
	//CartRouter(router)
	//PaymentRouter(router)
	return router
}

func configureSwaggerDoc() {
	docs.SwaggerInfo.Title = "Swagger Checkout API"
}

//Start ..
func Start() {

	shippingRepo := repository.ShippingAddrRepoFunc()
	shippingHandler = ShippingHandler{
		ShippingService: services.ShippingAddressServiceFunc(shippingRepo),
	}
	cartRepo := repository.CartRepoFunc()
	cartHandler = CartHandler{
		CartService: services.CartServiceFunc(cartRepo),
	}
	paymentRepo := repository.PaymentRepoFunc()
	paymentHandler = PaymentHandler{
		PaymentService: services.PaymentServiceFunc(paymentRepo),
	}

	//Custom Logger - Logs actions to 'checkoutServiceLogs.log' file
	file, err := os.OpenFile("./logs/checkoutServiceLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configure - ShippingAddress Server and Router

	//PORT := os.Getenv("CHECKOUT_SERVICE_PORT")
	router := setupRouter()

	configureSwaggerDoc()
	err = router.Run(":" + "9002")
	if err != nil {
		return
	}
}
