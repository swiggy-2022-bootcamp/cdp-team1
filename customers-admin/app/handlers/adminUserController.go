package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/service"
	"qwik.in/customers-admin/internal/errors"
)

var adminService service.AdminService

func GetAdminUsers(c *gin.Context) {
	c.JSON(200, adminService.GetAllAdminUsers())
}

func AddAdminUser(c *gin.Context) {
	adminUser := model.AdminUser{}
	json.NewDecoder(c.Request.Body).Decode(&adminUser)
	addedUser, err := adminService.AddAdminUser(adminUser)

	if err != nil {
		userErr, _ := err.(*errors.CustomerError)
		c.JSON(userErr.Status, userErr.ErrorMessage)
		return
	}

	c.JSON(200, *addedUser)
}
