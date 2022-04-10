package app

import (
	"io"
	"orderService/app/routes"
	"orderService/log"
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {

	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err == nil {
		gin.DefaultWriter = io.MultiWriter(file)
	}

	// dbClient := db.NewDbClient()
	// orderRepo := db.NewOrderRepositoryDB(dbClient)
	// orderService := domain.NeworderService(orderRepo)
	// orderHandlers := orderHandlers{service: orderService}

	orderRouter := gin.New()
	orderRouter.Use(log.UseLogger(log.DefaultLoggerFormatter), gin.Recovery())
	routes.InitRoutes(orderRouter)
	orderRouter.Run(":8000")

	// apiRouter := orderRouter.Group("/api")
	// apiRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// orderRoutesGroup := apiRouter.Group("/orders")
	// orderRoutesGroup.GET("/", orderHandlers.HelloWorldHandler)
	// orderRoutesGroup.POST("/", orderHandlers.Register)
	// orderRoutesGroup.GET("/:orderId", orderHandlers.GetorderById)
	// orderRoutesGroup.PUT("/:orderId", orderHandlers.Updateorder)
	// orderRoutesGroup.DELETE("/:orderId", orderHandlers.Deleteorder)
	// orderRouter.Run(":8000")
	// logger.Info(fmt.Sprintf("Starting server on %s:%s ...", "127.0.0.1", "8081"))

}
