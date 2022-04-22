package app

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/checkout/domain/services"
	"qwik.in/checkout/domain/tools/logger"
)

type CartHandler struct {
	CartService services.CartService
}

type CartDTO struct {
	Id               string `json:"id,omitempty" dynamodbav:"id"`
	CustomerId       string `json:"customer_id,omitempty" dynamodbav:"customer_id"`
	CartProductModel []struct {
		ProductId string `json:"product_id,omitempty" dynamodbav:"product_id"`
		Quantity  int    `json:"quantity" dynamodbav:"quantity"`
	} `json:"products" dynamodbav:"products"`
}

// GetCartDetailsFunc ..
// @Summary      Gets Default Shipping Address
// @Description  Returns default Shipping Address
// @Tags         ShippingAddress Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /existing/   [get]
func (cartH CartHandler) GetCartDetailsFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := cartH.CartService.GetCartDetails()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Provided Shipping Address does not exist"})
			return
		}
		logger.Info("Default Address Fetched ✅ ", res)
		ctx.JSON(http.StatusAccepted, res)
	}
}
