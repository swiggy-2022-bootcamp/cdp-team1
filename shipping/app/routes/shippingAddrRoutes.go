package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/shipping/app/handlers"
)

//Router ..

//ShippingAddrRoutes ..
type ShippingAddrRoutes struct {
	healthCheckHandlerObj  handlers.HealthCheckHandler
	shippingAddressHandler handlers.ShippingAddrHandler
}

//ShippingAddrRoutesFunc ..
func ShippingAddrRoutesFunc(healthCheckHandlerObj handlers.HealthCheckHandler, shippingAddressHandler handlers.ShippingAddrHandler) ShippingAddrRoutes {
	return ShippingAddrRoutes{healthCheckHandlerObj: healthCheckHandlerObj, shippingAddressHandler: shippingAddressHandler}
}

//InitRoutes ..
func (sar ShippingAddrRoutes) InitRoutes(shippingAddrRouter *gin.RouterGroup) {
	shippingAddrRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	shippingAddrRouter.GET("/", sar.healthCheckHandlerObj.HealthCheck)

	shippingAddrRouter.POST("/newAddress", sar.shippingAddressHandler.AddShippingAddress)
}
