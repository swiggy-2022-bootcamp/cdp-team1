package routes

import "github.com/gin-gonic/gin"

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/productsFrontStore/app/handlers"
	_ "qwik.in/productsFrontStore/docs"
)

func InitRoutes(router *gin.Engine, handler handlers.ProductHandler) {

	newRouter := router.Group("api/products/")

	newRouter.GET("/health", handlers.HealthCheck)
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	newRouter.GET("/", handler.GetProducts)
	newRouter.GET("/:id", handler.GetProductById)
	newRouter.GET("/category/:id", handler.GetProductByCategory)
}
