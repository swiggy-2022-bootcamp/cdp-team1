package routes

import "github.com/gin-gonic/gin"

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/productsAdmin/app/handlers"
	_ "qwik.in/productsAdmin/docs"
)

func InitRoutes(router *gin.Engine) {
	newRouter := router.Group("products/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)
}
