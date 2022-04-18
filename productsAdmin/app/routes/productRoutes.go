package routes

import (
	"github.com/gin-gonic/gin"
	"qwik.in/productsAdmin/app/handlers"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "qwik.in/productsAdmin/docs"
)

func InitRoutes(router *gin.Engine, handler handlers.ProductHandler) {
	newRouter := router.Group("products/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/health", handlers.HealthCheck)

	newRouter.POST("/", handler.AddProduct)
	newRouter.GET("/", handler.GetProduct)
	newRouter.PUT("/:id", handler.UpdateProduct)
	newRouter.DELETE("/:id", handler.DeleteProduct)
	newRouter.GET("/search", handler.SearchProduct)

	newRouter.GET("/quantity/:id", handler.GetQuantityForProductId)
}
