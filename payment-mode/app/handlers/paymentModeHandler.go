package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	apperrors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/domain/services"
)

var validate = validator.New()

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) PaymentHandler {
	return PaymentHandler{paymentService: paymentService}
}

// AddPaymentMode godoc
// @Summary To add a new payment method for a user.
// @Description To add a new payment method for a user(COD,Debit card, Credit Card).
// @Tags PaymentMode
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.PaymentMode true "Payment mode details"
// @Param userId path string true "User id"
// @Success	200  string 	Payment mode added successfully
// @Failure 400  string 	Bad request
// @Failure 500  string 	Internal server error
// @Failure 409  string 	Payment mode already exists
// @Router /paymentmethods/{userId} [POST]
func (ph PaymentHandler) AddPaymentMode(c *gin.Context) {

	userId := c.Param("userId") //To be replaced with grpc call to auth service.

	var paymentMode models.PaymentMode
	if err := c.BindJSON(&paymentMode); err != nil {
		c.Error(err)
		err_ := apperrors.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&paymentMode); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	err := ph.paymentService.AddPaymentMode(&paymentMode, userId)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment mode added successfully"})
}

// GetPaymentMode godoc
// @Summary To get available payment modes of a user.
// @Description To get available payment modes of a user.
// @Tags PaymentMode
// @Schemes
// @Accept json
// @Produce json
// @Param userId path string true "User id"
// @Success	200  {object} 	models.UserPaymentMode
// @Failure 500  string 	Internal server error
// @Failure 404  string 	User not found
// @Router /paymentmethods/{userId} [GET]
func (ph PaymentHandler) GetPaymentMode(c *gin.Context) {
	userId := c.Param("userId")
	userPaymentModes, err := ph.paymentService.GetPaymentMode(userId)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, userPaymentModes)
}

// SetPaymentMode godoc
// @Summary To set payment modes for an order.
// @Description To set payment modes for an order.
// @Tags PaymentMode
// @Schemes
// @Accept json
// @Produce json
// @Param userId path string true "User id"
// @Param req body models.PaymentMode true "Payment mode details"
// @Success	200  string 	http.StatusOK
// @Failure 500  string 	Internal server error
// @Failure 404  string 	User not found
// @Router /setpaymentmethods/{userId} [POST]
func (ph PaymentHandler) SetPaymentMode(c *gin.Context) {
	userId := c.Param("userId")

	var paymentMode models.PaymentMode
	if err := c.BindJSON(&paymentMode); err != nil {
		c.Error(err)
		err_ := apperrors.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&paymentMode); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	_, err := ph.paymentService.SetPaymentMode(userId, paymentMode)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment mode verified successfully"})
}

// CompletePayment godoc
// @Summary To set payment modes for an order.
// @Description To set payment modes for an order.
// @Tags PaymentRequest
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.PaymentRequest true "Payment request"
// @Success	201  string 	http.StatusAccepted
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Payment method is not added for the current user.
// @Failure 406  string		Insufficient funds, payment failed.
// @Router /pay [POST]
func (ph PaymentHandler) CompletePayment(c *gin.Context) {
	var paymentRequest models.PaymentRequest
	if err := c.BindJSON(&paymentRequest); err != nil {
		c.Error(err)
		err_ := apperrors.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&paymentRequest); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	_, err := ph.paymentService.CheckBalanceAndCompletePayment(&paymentRequest)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Payment successful"})
}
