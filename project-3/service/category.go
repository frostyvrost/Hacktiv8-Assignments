package service

import (
	"project-3/model"
	"project-3/pkg"
	"project-3/repo"

	"github.com/asaskevich/govalidator"
)

type categoryServiceRepo interface {
	CreateCategory(*model.Category) (*model.Category, pkg.Error)
	UpdateCategory(*model.CategoryUpdate, uint) (*model.Category, pkg.Error)
	GetAllCategories() ([]*model.Category, pkg.Error)
	DeleteCategory(uint) pkg.Error
}

type categoryService struct{}

var CategoryService categoryServiceRepo = &categoryService{}

func (t *categoryService) CreateCategory(category *model.Category) (*model.Category, pkg.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.CategoryModel.CreateCategory(category)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *categoryService) UpdateCategory(category *model.CategoryUpdate, categoryId uint) (*model.Category, pkg.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	result, err := repo.CategoryModel.UpdateCategory(category, categoryId)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *categoryService) GetAllCategories() ([]*model.Category, pkg.Error) {
	categories, err := repo.CategoryModel.GetAllCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryService) DeleteCategory(categoryId uint) pkg.Error {
	err := repo.CategoryModel.DeleteCategory(categoryId)

	if err != nil {
		return err
	}

	return nil
}
