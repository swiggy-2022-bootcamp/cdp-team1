package main

import "authService/app"

// @title Auth Service APIs
// @version 1.0
// @description Swagger API for Golang Project.
// @termsOfService http://swagger.io/terms/
// @contact.name Ayan Dutta
// @contact.email ayan59dutta@gmail.com
// @securityDefinitions.apikey  Bearer Token
// @in                          header
// @name                        Authorization

// @BasePath /api
func main() {
	app.Start()
}
