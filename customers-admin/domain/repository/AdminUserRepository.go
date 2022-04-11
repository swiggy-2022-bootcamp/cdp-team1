package repository

import "qwik.in/customers-admin/domain/model"

type AdminUserInterface interface {
	GetAll() []model.AdminUser
}

type AdminUserRepository struct {
}

func (adminUser *AdminUserRepository) GetAll() []model.AdminUser {
	return []model.AdminUser{}
}
