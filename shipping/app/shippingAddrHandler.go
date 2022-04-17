package app

import (
	"fmt"
	"github.com/ashwin2125/qwk/shipping/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//ShippingHandler ..
type ShippingHandler struct {
	ShippingAddrService services.ShippingAddrService
}

//ShippingAddrDTO ..
type ShippingAddrDTO struct {
	UserID         int    `json:"user_id" dynamodb:"user_id"`
	FirstName      string `json:"first_name" dynamodbav:"first_name"`
	LastName       string `json:"last_name" dynamodbav:"last_name"`
	AddressLine1   string `json:"address_line_1" dynamodbav:"address_line_1" `
	AddressLine2   string `json:"address_line_2" dynamodbav:"address_line_2"`
	City           string `json:"city" dynamodbav:"city"`
	State          string `json:"state" dynamodbav:"state"`
	Phone          string `json:"phone" dynamodbav:"phone"`
	Pincode        int    `json:"pincode" dynamodbav:"pincode"`
	AddressType    string `json:"address_type" dynamodbav:"address_type"`
	DefaultAddress bool   `json:"default_address" dynamodbav:"default_address"`
}

// ShippingAddrHandlerFunc
// Creates New Shipping Address
// @Summary      Creates New Shipping Address
// @Description  Creates a shipping address for the user.
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /newAddress    [post]
func (sh ShippingHandler) ShippingAddrHandlerFunc() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var saDto ShippingAddrDTO
		if err := ctx.BindJSON(&saDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := sh.ShippingAddrService.CreateNewShippingAddress(
			saDto.UserID,
			saDto.FirstName,
			saDto.LastName,
			saDto.AddressLine1,
			saDto.AddressLine2,
			saDto.City,
			saDto.State,
			saDto.Phone,
			saDto.Pincode,
			saDto.AddressType,
			saDto.DefaultAddress,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			fmt.Println("Sucks here!")
			fmt.Println(err)
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "New Shipping Address Created ✅", "Shipping Address ID": res})
	}
}
