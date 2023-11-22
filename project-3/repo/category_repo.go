package repo

import (
	"project-3/database"
	"project-3/model"
	"project-3/pkg"
)

type categoryModelRepo interface {
	CreateCategory(*model.Category) (*model.Category, pkg.Error)
	UpdateCategory(*model.CategoryUpdate, uint) (*model.Category, pkg.Error)
	GetAllCategories() ([]*model.Category, pkg.Error)
	DeleteCategory(uint) pkg.Error
	GetCategoryById(categoryId uint) (*model.Category, pkg.Error)
}

type categoryModel struct{}

var CategoryModel categoryModelRepo = &categoryModel{}

func (t *categoryModel) CreateCategory(category *model.Category) (*model.Category, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&category).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return category, nil
}

func (t *categoryModel) UpdateCategory(update *model.CategoryUpdate, categoryId uint) (*model.Category, pkg.Error) {
	db := database.GetDB()

	var category model.Category
	err := db.First(&category, categoryId).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&category).Updates(update)

	return &category, nil
}

func (c *categoryModel) GetAllCategories() ([]*model.Category, pkg.Error) {
	db := database.GetDB()
	var categories []*model.Category

	err := db.Preload("Tasks").Find(&categories).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return categories, nil
}

func (c *categoryModel) DeleteCategory(categoryId uint) pkg.Error {
	db := database.GetDB()
	var category model.Category

	err := db.First(&category, categoryId).Error

	if err != nil {
		return pkg.ParseError(err)
	}

	db.Delete(&category)

	return nil
}

func (u *categoryModel) GetCategoryById(categoryId uint) (*model.Category, pkg.Error) {
	db := database.GetDB()
	var category model.Category

	err := db.First(&category, categoryId).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return &category, nil
}
