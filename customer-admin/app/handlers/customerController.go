package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
	"qwik.in/customers-admin/domain/service"
	"qwik.in/customers-admin/internal/errors"
)

type CustomerControllerInterface interface {
	CreateCustomer(c *gin.Context)
	GetCustomerById(c *gin.Context)
	GetCustomerByEmail(c *gin.Context)
	UpdateCustomer(c *gin.Context)
	DeleteCustomer(c *gin.Context)
}

type CustomerController struct {
	customerService service.CustomerServiceInterface
}

func InitCustomerController(customerServiceToBeInjected service.CustomerServiceInterface) CustomerControllerInterface {
	customerController := new(CustomerController)
	if customerServiceToBeInjected == nil {
		customerController.customerService = service.InitCustomerService(&repository.CustomerRepository{})
	} else {
		customerController.customerService = customerServiceToBeInjected
	}
	return customerController
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
func (customerController *CustomerController) CreateCustomer(c *gin.Context) {
	newCustomer := model.Customer{}
	json.NewDecoder(c.Request.Body).Decode(&newCustomer)
	createdCustomer, err := customerController.customerService.CreateCustomer(newCustomer)

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
func (customerController *CustomerController) GetCustomerById(c *gin.Context) {
	fetchedCustomer, err := customerController.customerService.GetCustomerById(c.Param("customerId"))

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
func (customerController *CustomerController) GetCustomerByEmail(c *gin.Context) {
	fetchedCustomer, err := customerController.customerService.GetCustomerByEmail(c.Param("customerEmail"))

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
func (customerController *CustomerController) UpdateCustomer(c *gin.Context) {
	customer := model.Customer{}
	json.NewDecoder(c.Request.Body).Decode(&customer)
	updatedCustomer, err := customerController.customerService.UpdateCustomer(c.Param("customerId"), customer)

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
func (customerController *CustomerController) DeleteCustomer(c *gin.Context) {
	successMessage, err := customerController.customerService.DeleteCustomer(c.Param("customerId"))

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *successMessage)
}
