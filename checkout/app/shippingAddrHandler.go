package app

import (
	_ "qwik.in/checkout/docs" //GoSwagger
	"qwik.in/checkout/domain/services"
)

//ShippingHandler ..
type ShippingHandler struct {
	ShippingService services.ShippingService
}

//ShippingAddrDTO ..
type ShippingAddrDTO struct {
	UserID            int    `json:"user_id" dynamodb:"user_id"`
	ShippingAddressID string `json:"shipping_address_id" dynamodb:"shipping_address_id"`
	FirstName         string `json:"first_name" dynamodbav:"first_name"`
	LastName          string `json:"last_name" dynamodbav:"last_name"`
	AddressLine1      string `json:"address_line_1" dynamodbav:"address_line_1" `
	AddressLine2      string `json:"address_line_2" dynamodbav:"address_line_2"`
	City              string `json:"city" dynamodbav:"city"`
	State             string `json:"state" dynamodbav:"state"`
	Phone             string `json:"phone" dynamodbav:"phone"`
	Pincode           int    `json:"pincode" dynamodbav:"pincode"`
	AddressType       string `json:"address_type" dynamodbav:"address_type"`
	DefaultAddress    bool   `json:"default_address" dynamodbav:"default_address"`
	ShippingCost      int    `json:"shipping_cost" dynamodbav:"shipping_cost"`
}

//// GetDefaultShippingAddrHandlerFunc ..
//// @Summary      Gets Default Shipping Address
//// @Description  Returns default Shipping Address
//// @Tags         ShippingAddress Service
//// @Produce      json
//// @Success      200  {object}  map[string]interface{}
//// @Failure      400  {number} 	http.StatusBadRequest
//// @Router       /shipping/api/existing   [get]
//func (sh ShippingHandler) GetDefaultShippingAddrHandlerFunc() gin.HandlerFunc {
//	return func(ctx *gin.Context) {
//		res, err := sh.ShippingService.GetDefaultShippingAddr(true)
//		if err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Provided Shipping Address does not exist"})
//			return
//		}
//		logger.Info(" ✅ Default Address Fetched", res)
//		ctx.JSON(http.StatusAccepted, res)
//	}
//}
