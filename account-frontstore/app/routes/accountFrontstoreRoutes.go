package routes

import (
	"github.com/gin-gonic/gin"
	"qwik.in/account-frontstore/app/handlers"
)

//Router ..
func Router(router *gin.Engine) {
	newRouter := router.Group("account-frontstore/api")
	//newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)
	newRouter.POST("/register", handlers.RegisterAccount)
	newRouter.GET("/account", handlers.GetAccountById)
	newRouter.PUT("/account", handlers.UpdateAccount)
}
