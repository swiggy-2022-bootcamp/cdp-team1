package main

import (
	"qwik.in/customers-admin/app"
	_ "qwik.in/customers-admin/docs"
)

// @title          Swiggy Qwik - Customer-Admin module
// @version        1.0
// @description    This microservice is for Customer-Admin.
// @contact.name   Ravikumar S
// @contact.email  ravikumarsravi1999@gmail.com
// @license.name  Apache 2.0
// @host      localhost:7002
// @BasePath /customer-admin/api
func main() {
	app.Start()
}
