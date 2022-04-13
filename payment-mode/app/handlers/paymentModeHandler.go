package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	apperrors "qwik.in/payment-mode/app-errors"
	"qwik.in/payment-mode/domain/models"
	"qwik.in/payment-mode/domain/services"
)

type PaymentHandler struct {
	paymentService services.PaymentService
}

func NewPaymentHandler(paymentService services.PaymentService) PaymentHandler {
	return PaymentHandler{paymentService: paymentService}
}

func (ph PaymentHandler) AddPaymentMode(c *gin.Context) {
	//TODO
	//1. Take paymentMode object
	//2. Fetch userId from JWT token
	//3. Fetch userPaymentMode record from DB.
	//4. Check if passed paymentMode object is there or not in userPaymentMode, if there return Conflict
	//5. If not there add the payment mode object in userPaymentMode's paymentMethods field.

	var userPaymentMode models.UserPaymentMode
	if err := c.ShouldBindJSON(&userPaymentMode); err != nil {
		c.Error(err)
		err_ := apperrors.NewBadRequestError(err.Error())
		c.JSON(err_.Code, err_.Message)
		return
	}

	newUserPaymentMode := models.UserPaymentMode{
		UserId:       uuid.New().String(),
		PaymentModes: userPaymentMode.PaymentModes,
	}

	err := ph.paymentService.AddPaymentMode(&newUserPaymentMode)
	if err != nil {
		c.Error(err)
		err_ := apperrors.NewUnexpectedError(err.Error())
		c.JSON(err_.Code, err_.Message)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment mode added successfully"})
}

func (ph PaymentHandler) GetPaymentMode(c *gin.Context) {
	userId := c.Param("userId")
	userPaymentModes, err := ph.paymentService.GetPaymentMode(userId)
	if err != nil {
		c.Error(err)
		err_ := apperrors.NewUnexpectedError(err.Error())
		c.JSON(err_.Code, err_.Message)
		return
	}
	c.JSON(http.StatusOK, userPaymentModes)
}
