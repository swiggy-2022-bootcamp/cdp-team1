package main

import (
	"fmt"
	"qwik.in/productsAdmin/api/handler"
)

func init() {
	fmt.Println("ProductAdmin service started")
}

func main() {

	handler.Start()
}
