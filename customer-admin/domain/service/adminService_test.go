package service

import (
	"github.com/stretchr/testify/assert"
	"qwik.in/customers-admin/domain/repository"
	"testing"
)

func TestShouldGetAtleastOneAdminRecord(t *testing.T) {
	customerRepo := &repository.AdminUserRepository{}
	adminService := InitAdminService(customerRepo)

	admins := adminService.GetAllAdminUsers()

	assert.Greater(t, len(admins), 0)
}
