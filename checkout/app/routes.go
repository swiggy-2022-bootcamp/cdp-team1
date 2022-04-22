package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "qwik.in/checkout/docs"
)

//HealthCheckRouter ..
func HealthCheckRouter(gin *gin.Engine) {
	gin.GET("/checkout/api/", HealthCheck())
}

//SwagHandler ..
func SwagHandler(router *gin.Engine) {
	router.GET("/shipping/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

//ShippingRouter ..
func ShippingRouter(router *gin.Engine) {
	//Actual Endpoints
	router.GET("/shipping/api/existing", shippingHandler.GetDefaultShippingAddrHandlerFunc())
}

//CartRouter ..
func CartRouter(router *gin.Engine) {
	router.GET("/cart/api/cartItems", cartHandler.GetCartDetailsFunc())
}

//CheckoutRouter ..
func CheckoutRouter(router *gin.Engine) {
	router.GET("/checkout/api/shippingAddress", checkoutHandler.CheckoutGetShippingAddressFlow())
	router.GET("/checkout/api/cartItems", checkoutHandler.CheckoutGetCartItemsFlow())
}
