package routes

import (
	"authService/domain"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"authService/app/handlers"
	"authService/docs"
)

func RegisterAuthRoutes(authRouter *gin.Engine, authSvc domain.AuthService) {

	authHandlers := handlers.AuthHandlers{AuthSvc: authSvc}

	apiRouter := authRouter.Group("/api")

	docs.SwaggerInfo.BasePath = "/api"
	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	authRoutesGroup := apiRouter.Group("/auth")

	authRoutesGroup.GET("/", authHandlers.HealthCheckHandler)
	authRoutesGroup.POST("/login", authHandlers.LoginHandler)
	authRoutesGroup.POST("/logout", authHandlers.LogoutHandler)
	authRoutesGroup.POST("/verify", authHandlers.VerificationHandler)
}
