package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"qwik.in/shipping/domain/services"
	"qwik.in/shipping/domain/tools/logger"
)

//ShippingHandler ..
type ShippingHandler struct {
	ShippingAddrService services.ShippingAddrService
}

//ShippingAddrDTO ..
type ShippingAddrDTO struct {
	UserID            string `json:"user_id" dynamodb:"user_id"`
	ShippingAddressId string `json:"shipping_address_id" dynamodb:"shipping_address_id"`
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
}

// CreateShippingAddrHandlerFunc ..
// @Summary      Creates New Shipping Address
// @Description  Creates a new shipping address for the user.
// @Tags         ShippingAddress Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /newAddress    [post]
func (sh ShippingHandler) CreateShippingAddrHandlerFunc() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		var saDto ShippingAddrDTO
		if err := ctx.BindJSON(&saDto); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		res, err := sh.ShippingAddrService.CreateNewShippingAddress(
			saDto.UserID,
			saDto.ShippingAddressId,
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
			0,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			fmt.Println(err)
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"message": "New Shipping Address Created ✅", "Shipping Address ID": res})
	}
}

// GetShippingAddrHandlerFunc ..
// @Summary      Get Shipping Address By ShippingAddressId
// @Description  Returns shippingAddress from ShippingAddressId
// @Tags         ShippingAddress Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /getAddress/:id    [get]
func (sh ShippingHandler) GetShippingAddrHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		fmt.Println(id)
		res, err := sh.ShippingAddrService.GetShippingAddressById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Record not found"})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"record": res})
	}
}

// GetDefaultShippingAddrHandlerFunc ..
// @Summary      Gets Default Shipping Address
// @Description  Returns default Shipping Address
// @Tags         ShippingAddress Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /existing/   [get]
func (sh ShippingHandler) GetDefaultShippingAddrHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		res, err := sh.ShippingAddrService.GetDefaultShippingAddr(userId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Provided Shipping Address does not exist"})
			return
		}
		logger.Info("Default Address Fetched ✅ ", res)
		ctx.JSON(http.StatusAccepted, res)
	}
}

//GetAllShippingAddrOfUserHandlerFunc ..
func (sh ShippingHandler) GetAllShippingAddrOfUserHandlerFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userId := ctx.Param("id")
		res, err := sh.ShippingAddrService.GetAllShippingAddressOfUser(userId)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Provided Shipping Address does not exist"})
			return
		}
		logger.Info("Default Address Fetched ✅ ", res)
		ctx.JSON(http.StatusAccepted, res)
	}
}
