package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/transaction/app/handlers"
	_ "qwik.in/transaction/docs"
)

func InitRoutes(router *gin.Engine) {
	newRouter := router.Group("transaction/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)
}
