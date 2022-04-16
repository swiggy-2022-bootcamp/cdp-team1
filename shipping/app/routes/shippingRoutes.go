package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/shipping/app/handlers"
)

//Router ..

//ShippingAddressRoutes ..
type ShippingAddressRoutes struct {
	healthCheckHandler handlers.HealthCheckHandler
}

//NewShippingAddressRoutes ..
func NewShippingAddressRoutes(healthCheckHandler handlers.HealthCheckHandler) ShippingAddressRoutes {
	return ShippingAddressRoutes{healthCheckHandler: healthCheckHandler}
}

//InitRoutes ..
func (sar ShippingAddressRoutes) InitRoutes(shippingAddressRouter *gin.RouterGroup) {
	shippingAddressRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	shippingAddressRouter.GET("/", sar.healthCheckHandler.HealthCheck)
}
