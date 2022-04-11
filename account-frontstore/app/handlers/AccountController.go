package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/service"
	"qwik.in/account-frontstore/internal/errors"
)

var accountService service.AccountServiceInterface

func RegisterAccount(c *gin.Context) {
	newAccount := model.Account{}
	json.NewDecoder(c.Request.Body).Decode(&newAccount)
	createdAccount, err := accountService.CreateAccount(newAccount)

	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	c.JSON(200, *createdAccount)
}

func GetAccountById(c *gin.Context) {
	fetchedAccount, err := accountService.GetAccountById(c.Param("accessorId"))

	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedAccount)
}

func UpdateAccount(c *gin.Context) {
	account := model.Account{}
	json.NewDecoder(c.Request.Body).Decode(&account)
	updatedAccount, err := accountService.UpdateAccount(c.Param("accessorId"), account)

	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	c.JSON(200, *updatedAccount)
}
