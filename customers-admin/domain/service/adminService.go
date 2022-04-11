package service

import (
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
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
