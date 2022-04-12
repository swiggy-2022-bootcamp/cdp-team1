package routes

import (
	"github.com/gin-gonic/gin"
)

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/productsAdmin/app/handlers"
	_ "qwik.in/productsAdmin/docs"
)

func InitRoutes(router *gin.Engine, controller handlers.ProductController) {
	newRouter := router.Group("products/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/health", handlers.HealthCheck)

	newRouter.POST("/", controller.AddProduct)
	newRouter.GET("/", controller.GetProduct)
	newRouter.PUT("/:id", controller.UpdateProduct)
	newRouter.DELETE("/:id", controller.DeleteProduct)
	newRouter.GET("/search/:query", controller.SearchProduct)
}
