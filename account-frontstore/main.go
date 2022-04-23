package main

import (
	"qwik.in/account-frontstore/app"
	_ "qwik.in/account-frontstore/docs"
)

// @title          Swiggy Qwik - Account-FrontStore module
// @version        1.0
// @description    This microservice is for Account-FrontStore.
// @contact.name   Ravikumar S
// @contact.email  ravikumarsravi1999@gmail.com
// @license.name  Apache 2.0
// @host      localhost:7001
// @BasePath /api/account-frontstore
func main() {
	app.Start()
}
