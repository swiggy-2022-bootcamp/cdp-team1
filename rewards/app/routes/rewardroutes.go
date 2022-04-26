package routes

import (
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/rewards/app/handlers"

	_ "qwik.in/rewards/docs"
)

func InitRoutes(router *gin.Engine, handler handlers.RewardHandler) {
	newRouter := router.Group("rewards/api")
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/health", handlers.HealthCheck)
	newRouter.GET("/search/:id", handler.Searchreward)
	newRouter.GET("/", handler.Getall)
	newRouter.PUT("/:id", handler.UpdateReward)
	newRouter.DELETE("/:id", handler.DeleteReward)
	newRouter.POST("/", handler.AddReward)

}
