package main

import (
	prometheusUtility "github.com/swiggy-2022-bootcamp/cdp-team1/common-utilities/prometheus-utility"
	"qwik.in/transaction/app"
)

// @title          Swiggy Qwik - Transaction module
// @version        1.0
// @description    This microservice is for transaction service.
// @contact.name   Aaditya Khetan
// @contact.email  aadityakhetan123@gmail.com
// @license.name  Apache 2.0
// @BasePath /api
func main() {
	prometheusUtility.RegisterMetrics()
	app.Start()
}
