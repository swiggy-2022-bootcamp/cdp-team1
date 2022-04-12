package routes

import (
	"orderService/app/handlers"
	_ "orderService/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type OrderRoutes struct {
	orderHandler handlers.OrderHandler
}

func (or OrderRoutes) InitRoutes(router *gin.Engine) {

	newRouter := router.Group("order/api")

	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)

	newRouter.GET("/orders", or.orderHandler.GetAllOrder)
	newRouter.GET("/orders/status/:status", or.orderHandler.GetOrderByStatus)
	newRouter.GET("/orders/:id", or.orderHandler.GetOrderById)
	newRouter.PUT("orders/:id", or.orderHandler.UpdateOrder)
	newRouter.DELETE("/orders/:id", or.orderHandler.DeleteOrderById)
	newRouter.GET("/orders/users/:id", or.orderHandler.GetOrderByCustomerId)
	newRouter.POST("/orders/invoice/:id", or.orderHandler.CreateInvoice)
}
