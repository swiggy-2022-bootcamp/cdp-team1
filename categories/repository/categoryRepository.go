package repository

import (
	"qwik.in/categories/entity"
)

type CategoryRepository interface {
	Connect() error
	FindAll() ([]entity.Category, error)
	FindOne(category_id string) (entity.Category, error)
	SaveCategory(category entity.Category) error
	DeleteCategory(category_id string) error
	// FindAndUpdate(rewardId string, reward entity.Reward) error
}
