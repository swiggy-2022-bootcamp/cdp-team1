package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
	"qwik.in/customers-admin/domain/service"
	"qwik.in/customers-admin/internal/errors"
)

var customerService service.CustomerServiceInterface

//inject customerService with normal repo (non-mock repo)
func init() {
	customerService = service.InitCustomerService(&repository.CustomerRepository{})
}

// CreateCustomer godoc
// @Summary Create Customer
// @Description Lets admin create customer.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param req body model.Customer true "New customer to add"
// @Success	200  {object} model.Customer
// @Router /customer [POST]
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

// GetCustomerById godoc
// @Summary Get Customer by ID
// @Description Lets admin get customer by customer id.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param customerId path string true "customer id"
// @Success	200  {object} model.Customer
// @Router /customer/{customerId} [GET]
func GetCustomerById(c *gin.Context) {
	fetchedCustomer, err := customerService.GetCustomerById(c.Param("customerId"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedCustomer)
}

// GetCustomerByEmail godoc
// @Summary Get Customer by Email
// @Description Lets admin get customer by customer email.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param customerEmail path string true "customer email"
// @Success	200  {object} model.Customer
// @Router /customer/email/{customerEmail} [GET]
func GetCustomerByEmail(c *gin.Context) {
	fetchedCustomer, err := customerService.GetCustomerByEmail(c.Param("customerEmail"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedCustomer)
}

// UpdateCustomer godoc
// @Summary Update Customer
// @Description Lets admin update customer by customer id.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param customerId path string true "customer id"
// @Param req body model.Customer true "Updated customer"
// @Success	200  {object} model.Customer
// @Router /customer/{customerId} [PUT]
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

// DeleteCustomer godoc
// @Summary Delete Customer
// @Description Lets admin delete customer by customer id.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param customerId path string true "customer id"
// @Success	200  {string} deletion successful
// @Router /customer/{customerId} [DELETE]
func DeleteCustomer(c *gin.Context) {
	successMessage, err := customerService.DeleteCustomer(c.Param("customerId"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *successMessage)
}
