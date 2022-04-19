package routes

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"qwik.in/customers-admin/app/handlers"
	_ "qwik.in/customers-admin/docs"
)

func Router(router *gin.Engine) {
	newRouter := router.Group("customer-admin/api")
	newRouter.Use(cors.AllowAll())

	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	newRouter.GET("/", handlers.HealthCheck)
	newRouter.POST("/customer", handlers.CreateCustomer)
	newRouter.GET("/customer/:customerId", handlers.GetCustomerById)
	newRouter.GET("/customer/email/:customerEmail", handlers.GetCustomerByEmail)
	newRouter.PUT("/customer/:customerId", handlers.UpdateCustomer)
	newRouter.DELETE("/customer/:customerId", handlers.DeleteCustomer)
	newRouter.GET("/user", handlers.GetAdminUsers)
	newRouter.POST("/user", handlers.AddAdminUser)

}
