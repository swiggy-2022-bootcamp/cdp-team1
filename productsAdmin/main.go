package main

import (
	prometheusUtility "github.com/swiggy-2022-bootcamp/cdp-team1/common-utilities/prometheus-utility"
	"qwik.in/productsAdmin/app"
)

// @title          Swiggy Qwik - Product_Admin module
// @version        1.0
// @description    This microservice is for product_admin service.
// @contact.name   Akash Yadav
// @contact.email  akash283y@gmail.com
// @license.name  Apache 2.0
// @host      localhost:9119
// @BasePath /api/admin/products/
func main() {
	prometheusUtility.RegisterMetrics()
	app.Start()
}
