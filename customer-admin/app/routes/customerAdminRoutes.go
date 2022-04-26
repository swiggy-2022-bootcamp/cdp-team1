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
	newRouter := router.Group("api/customer-admin")
	newRouter.Use(cors.AllowAll())

	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	newRouter.GET("/", handlers.HealthCheck)

	customerController := handlers.InitCustomerController(nil)
	newRouter.POST("/customer", customerController.CreateCustomer)
	newRouter.GET("/customer/:customerId", customerController.GetCustomerById)
	newRouter.GET("/customer/email/:customerEmail", customerController.GetCustomerByEmail)
	newRouter.PUT("/customer/:customerId", customerController.UpdateCustomer)
	newRouter.DELETE("/customer/:customerId", customerController.DeleteCustomer)

	newRouter.GET("/user", handlers.GetAdminUsers)
	newRouter.POST("/user", handlers.AddAdminUser)

}
