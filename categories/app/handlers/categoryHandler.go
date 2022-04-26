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

// UpdateCategory godoc
// @Summary Update Categorys
// @Description Update Category with given id
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [PUT]
func (ct CategoryHandler) UpdateCategory(c *gin.Context) {

	categoryId := c.Param("id")

	var category entity.Category
	category.Category_id = categoryId
	if err := c.BindJSON(&category); err != nil {
		log.Error(err)
	}

	log.Info("Update Category having id : ", categoryId, " with values: ", category)

	err := ct.categoryService.UpdateCategory(categoryId, category)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Category Updated Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}

// DeleteCategory godoc
// @Summary Delete Categorys
// @Description Delete Category with given id
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200
// @Router / [DELETE]
func (ct CategoryHandler) Deletecategory(c *gin.Context) {

	categoryId := c.Param("id")
	log.Info("Delete Category with id : ", categoryId)

	err := ct.categoryService.DeleteCategory(categoryId)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Category Deleted Successfully"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Something went wrong"})
	}
}
// AddCategory godoc
// @Summary AddCategory
// @Description Create a new Category object, generate id and save in DB
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Success	200  {string} 	Category Created
// @Router / [POST]
func (ct CategoryHandler) AddCategory(c *gin.Context) {
	var category entity.Category
	if err := c.BindJSON(&category); err != nil {
		log.Error(err)
	}
	category.SetId()

	log.Info("Add Category with values ", category)

	err := ct.categoryService.CreateCategory(category)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"messaege": "Something went wrong"})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Product Created"})
	}
}