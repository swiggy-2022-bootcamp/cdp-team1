package app

import (
	"cartService/utils/logger"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Start() {

	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	// dbClient := db.NewDbClient()
	// cartRepo := db.NewcartRepositoryDB(dbClient)
	// cartService := domain.NewcartService(cartRepo)
	// cartHandlers := cartHandlers{service: cartService}

	cartRouter := gin.Default()
	apiRouter := cartRouter.Group("/api")

	apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// cartRoutesGroup := apiRouter.Group("/carts")
	// cartRoutesGroup.GET("/", cartHandlers.HelloWorldHandler)
	// cartRoutesGroup.POST("/", cartHandlers.Register)
	// cartRoutesGroup.GET("/:cartId", cartHandlers.GetcartById)
	// cartRoutesGroup.PUT("/:cartId", cartHandlers.Updatecart)
	// cartRoutesGroup.DELETE("/:cartId", cartHandlers.Deletecart)

	cartRouter.Run(":8081")
	logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", "8081"))

}
