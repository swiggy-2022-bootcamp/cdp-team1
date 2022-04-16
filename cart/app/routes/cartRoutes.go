package routes

import (
	"cartService/app/handlers"
	_ "cartService/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type CartRoutes struct {
	cartHandler        handlers.CartHandler
	healthCheckHandler handlers.HealthCheckHandler
}

func NewCartRoutes(cartHandler handlers.CartHandler, healthCheck handlers.HealthCheckHandler) CartRoutes {
	return CartRoutes{cartHandler: cartHandler, healthCheckHandler: healthCheck}
}

func (cr CartRoutes) InitRoutes(newRouter *gin.RouterGroup) {

	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", cr.healthCheckHandler.HealthCheck)

	newRouter.POST("/cart", cr.cartHandler.CreateCart)
	newRouter.PUT("/cart", cr.cartHandler.UpdateCart)
	newRouter.GET("/cart", cr.cartHandler.GetCart)
	newRouter.DELETE("/cart/:id", cr.cartHandler.DeleteCart)
	newRouter.DELETE("/cart", cr.cartHandler.DeleteCartAll)

}
