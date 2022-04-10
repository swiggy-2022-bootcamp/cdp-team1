package app

import (
	"cartService/app/routes"
	"cartService/log"
	"io"
	"os"

	"github.com/gin-gonic/gin"
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

	cartRouter := gin.New()
	cartRouter.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	routes.InitRoutes(cartRouter)
	cartRouter.Run(":7000")

	// apiRouter := cartRouter.Group("/api")
	// apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// cartRoutesGroup := apiRouter.Group("/carts")
	// cartRoutesGroup.GET("/", cartHandlers.HelloWorldHandler)
	// cartRoutesGroup.POST("/", cartHandlers.Register)
	// cartRoutesGroup.GET("/:cartId", cartHandlers.GetcartById)
	// cartRoutesGroup.PUT("/:cartId", cartHandlers.Updatecart)
	// cartRoutesGroup.DELETE("/:cartId", cartHandlers.Deletecart)
	// logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", "8081"))
}
