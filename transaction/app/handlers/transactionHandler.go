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
