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

// CreaterOrder godoc
// @Summary To place an order for a user
// @Description To create orders
// @Tags Order
// @Schemes
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

	c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})

}

// GetAllOrders godoc
// @Summary To get all orders.
// @Description Fetch all orders in the database
// @Tags Order
// @Schemes
// @Produce json
// @Success	200  {object} 	model.Order
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /orders [GET]
func (oh OrderHandler) GetAllOrder(c *gin.Context) {

	result, err := oh.orderService.GetAllOrders()
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetOrderByOrderStatus godoc
// @Summary To get all order of a specific status.
// @Description Fetch all the orders with a specific status
// @Tags Order
// @Schemes
// @Accept json
// @Produce json
// @Param status path string true "status"
// @Success	200  {object} 	model.Order
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /orders/status/{status} [GET]
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

// GetOrderByOrderId godoc
// @Summary To get an order by order id.
// @Description Fetch the order by the order id
// @Tags Order
// @Schemes
// @Accept json
// @Produce json
// @Param id path string true "order id"
// @Success	200  {object} 	model.Order
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /orders/{id} [GET]
func (oh OrderHandler) GetOrderById(c *gin.Context) {

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

	result, err := oh.orderService.GetOrderById(order_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// GetOrderByOrderId godoc
// @Summary To get an order by order id.
// @Description Fetch the order by the order id
// @Tags Order
// @Schemes
// @Accept json
// @Produce json
// @Param id path string true "order id"
// @Success	200  string 	Order updated successfully
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /orders/{id} [PUT]
func (oh OrderHandler) UpdateOrder(c *gin.Context) {

	order_id := c.Param("id")

	// get order id from parameter and status from body
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

	err := oh.orderService.UpdateOrder(order_id, status)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// DeleteOrderById godoc
// @Summary To delete an order by order id.
// @Description Delete the order by the order id
// @Tags Order
// @Schemes
// @Accept json
// @Param id path string true "order id"
// @Success	200  string 	Order deleted successfully
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /orders/{id} [DELETE]
func (oh OrderHandler) DeleteOrderById(c *gin.Context) {

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

	err := oh.orderService.DeleteOrderById(order_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}

// GetOrderByCustomerId godoc
// @Summary To get an order by customer id.
// @Description Fetch all the order by the customer id
// @Tags Order
// @Schemes
// @Accept json
// @Produce json
// @Param id path string true "order id"
// @Success	200  {object} 	model.Order
// @Failure 500  string 	Internal server error
// @Failure 404  string 	Order not found
// @Router /orders/user/{id} [GET]
func (oh OrderHandler) GetOrderByCustomerId(c *gin.Context) {

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

	result, err := oh.orderService.GetOrderByCustomerId(customer_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// CreateInvoice godoc
// @Summary To create an invoice
// @Description To create an invoice for a given order id
// @Tags Order
// @Schemes
// @Param id path string true "order id"
// @Success	200  string 	Invoice Number created successfully
// @Failure 400  string 	Bad request
// @Failure 500  string 	Internal server error
// @Router /orders/invoice/{id} [POST]
func (oh OrderHandler) CreateInvoice(c *gin.Context) {

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

	err := oh.orderService.CreateInvoice(order_id)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Invoice created successfully"})
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
