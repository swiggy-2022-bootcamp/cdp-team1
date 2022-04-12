package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"authService/app/handlers"
	"authService/docs"
)

func RegisterAuthRoutes(authRouter *gin.Engine) {

	authHandlers := handlers.AuthHandlers{}

	apiRouter := authRouter.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutesGroup := apiRouter.Group("/auth")

	authRoutesGroup.GET("/", authHandlers.HealthCheckHandler)
}
