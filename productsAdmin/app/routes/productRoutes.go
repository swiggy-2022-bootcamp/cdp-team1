package routes

import (
	"github.com/gin-gonic/gin"
	"qwik.in/productsAdmin/app/handlers"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	prometheusUtility "github.com/swiggy-2022-bootcamp/cdp-team1/common-utilities/prometheus-utility"
	_ "qwik.in/productsAdmin/docs"
)

func InitRoutes(router *gin.Engine, handler handlers.ProductHandler) {
	newRouter := router.Group("api/admin/products")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/health", handlers.HealthCheck)

	newRouter.POST("/", handler.AddProduct)
	newRouter.GET("/", handler.GetProduct)
	newRouter.PUT("/:id", handler.UpdateProduct)
	newRouter.DELETE("/:id", handler.DeleteProduct)
	newRouter.GET("/search", handler.SearchProduct)

	newRouter.GET("/quantity/:id", handler.GetQuantityForProductId)

	newRouter.GET("/metrics", prometheusUtility.PrometheusHandler())
}
