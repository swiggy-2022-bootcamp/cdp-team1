package handlers

import (
	"github.com/gin-gonic/gin"
	"qwik.in/customers-admin/domain/service"
)

var adminService service.AdminService

func GetAdminUsers(c *gin.Context) {
	c.JSON(200, adminService.GetAllAdminUsers())
}
