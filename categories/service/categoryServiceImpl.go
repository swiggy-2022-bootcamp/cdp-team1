package service

import (
	"qwik.in/categories/entity"
	"qwik.in/categories/repository"
)

type CategoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{categoryRepository: categoryRepository}
}

func (c CategoryServiceImpl) GetAll() ([]entity.Category, error) {
	all, err := c.categoryRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return all, nil
}
func (c CategoryServiceImpl) SearchCategory(category_id string) (entity.Category, error) {
	all, err := c.categoryRepository.FindOne(category_id)
	r := entity.Category{}
	if err != nil {
		return r, err
	}
	return all, nil
}
func (c CategoryServiceImpl) UpdateCategory(categoryId string, category entity.Category) error {
	category.Category_id = categoryId
	if err := c.categoryRepository.SaveCategory(category); err != nil {
		return err
	}
	return nil
}

func (c CategoryServiceImpl) DeleteCategory(categoryid string) error {
	err := c.categoryRepository.DeleteCategory(categoryid)
	if err != nil {
		return err
	}
	return nil
}
func (c CategoryServiceImpl) CreateCategory(category entity.Category) error {
	if err := c.categoryRepository.SaveCategory(category); err != nil {
		return err
	}
	return nil
}
