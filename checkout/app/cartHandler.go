package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "qwik.in/checkout/docs" //GoSwagger
	"qwik.in/checkout/domain/services"
	"qwik.in/checkout/domain/tools/logger"
)

// CartHandler ..
type CartHandler struct {
	CartService services.CartService
}

//CartDTO ..
type CartDTO struct {
	ID               string `json:"id,omitempty" dynamodbav:"id"`
	CustomerID       string `json:"customer_id,omitempty" dynamodbav:"customer_id"`
	CartProductModel []struct {
		ProductID string `json:"product_id,omitempty" dynamodbav:"product_id"`
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
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Cart Details does not exist"})
			return
		}
		logger.Info("✅ Cart Details Fetched ", res)
		ctx.JSON(http.StatusAccepted, res)
	}
}
