package handlers

import (
	"orderService/domain/service"

	"github.com/gin-gonic/gin"
)

// var validate = validator.New()

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) OrderHandler {
	return OrderHandler{orderService: orderService}
}

func (oh OrderHandler) GetAllOrder(c *gin.Context) {
	//todo
}

func (oh OrderHandler) GetOrderByStatus(c *gin.Context) {
	//todo
}

func (oh OrderHandler) GetOrderById(c *gin.Context) {
	//todo
}

func (oh OrderHandler) UpdateOrder(c *gin.Context) {
	//todo
}

func (oh OrderHandler) DeleteOrderById(c *gin.Context) {
	//todo
}

func (oh OrderHandler) GetOrderByCustomerId(c *gin.Context) {
	//todo
}

func (oh OrderHandler) CreateInvoice(c *gin.Context) {
	//todo
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
