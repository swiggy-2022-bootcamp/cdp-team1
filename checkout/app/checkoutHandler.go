package app

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"qwik.in/checkout/domain/services"
	"qwik.in/checkout/domain/tools/errs"
	"qwik.in/checkout/domain/tools/logger"
)

//CheckoutHandler
type CheckoutHandler struct {
	CheckoutService services.CheckoutService
}

// CheckoutGetShippingAddressFlow ..
// @Summary      Gets Default Shipping Address
// @Description  Returns default Shipping Address
// @Tags         ShippingAddress Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /checkout/api/shippingaddress [get]
func (ch CheckoutHandler) CheckoutGetShippingAddressFlow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Getting Shipping Address !")
		var shippingAddressRespBody ShippingAddrDTO
		////INITIAL GET REQUEST
		req, err := http.Get("http://localhost:9002/shipping/api/existing")
		//ADDING HEADER FOR REQUEST
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		//HANDLING ERROR / ELSE -> CONSUME REQUEST
		if err != nil {
			logger.Error("Default Shipping Address does not exist")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Shipping Address does not exist"))
			return
		} else {
			data, err2 := ioutil.ReadAll(req.Body)
			if err2 != nil {
				logger.Error("Error during RESPONSE value from GET Request")
				ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Some shit doesn't work during GET - CONSUME"))
			}
			err3 := json.Unmarshal(data, &shippingAddressRespBody)
			if err3 != nil {
				ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Default Shipping Address couldn't be UNMARSHALLED."))
				return
			}
			fmt.Println("Default ShippingAddress Fetched ✅")
			logger.Info("Default ShippingAddress Fetched ✅", shippingAddressRespBody)
			ctx.JSON(http.StatusAccepted, shippingAddressRespBody)
		}
	}
}

// CheckoutGetCartItemsFlow ..
// @Summary      Gets Cart Items
// @Description  Returns Cart Items using GET Request.
// @Tags         Cart Service
// @Produce      json
// @Success      200  {object}  map[string]interface{}
// @Failure      400  {number} 	http.StatusBadRequest
// @Router       /checkout/api/cartItems   [get]
func (ch CheckoutHandler) CheckoutGetCartItemsFlow() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("Getting Cart Items !")
		var cartItemsRespBody CartDTO
		////INITIAL GET REQUEST
		req, err := http.Get("http://localhost:9002/cart/api/cartItems")
		//ADDING HEADER FOR REQUEST
		req.Header.Add("Accept", "application/json")
		req.Header.Add("Content-Type", "application/json")
		//HANDLING ERROR
		if err != nil {
			logger.Error("Cart Items do not exist")
			ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Cart Items do not exist"))
			return
		} else {
			data, err2 := ioutil.ReadAll(req.Body)
			if err2 != nil {
				logger.Error("Error during RESPONSE value from GET Request")
				ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Some shit doesn't work during GET - CONSUME"))
			}
			err3 := json.Unmarshal(data, &cartItemsRespBody)
			if err3 != nil {
				ctx.JSON(http.StatusBadRequest, errs.NewBadRequest("Cart Items couldn't be UNMARSHALLED."))
				return
			}

			fmt.Println("Cart Items Fetched ✅")
			logger.Info("Cart Items Fetched ✅", cartItemsRespBody)
			ctx.JSON(http.StatusAccepted, cartItemsRespBody)
		}
	}
}
