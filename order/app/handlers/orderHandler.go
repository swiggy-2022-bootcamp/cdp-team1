package handlers

import (
	"net/http"
	"orderService/domain/model"
	"orderService/domain/service"
	"orderService/internal/error"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) OrderHandler {
	return OrderHandler{orderService: orderService}
}

// Order godoc
// @Summary To place an order for a user
// @Description To create orders
// @Tags Order
// @Schemes
// @Accept json
// @Produce json
// @Success	200  string 	Order created successfully
// @Failure 400  string 	Bad request
// @Failure 500  string 	Internal server error
// @Router /order [POST]
func (oh OrderHandler) CreateOrder(c *gin.Context) {

	var order model.Order
	if err := c.BindJSON(&order); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&order); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	err := oh.orderService.CreateOrder(&order)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment mode added successfully"})

}

func (oh OrderHandler) GetAllOrder(c *gin.Context) {

	result, err := oh.orderService.GetAllOrders()
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (oh OrderHandler) GetOrderByStatus(c *gin.Context) {

	var status string
	if err := c.BindJSON(&status); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}
	result, err := oh.orderService.GetOrderByStatus(status)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func (oh OrderHandler) GetOrderById(c *gin.Context) {

	customer_id := "dummy" // To be replaced with grpc call to auth service

	var order_id string
	if err := c.BindJSON(&order_id); err != nil {
		c.Error(err)
		requestError := error.NewBadRequestError(err.Error())
		c.JSON(requestError.Code, gin.H{"message": requestError.Message})
		return
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&order_id); validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	result, err := oh.orderService.GetOrderById(order_id, customer_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
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
