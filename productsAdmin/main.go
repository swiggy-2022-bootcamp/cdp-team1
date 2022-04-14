package main

import "qwik.in/productsAdmin/app"

// @title          Swiggy Qwik - Product_Admin module
// @version        1.0
// @description    This microservice is for product_admin service.
// @contact.name   Akash Yadav
// @contact.email  akash283y@gmail.com
// @license.name  Apache 2.0
// @host      localhost:9000
// @BasePath /products/api
func main() {
	app.Start()
	//
	//// create an aws session
	//sess := session.Must(session.NewSession(&aws.Config{
	//	Region:   aws.String("us-east-1"),
	//	Endpoint: aws.String("https://dynamodb.us-east-1.amazonaws.com"),
	//}))
	//
	//// create a dynamodb instance
	//db := dynamodb.New(sess)

}
