package service

import "qwik.in/categories/entity"

type CategoryService interface {
	GetAll() ([]entity.Category, error)
	SearchCategory(categoryid string) (entity.Category, error)
	UpdateCategory(categoryid string, category entity.Category) error
	
}
