package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/productsFrontStore/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return ProductHandler{productService: productService}
}

// GetProducts godoc
// @Summary Get Products
// @Description Get a list of all products
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (p ProductHandler) GetProducts(c *gin.Context) {
	products, err := p.productService.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// GetProductById godoc
// @Summary GetProductById
// @Description Get Product By Id
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (p ProductHandler) GetProductById(c *gin.Context) {
	productId := c.Param("id")
	product, err := p.productService.GetProductById(productId)
	if err == nil {
		c.JSON(http.StatusOK, product)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// GetProductByCategory godoc
// @Summary GetProductByCategory
// @Description Get all products belonging to a category
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (p ProductHandler) GetProductByCategory(c *gin.Context) {
	categoryId := c.Param("id")
	products, err := p.productService.GetProductsByCategory(categoryId)
	if err == nil {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}
