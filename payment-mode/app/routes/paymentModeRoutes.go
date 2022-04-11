package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"qwik.in/payment-mode/app/handlers"
	_ "qwik.in/payment-mode/docs"
)

type PaymentRoutes struct {
	paymentHandler handlers.PaymentHandler
}

func NewPaymentRoutes(paymentHandler handlers.PaymentHandler) PaymentRoutes {
	return PaymentRoutes{paymentHandler: paymentHandler}
}
func (pr PaymentRoutes) InitRoutes(newRouter *gin.RouterGroup) {

	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	newRouter.GET("/", handlers.HealthCheck)

	newRouter.POST("/paymentmethods/:userId", pr.paymentHandler.AddPaymentMode)
	newRouter.GET("/paymentmethods/:userId", pr.paymentHandler.GetPaymentMode)
}
