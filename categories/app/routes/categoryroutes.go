package routes

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/categories/app/handlers"

	_ "qwik.in/categories/docs"
)

func InitRoutes(router *gin.Engine, handler handlers.CategoryHandler) {
	newRouter := router.Group("api/categories")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/health", handlers.HealthCheck)
	newRouter.GET("/", handler.Getall)
	newRouter.PUT("/:id", handler.UpdateCategory)
	newRouter.GET("/search/:id", handler.Searchcategory)
	newRouter.DELETE("/:id", handler.Deletecategory)
	newRouter.POST("/categories", handler.AddCategory)
}
