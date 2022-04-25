package routes

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"qwik.in/account-frontstore/app/handlers"
	_ "qwik.in/account-frontstore/docs"
)

func Router(router *gin.Engine) {
	newRouter := router.Group("api/account-frontstore")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)

	accountController := handlers.InitAccountController(nil)
	newRouter.POST("/register", accountController.RegisterAccount)
	newRouter.GET("/account/:accessorId", accountController.GetAccountById)
	newRouter.PUT("/account/:accessorId", accountController.UpdateAccount)
}
