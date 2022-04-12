package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
)

// AddProduct godoc
// @Summary Add new product to store
// @Description Create a new product object, generate id and save in DB
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	Product Created
// @Router / [POST]
func AddProduct(c *gin.Context) {

	var product entity.Product

	if err := c.BindJSON(&product); err != nil {
		log.Error(err)
	}
	product.SetId()

	log.Info(product)

	c.JSON(http.StatusOK, gin.H{"message": "Product Created"})
}
