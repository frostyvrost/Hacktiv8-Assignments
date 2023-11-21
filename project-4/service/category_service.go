package service

import (
	"project-4/models"
	"project-4/pkg"
	"project-4/repo"

	"github.com/asaskevich/govalidator"
)

type categoryServiceRepo interface {
	CreateCategory(*models.Category) (*models.Category, pkg.Error)
	GetAllCategories() ([]*models.Category, pkg.Error)
	UpdateCategory(*models.CategoryUpdate, int) (*models.Category, pkg.Error)
	DeleteCategory(int) pkg.Error
}

type categoryService struct{}

var CategoryService categoryServiceRepo = &categoryService{}

func (c *categoryService) CreateCategory(category *models.Category) (*models.Category, pkg.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	categoryResponse, err := repo.CategoryRepo.CreateCategory(category)

	if err != nil {
		return nil, err
	}

	return categoryResponse, nil
}

func (s *categoryService) GetAllCategories() ([]*models.Category, pkg.Error) {
	categories, err := repo.CategoryRepo.GetAllCategories()

	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (c *categoryService) UpdateCategory(category *models.CategoryUpdate, categoryId int) (*models.Category, pkg.Error) {
	if _, err := govalidator.ValidateStruct(category); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	categoryResponse, err := repo.CategoryRepo.UpdateCategory(category, categoryId)

	if err != nil {
		return nil, err
	}

	return categoryResponse, nil
}

func (s *categoryService) DeleteCategory(categoryId int) pkg.Error {
	err := repo.CategoryRepo.DeleteCategory(categoryId)

	if err != nil {
		return err
	}

	return nil
}
