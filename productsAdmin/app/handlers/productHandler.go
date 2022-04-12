package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
	"qwik.in/productsAdmin/service"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) ProductController {
	return ProductController{productService: productService}
}

// AddProduct godoc
// @Summary Add new product to store
// @Description Create a new product object, generate id and save in DB
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	Product Created
// @Router / [POST]
func (p ProductController) AddProduct(c *gin.Context) {

	var product entity.Product
	if err := c.BindJSON(&product); err != nil {
		log.Error(err)
	}
	product.SetId()
	log.Info("Parsed product: ", product)
	err := p.productService.CreateProduct(product)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"messaege": "Something went wrong"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Product Created"})
	}
}

// GetProduct godoc
// @Summary Get Products
// @Description Get a list of all products
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (p ProductController) GetProduct(c *gin.Context) {
	products, err := p.productService.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// UpdateProduct godoc
// @Summary Update Products
// @Description Update Product with given id
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [PUT]
func (p ProductController) UpdateProduct(c *gin.Context) {

	productId := c.Param("id")
	log.Info("Update product with id : ", productId)

	// TODO check if exist

	var product entity.Product
	if err := c.BindJSON(&product); err != nil {
		log.Error(err)
	}
	log.Info("Parsed product: ", product)
	err := p.productService.UpdateProduct(product)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Product Updated Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// DeleteProduct godoc
// @Summary Delete Products
// @Description Delete Product with given id
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [DELETE]
func (p ProductController) DeleteProduct(c *gin.Context) {

	productId := c.Param("id")
	log.Info("Delete product with id : ", productId)

	// TODO check if exist

	err := p.productService.DeleteProduct("")
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Product Deleted Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// SearchProduct godoc
// @Summary Search Products
// @Description Search Product with given query
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (p ProductController) SearchProduct(c *gin.Context) {

	query := c.Param("query")
	log.Info("Find products with query : ", query)

	products, err := p.productService.SearchProduct(query)
	if err == nil {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}
