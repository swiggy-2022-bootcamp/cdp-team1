package handlers

import (
	"cartService/domain/model"
	"cartService/domain/service"
	"cartService/internal/error"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartService service.CartService) CartHandler {
	return CartHandler{cartService: cartService}
}

// CreateCart godoc
// @Summary To create a new cart
// @Description To create a new cart for the logged in user
// @Tags Cart
// @Schemes
// @Produce json
// @Success	200  string 	Cart created successfully
// @Failure 400  string 	Bad request
// @Failure 500  string 	Internal server error
// @Router /cart [POST]
func (ch CartHandler) CreateCart(c *gin.Context) {

	customer_id := c.Param("id")

	var product model.Product
	if err := c.BindJSON(&product); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&product); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	err := ch.cartService.AddToCart(&product, customer_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart created successfully"})
}

// UpdateCart godoc
// @Summary To update cart
// @Description Update cart's quantity
// @Tags Cart
// @Schemes
// @Accept json
// @Produce json
// @Param id string true "id"
// @Param req body string true "Updated quantity"
// @Success	200  string 	Cart updated successfully
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /cart [PUT]
func (ch CartHandler) UpdateCart(c *gin.Context) {

	customer_id := c.Param("id")

	// get cart id from parameter and status from body
	var product model.Product
	if err := c.BindJSON(&product); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&product); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	productId := product.ProductId
	quantity := product.Quantity

	err := ch.cartService.UpdateCart(customer_id, productId, quantity)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully"})
}

// GetAllCart godoc
// @Summary To get all carts.
// @Description Fetch all carts in the database
// @Tags Cart
// @Schemes
// @Produce json
// @Success	200  {object} 	[]model.Order
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /cart [GET]
func (ch CartHandler) GetAllCart(c *gin.Context) {

	result, err := ch.cartService.GetAllCart()
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// DeleteCartById godoc
// @Summary To delete a cart
// @Description Delete the cart by the customer id
// @Tags Cart
// @Schemes
// @Accept json
// @Param id string true "id"
// @Success	200  string 	Cart deleted successfully
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /cart/:id [DELETE]
func (ch CartHandler) DeleteCart(c *gin.Context) {

	var customer_id string
	if err := c.BindJSON(&customer_id); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&customer_id); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	err := ch.cartService.DeleteCartByCustomerId(customer_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}

// DeleteCartAll godoc
// @Summary To delete all the cart
// @Description Delete all the cart by the customer id
// @Tags Cart
// @Schemes
// @Accept json
// @Param id string true "id"
// @Success	200  string 	Cart deleted successfully
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /cart/empty [DELETE]
func (ch CartHandler) DeleteCartAll(c *gin.Context) {

	err := ch.cartService.DeleteAllCart()
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart deleted successfully"})
}

// set SECRET=sUpErCaLiFrAgIlIsTiCeXpIaLiDoCiOuS in .env file
// const secret string = os.Getenv("SECRET")

// func UserIDFromAuthToken(authToken string) (string, *errs.AppError) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		return "", errs.NewAuthenticationError("unexpected signing method")
// 	}
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

// 		return claims["user_id"].(string), nil
// 	}
// 	return "", errs.NewAuthenticationError("invalid token")
// }
