package routes

import (
	"github.com/gin-gonic/gin"
	"qwik.in/customers-admin/app/handlers"
)

//Router ..
func Router(router *gin.Engine) {
	newRouter := router.Group("customer-admin/api")
	//newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)
	newRouter.POST("/customer", handlers.CreateCustomer)
	newRouter.GET("/customer/:customerId", handlers.GetCustomerById)
	newRouter.GET("/customer/email/:customerEmail", handlers.GetCustomerByEmail)
	newRouter.PUT("/customer/:customerId", handlers.UpdateCustomer)
	newRouter.DELETE("/customer/:customerId", handlers.DeleteCustomer)
	newRouter.GET("/user", handlers.GetAdminUsers)
	newRouter.POST("/user", handlers.AddAdminUser)

}
