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

	userId := c.Param("userId") //To be replaced with grpc call to auth service.

	var paymentMode models.PaymentMode
	if err := c.ShouldBindJSON(&paymentMode); err != nil {
		c.Error(err)
		err_ := apperrors.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
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
