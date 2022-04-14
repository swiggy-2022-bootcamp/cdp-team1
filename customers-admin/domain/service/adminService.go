package service

import (
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
	"time"
)

type AdminServiceInterface interface {
	GetAllAdminUsers() []model.AdminUser
}

type AdminService struct {
	adminRepository repository.AdminUserRepository
}

func (adminService *AdminService) GetAllAdminUsers() []model.AdminUser {
	return adminService.adminRepository.GetAll()
}

func (adminService *AdminService) AddAdminUser(adminUser model.AdminUser) (*model.AdminUser, error) {
	adminUser.DateAdded = time.Now()
	return adminService.adminRepository.Create(adminUser)
}
