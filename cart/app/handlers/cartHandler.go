package handlers

import (
	"cartService/domain/model"
	"cartService/domain/service"
	apperrors "cartService/internal/error"
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

	// get userid from jwt
	// userID := c.MustGet("userID").(string)

	var cartModel model.Cart
	if err := c.BindJSON(&cartModel); err != nil {
		c.Error(err)
		err2 := apperrors.NewBadRequestError(err.Error())
		c.JSON(err2.Code, gin.H{"message": err2.Message})
		return
	}

	if validationErr := validate.Struct(&cartModel); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	err := ch.cartService.AddToCart(cartModel)

	if err != nil {
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Added to cart successfully"})
}

func (ch CartHandler) UpdateCart(c *gin.Context) {
	// todo check
}

func (ch CartHandler) GetCart(c *gin.Context) {
	// todo
}

func (ch CartHandler) DeleteCart(c *gin.Context) {
	// todo
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
