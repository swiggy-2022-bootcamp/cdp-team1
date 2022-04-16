package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"qwik.in/categories/entity"
	"qwik.in/categories/log"
	"qwik.in/categories/service"
)

type CategoryHandler struct {
	categoryService service.CategoryService
}

func NewCategoryHandler(categoryService service.CategoryService) CategoryHandler {
	return CategoryHandler{categoryService: categoryService}
}

// Getall godoc
// @Summary Get Rewards
// @Description Get a list of all Rewards
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (ct CategoryHandler) Getall(c *gin.Context) {
	rewards, err := ct.categoryService.GetAll()
	if err == nil {
		c.JSON(http.StatusOK, rewards)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// Searchreward godoc
// @Summary Search rewards
// @Description Search reward with given query
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [GET]
func (ct CategoryHandler) Searchcategory(c *gin.Context) {

	id := c.Param("id")

	log.Info("Searching category with id: ", id)
	category, err := ct.categoryService.SearchCategory(id)
	if err == nil {
		c.JSON(http.StatusOK, category)
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
func (ct CategoryHandler) UpdateProduct(c *gin.Context) {

	categoryId := c.Param("id")

	var category entity.Category
	if err := c.BindJSON(&category); err != nil {
		log.Error(err)
	}

	log.Info("Update product having id : ", categoryId, " with values: ", category)

	err := ct.categoryService.UpdateCategory(categoryId, category)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Product Updated Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}
