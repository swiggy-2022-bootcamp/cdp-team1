package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"qwik.in/account-frontstore/app/handlers"
	_ "qwik.in/account-frontstore/docs"
)

//Router ..
func Router(router *gin.Engine) {
	newRouter := router.Group("account-frontstore/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)
	newRouter.POST("/register", handlers.RegisterAccount)
	newRouter.GET("/account/:accessorId", handlers.GetAccountById)
	newRouter.PUT("/account/:accessorId", handlers.UpdateAccount)
}
