package main

import "github.com/ashwin2125/qwk/shipping/app"

// @title Gin Swagger Example API
// @version 2.0
// @description Shipping Service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9003
// @BasePath /
// @schemes http
func main() {
	app.Start()
}
