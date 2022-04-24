package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qwik.in/account-frontstore/domain/model"
	"qwik.in/account-frontstore/domain/repository"
	"qwik.in/account-frontstore/domain/service"
	"qwik.in/account-frontstore/internal/errors"
)

var accountService service.AccountServiceInterface

func init() {
	accountService = service.InitAccountService(&repository.AccountRepository{}, &repository.GrpcClientRepository{})
}

// RegisterAccount godoc
// @Summary Register Customer
// @Description Lets register customer in front store.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param req body model.Account true "New account to add"
// @Success	200  {object} model.Account
// @Router /customer [POST]
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

// GetAccountById godoc
// @Summary Get account by ID
// @Description Lets front store get customer account by customer id.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param accessorId path string true "accessor id"
// @Success	200  {object} model.Account
// @Router /account/{accessorId} [GET]
func GetAccountById(c *gin.Context) {
	fetchedAccount, err := accountService.GetAccountById(c.Param("accessorId"))

	if err != nil {
		accountErr, _ := err.(*errors.AccountError)
		c.JSON(accountErr.Status, accountErr.ErrorMessage)
		return
	}

	c.JSON(200, *fetchedAccount)
}

// UpdateAccount godoc
// @Summary Update Account
// @Description Lets front store update account by id.
// @Tags
// @Schemes
// @Accept json
// @Produce json
// @Param accessorId path string true "accessorId id"
// @Param req body model.Account true "Updated account"
// @Success	200  {object} model.Account
// @Router /account/{accessorId} [PUT]
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
