package service

import (
	"golang.org/x/crypto/bcrypt"
	"qwik.in/customers-admin/domain/model"
	"qwik.in/customers-admin/domain/repository"
	"time"
)

type AdminServiceInterface interface {
	GetAllAdminUsers() []model.AdminUser
	AddAdminUser(adminUser model.AdminUser) (*model.AdminUser, error)
}

type AdminService struct {
	adminRepository repository.AdminUserRepositoryInterface
}

func InitAdminService(repositoryToInject repository.AdminUserRepositoryInterface) AdminServiceInterface {
	adminService := new(AdminService)
	adminService.adminRepository = repositoryToInject
	return adminService
}

func (adminService *AdminService) GetAllAdminUsers() []model.AdminUser {
	return adminService.adminRepository.GetAll()
}

func (adminService *AdminService) AddAdminUser(adminUser model.AdminUser) (*model.AdminUser, error) {
	//encrypt password
	adminPassword, err := bcrypt.GenerateFromPassword([]byte(adminUser.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	adminUser.Password = string(adminPassword)

	adminUser.DateAdded = time.Now()
	return adminService.adminRepository.Create(adminUser)
}
