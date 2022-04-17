package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/shipping/internal/models"
	"qwik.in/shipping/internal/services"
	"qwik.in/shipping/internal/tools/err"
)

type ShippingAddrHandler struct {
	shippingAddrService services.ShippingAddrService
}

func ShippingAddrHandlerFunc(shippingAddrService services.ShippingAddrService) ShippingAddrHandler {
	return ShippingAddrHandler{shippingAddrService: shippingAddrService}
}

func (ph ShippingAddrHandler) AddShippingAddress(c *gin.Context) {

	var shippingAddress models.ShippingAddressDetails
	if err := c.BindJSON(&shippingAddress); err != nil {
		c.Error(err)
		err_ := apperrors.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	err := ph.shippingAddrService.AddShippingAddress(&shippingAddress)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	defer c.JSON(http.StatusOK, gin.H{"message": "New Shipping Address Created & saved successfully! 📧"})
}
