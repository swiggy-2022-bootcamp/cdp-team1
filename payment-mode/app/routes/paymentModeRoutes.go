package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/payment-mode/app/handlers"
	_ "qwik.in/payment-mode/docs"
)

func InitRoutes(router *gin.Engine) {
	newRouter := router.Group("payment-mode/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)
}
