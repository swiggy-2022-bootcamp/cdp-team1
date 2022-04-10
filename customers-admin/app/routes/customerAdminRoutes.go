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
}
