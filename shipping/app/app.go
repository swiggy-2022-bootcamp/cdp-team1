package app

import (
	"github.com/ashwin2125/qwk/shipping/docs"
	"github.com/ashwin2125/qwk/shipping/domain/repository"
	"github.com/ashwin2125/qwk/shipping/domain/services"
	"github.com/ashwin2125/qwk/shipping/domain/tools/logger"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var shippingHandler ShippingHandler

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.SetTrustedProxies(nil)
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
	file, err := os.OpenFile("./logs/shippingAddressServiceLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	//Configure - ShippingAddress Server and Router

	PORT := os.Getenv("SHIPPING_SERVICE_PORT")
	router := setupRouter()

	configureSwaggerDoc()
	router.Run(":" + PORT)
}
