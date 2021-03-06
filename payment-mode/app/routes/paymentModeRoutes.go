package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	prometheusUtility "github.com/swiggy-2022-bootcamp/cdp-team1/common-utilities/prometheus-utility"
	"qwik.in/payment-mode/app/handlers"
	_ "qwik.in/payment-mode/docs"
)

type PaymentRoutes struct {
	paymentHandler     handlers.PaymentHandler
	healthCheckHandler handlers.HealthCheckHandler
}

func NewPaymentRoutes(paymentHandler handlers.PaymentHandler, healthCheck handlers.HealthCheckHandler) PaymentRoutes {
	return PaymentRoutes{paymentHandler: paymentHandler, healthCheckHandler: healthCheck}
}

func (pr PaymentRoutes) InitRoutes(newRouter *gin.RouterGroup) {

	//Swagger route
	newRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//HealthCheck Route
	newRouter.GET("/paymentmethods/health", pr.healthCheckHandler.HealthCheck)

	//Prometheus metrics route
	newRouter.GET("/paymentmethods/metrics", prometheusUtility.PrometheusHandler())

	//Payment mode service routes
	newRouter.POST("/paymentmethods", pr.paymentHandler.AddPaymentMode)
	newRouter.GET("/paymentmethods", pr.paymentHandler.GetPaymentMode)
	newRouter.POST("/pay", pr.paymentHandler.CompletePayment)
}
