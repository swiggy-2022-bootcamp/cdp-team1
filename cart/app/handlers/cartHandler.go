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

func (ch CartHandler) CreateCart(c *gin.Context) {

	var cart model.Cart
	if err := c.BindJSON(&cart); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&cart); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	err := ch.cartService.AddToCart(&cart)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment mode added successfully"})
}

func (ch CartHandler) UpdateCart(c *gin.Context) {

	cart_id := c.Param("id")

	// get cart id from parameter and status from body
	var status string
	if err := c.BindJSON(&status); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&status); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	err := ch.cartService.UpdateCart(cart_id, status)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Cart updated successfully"})
}

func (ch CartHandler) GetAllCart(c *gin.Context) {

	result, err := ch.cartService.GetAllCart()
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

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

func (ch CartHandler) DeleteCartAll(c *gin.Context) {
	// todo
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
