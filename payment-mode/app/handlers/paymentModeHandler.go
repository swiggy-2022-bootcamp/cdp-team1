package handlers

import (
	"github.com/gin-gonic/gin"
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
	var userPaymentMode models.UserPaymentMode
	if err := c.ShouldBindJSON(&userPaymentMode); err != nil {
		c.Error(err)
		err_ := apperrors.NewBadRequestError(err.Error())
		c.JSON(err_.Code, err_.Message)
		return
	}
	err := ph.paymentService.AddPaymentMode(&userPaymentMode)
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
