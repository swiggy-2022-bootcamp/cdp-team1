package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	apperros "qwik.in/transaction/app-erros"
	"qwik.in/transaction/domain/models"
	"qwik.in/transaction/domain/services"
)

var validate = validator.New()

type TransactionHandler struct {
	transactionService services.TransactionService
}

func NewTransactionHandler(transactionService services.TransactionService) TransactionHandler {
	return TransactionHandler{transactionService: transactionService}
}

// AddTransactionPoints godoc
// @Summary To add transaction points for a user.
// @Description This request will add transaction points to a user account.
// @Tags Transaction
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.TransactionAmount true "Transaction details"
// @Param userId path string true "User id"
// @Success	200  string		Transaction points added successfully
// @Failure 400  string   	Bad request
// @Failure 500  string		Internal Server Error
// @Router /transaction/{userId} [POST]
func (th TransactionHandler) AddTransactionPoints(c *gin.Context) {
	userId := c.Param("userId")
	var transactionAmount models.TransactionAmount

	if err := c.BindJSON(&transactionAmount); err != nil {
		c.Error(err)
		err_ := apperros.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	//validate request body
	if validationErr := validate.Struct(&transactionAmount); validationErr != nil {
		c.Error(validationErr)
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	transactionAmount.UserId = userId
	err := th.transactionService.AddTransactionPoints(&transactionAmount)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction points added successfully"})
}

// GetTransactionPointsByUserID godoc
// @Summary To get transaction points of a user.
// @Description This request will get transaction points from a user account.
// @Tags Transaction
// @Schemes
// @Accept json
// @Produce json
// @Param userId path string true "User id"
// @Success	200  int		Transaction points
// @Failure 404  string   	User not found
// @Failure 500  string		Internal Server Error
// @Router /transaction/{userId} [GET]
func (th TransactionHandler) GetTransactionPointsByUserID(c *gin.Context) {
	userId := c.Param("userId")
	points, err := th.transactionService.GetTransactionPointsByUserId(userId)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"transaction_points": points})
}

// UseTransactionPoints godoc
// @Summary To use transaction points for an order.
// @Description This request will use transaction points present in user account for passed order ID.
// @Tags Transaction
// @Schemes
// @Accept json
// @Produce json
// @Param req body models.TransactionAmount true "Transaction details"
// @Param userId path string true "User id"
// @Success	200  {object}	models.TransactionAmount
// @Failure 400  string   	Bad request
// @Failure 417  string		Expectation failed
// @Failure 500  string		Internal Server Error
// @Router /transaction/use-transaction-points/{userId} [POST]
func (th TransactionHandler) UseTransactionPoints(c *gin.Context) {
	userId := c.Param("userId")
	var transactionAmount models.TransactionAmount

	if err := c.BindJSON(&transactionAmount); err != nil {
		c.Error(err)
		err_ := apperros.NewBadRequestError(err.Error())
		c.JSON(err_.Code, gin.H{"message": err_.Message})
		return
	}

	//validate request body
	if validationErr := validate.Struct(&transactionAmount); validationErr != nil {
		c.Error(validationErr)
		c.JSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	transactionAmount.UserId = userId
	_, newTransactionAmount, err := th.transactionService.UseTransactionPoints(&transactionAmount)
	if err != nil {
		c.Error(err.Error())
		c.JSON(err.Code, gin.H{"message": err.Message})
		return
	}
	c.JSON(http.StatusOK, newTransactionAmount)
}
