package route

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	handler "qwik.in/checkout/app/handlers"
)

//Router ..
func Router(router *gin.Engine) {
	newRouter := router.Group("checkout/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handler.HealthCheck)
}
