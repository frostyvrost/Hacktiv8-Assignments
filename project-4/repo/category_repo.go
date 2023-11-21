package repo

import (
	"fmt"
	"project-4/database"
	"project-4/models"
	"project-4/pkg"
)

type categoryDomainRepo interface {
	CreateCategory(*models.Category) (*models.Category, pkg.Error)
	GetAllCategories() ([]*models.Category, pkg.Error)
	GetCategoryById(int) (*models.Category, pkg.Error)
	UpdateCategory(*models.CategoryUpdate, int) (*models.Category, pkg.Error)
	DeleteCategory(int) pkg.Error
}

type categoryRepo struct{}

var CategoryRepo categoryDomainRepo = &categoryRepo{}

func (c *categoryRepo) CreateCategory(category *models.Category) (*models.Category, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&category).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return category, nil
}

func (s *categoryRepo) GetAllCategories() ([]*models.Category, pkg.Error) {
	db := database.GetDB()
	var categories []*models.Category

	err := db.Preload("Products").Find(&categories).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return categories, nil
}

func (c *categoryRepo) GetCategoryById(categoryId int) (*models.Category, pkg.Error) {
	db := database.GetDB()

	var category models.Category

	err := db.First(&category, categoryId).Error

	if err != nil {
		return nil, pkg.NotFound(fmt.Sprintf("Category with id %d not found", categoryId))
	}

	return &category, nil
}

func (c *categoryRepo) UpdateCategory(categoryUpdated *models.CategoryUpdate, categoryId int) (*models.Category, pkg.Error) {
	db := database.GetDB()

	var category models.Category

	err := db.First(&category, categoryId).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&category).Updates(categoryUpdated)

	return &category, nil
}

func (s *categoryRepo) DeleteCategory(categoryId int) pkg.Error {
	db := database.GetDB()
	var category models.Category

	err := db.Where("id = ?", categoryId).Delete(&category).Error

	if err != nil {
		return pkg.ParseError(err)
	}

	return nil
}
