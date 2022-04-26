package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//HealthCheckRouter ..
func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/api/shipping/", HealthCheck())
}

//ShippingRouter ..
func ShippingRouter(router *gin.Engine) {
	router.GET("/api/shipping/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//Actual Endpoints
	router.POST("/api/shipping/newAddress", shippingHandler.CreateShippingAddrHandlerFunc())
	router.GET("/api/shipping/getAddress/:id", shippingHandler.GetShippingAddrHandlerFunc())
	router.GET("/api/shipping/existing/:id", shippingHandler.GetDefaultShippingAddrHandlerFunc())
	router.GET("/api/shipping/allAddressOfCustomer/:id", shippingHandler.GetAllShippingAddrOfUserHandlerFunc())
}
