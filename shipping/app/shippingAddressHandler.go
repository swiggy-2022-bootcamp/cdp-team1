package app

import (
	"fmt"
	"github.com/ashwin2125/qwk/shipping/domain/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

//ShippingHandler ..
type ShippingHandler struct {
	ShippingAddressService services.ShippingAddressService
}

//ShippingAddressRecordDTO ..
type ShippingAddressRecordDTO struct {
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

// HandleShippingAddress
// Create Shipping Address
// @Summary      Create Shipping Address
// @Description  Creates New Shipping Address
// @Tags         Shipping Address
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /shippingaddress    [post]
func (sh ShippingHandler) HandleShippingAddress() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var saDto ShippingAddressRecordDTO
		if err := ctx.BindJSON(&saDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := sh.ShippingAddressService.CreateShippingAddress(
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
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Shipping Address Record Added", "Shipping Address ID": res})
	}
}
