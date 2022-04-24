package main

import "qwik.in/payment-mode/app"

// @title          Swiggy Qwik - Payment_Mode module
// @version        1.0
// @description    This microservice is for payment_mode service.
// @contact.name   Aaditya Khetan
// @contact.email  aadityakhetan123@gmail.com
// @license.name  Apache 2.0
// @securityDefinitions.apikey  Bearer Token
// @in                          header
// @name                        Authorization
// @BasePath /api
func main() {
	app.Start()
}
