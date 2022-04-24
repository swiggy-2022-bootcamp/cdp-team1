package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"qwik.in/checkout/domain/services"
	"qwik.in/checkout/domain/tools/logger"
)

//PaymentHandler ..
type PaymentHandler struct {
	PaymentService services.PaymentService
}

//PaymentModeDTO ..
type PaymentModeDTO struct {
	OrderAmount      float64 `json:"order_amount" dynamodbav:"order_amount,omitempty"`
	UserID           string  `json:"user_id" dynamodbav:"user_id,omitempty"`
	PaymentModeModel []struct {
		Mode       string `json:"mode_of_payment" dynamodbav:"mode_of_payment"`
		Credential string `json:"credential" dynamodbav:"credential"`
	} `json:"payment_mode" dynamodbav:"payment_mode"`
}

// GetDefaultPaymentModeHandlerFunc ..
// @Summary      Gets Default Payment Mode
// @Description  Returns default Payment Mode
// @Tags         Payment Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /payment/api/payments   [get]
func (p PaymentHandler) GetDefaultPaymentModeHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res, err := p.PaymentService.GetDefaultPaymentMode()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Payment Mode does not exist"})
			return
		}
		logger.Info(" ✅ Payment Mode Fetched", res)
		ctx.JSON(http.StatusAccepted, res)
	}
}
