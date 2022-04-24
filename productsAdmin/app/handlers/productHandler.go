package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"qwik.in/productsAdmin/entity"
	"qwik.in/productsAdmin/log"
	"qwik.in/productsAdmin/service"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productService service.ProductService) ProductHandler {
	return ProductHandler{productService: productService}
}

// AddProduct godoc
// @Summary AddProduct
// @Description Create a new product object, generate id and save in DB
// @Tags ProductAdmin
// @Schemes
// @Accept json
// @Produce json
// @Param req body entity.Product true "Product details"
// @Success	200 "Product Created"
// @Failure 400 "Something went wrong"
// @Router /api/admin/products [POST]
func (p ProductHandler) AddProduct(c *gin.Context) {
	var product entity.Product
	if err := c.BindJSON(&product); err != nil {
		log.Error(err)
	}

	log.Info("Add product with values ", product)

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
// @Tags ProductAdmin
// @Schemes
// @Accept json
// @Produce json
// @Success	200 {array} entity.Product
// @Failure 400 "Something went wrong"
// @Router /api/admin/products [GET]
func (p ProductHandler) GetProduct(c *gin.Context) {
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
// @Tags ProductAdmin
// @Schemes
// @Accept json
// @Produce json
// @Param productId path string true "Product id"
// @Success	200 "Product Updated Successfully"
// @Failure 400 "Something went wrong"
// @Router /api/admin/products/{id} [PUT]
func (p ProductHandler) UpdateProduct(c *gin.Context) {

	productId := c.Param("id")

	var product entity.Product
	if err := c.BindJSON(&product); err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse product"})
		return
	}

	log.Info("Update product having id : ", productId, " with values: ", product)

	err := p.productService.UpdateProduct(productId, product)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Product Updated Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// DeleteProduct godoc
// @Summary Delete Products
// @Description Delete Product with given id
// @Tags ProductAdmin
// @Schemes
// @Accept json
// @Produce json
// @Param productId path string true "Product id"
// @Success	200 "Product Deleted Successfully"
// @Failure 400 "Something went wrong"
// @Router /api/admin/products/{id} [DELETE]
func (p ProductHandler) DeleteProduct(c *gin.Context) {

	productId := c.Param("id")
	log.Info("Delete product with id : ", productId)

	err := p.productService.DeleteProduct(productId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Product Deleted Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// SearchProduct godoc
// @Summary Search Products
// @Description Search Product with given query
// @Tags ProductAdmin
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Failure 400 "Something went wrong"
// @Router /api/admin/products/search [GET]
func (p ProductHandler) SearchProduct(c *gin.Context) {

	limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
	log.Info("Find products with limit : ", limit)

	products, err := p.productService.SearchProduct(limit)
	if err == nil {
		c.JSON(http.StatusOK, products)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// GetQuantityForProductId godoc
// @Summary Quantity of Products
// @Description Get Quantity of Product with given id
// @Tags ProductAdmin
// @Schemes
// @Accept json
// @Produce json
// @Param productId path string true "Product id"
// @Success	200 {object} proto.Response
// @Failure 400 "Something went wrong"
// @Router /api/admin/products/quantity/{id} [GET]
func (p ProductHandler) GetQuantityForProductId(c *gin.Context) {

	productId := c.Param("id")
	log.Info("Get quantity for product with id : ", productId)

	response, err := p.productService.GetQuantityForProductId(productId)
	if err == nil {
		c.JSON(http.StatusOK, response)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}
