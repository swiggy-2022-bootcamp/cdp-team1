package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//HealthCheckRouter ..
func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/shipping/api/", HealthCheck())
}

//ShippingRouter ..
func ShippingRouter(router *gin.Engine) {
	router.GET("/shipping/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.POST("/shipping/api/newAddress", shippingHandler.HandleShippingAddress())
}
