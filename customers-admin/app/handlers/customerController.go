package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/service"
	"qwik.in/customers-admin/internal/errors"
)

var customerService service.CustomerServiceInterface

func init() {
	customerService = &service.CustomerService{}
}

func CreateCustomer(c *gin.Context) {
	newCustomer := model.Customer{}
	json.NewDecoder(c.Request.Body).Decode(&newCustomer)
	createdCustomer, err := customerService.CreateCustomer(newCustomer)

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *createdCustomer)
}

func GetCustomerById(c *gin.Context) {
	fetchedCustomer, err := customerService.GetCustomerById(c.Param("customerId"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedCustomer)
}

func GetCustomerByEmail(c *gin.Context) {
	fetchedCustomer, err := customerService.GetCustomerByEmail(c.Param("customerEmail"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedCustomer)
}

func UpdateCustomer(c *gin.Context) {
	customer := model.Customer{}
	json.NewDecoder(c.Request.Body).Decode(&customer)
	updatedCustomer, err := customerService.UpdateCustomer(c.Param("customerId"), customer)

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *updatedCustomer)
}

func DeleteCustomer(c *gin.Context) {
	successMessage, err := customerService.DeleteCustomer(c.Param("customerId"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *successMessage)
}
