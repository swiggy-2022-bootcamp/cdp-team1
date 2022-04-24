package main

import "github.com/ashwin2125/qwk/shipping/app"

// @title Gin Swagger Example API
// @version 2.0
// @description ShippingAddress Service
// @termsOfService https://swagger.io/terms/

// @contact.name Ashwin Gopalsamy
// @contact.email ashwinyaal@gmail.com

// @securityDefinitions.apikey  Bearer Token
// @in                          header
// @name                        Authorization
// @BasePath /
// @schemes http
func main() {
	app.Start()
}
