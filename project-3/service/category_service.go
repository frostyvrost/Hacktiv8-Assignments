package service

import (
	"net/http"
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
	"project-3/repo"
)

type CategoryService interface {
	Create(categoryPayLoad *dto.NewCategoryRequest) (*dto.NewCategoryResponse, pkg.MessageErr)
	Get() (*dto.GetResponse, pkg.MessageErr)
	Update(categoryId int, categoryPayLoad *dto.UpdateRequest) (*dto.UpdateCategoryResponse, pkg.MessageErr)
	Delete(categoryId int) (*dto.DeleteCategoryByIdResponse, pkg.MessageErr)
}

type categoryService struct {
	categoryRepo repo.CategoryRepository
	taskRepo     repo.TaskRepository
}

// factory function
func NewCategorySevice(categoryRepo repo.CategoryRepository, taskRepo repo.TaskRepository) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
		taskRepo:     taskRepo,
	}
}

// Implements service interface
func (cs *categoryService) Create(categoryPayLoad *dto.NewCategoryRequest) (*dto.NewCategoryResponse, pkg.MessageErr) {
	err := pkg.ValidateStruct(categoryPayLoad)

	if err != nil {
		return nil, err
	}

	category := &models.Category{
		Type: categoryPayLoad.Type,
	}

	response, err := cs.categoryRepo.Create(category)

	if err != nil {
		return nil, err
	}

	response = &dto.NewCategoryResponse{
		Id:        response.Id,
		Type:      response.Type,
		CreatedAt: response.CreatedAt,
	}

	return response, nil
}

func (cs *categoryService) Get() (*dto.GetResponse, pkg.MessageErr) {
	categories, err := cs.categoryRepo.GetCategory()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	response := dto.GetResponse{
		StatusCode: http.StatusOK,
		Message:    "categories successfully fetched",
		Data:       categories,
	}

	return &response, nil
}

func (cs *categoryService) Update(categoryId int, categoryPayLoad *dto.UpdateRequest) (*dto.UpdateCategoryResponse, pkg.MessageErr) {
	err := pkg.ValidateStruct(categoryPayLoad)

	if err != nil {
		return nil, err
	}

	updateCategory, err := cs.categoryRepo.CheckCategoryId(categoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if updateCategory.Id != categoryId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	category := &models.Category{
		Id:   categoryId,
		Type: categoryPayLoad.Type,
	}

	response, err := cs.categoryRepo.UpdateCategory(category)

	if err != nil {
		return nil, err
	}

	return &dto.UpdateCategoryResponse{
		StatusCode: http.StatusOK,
		Message:    "Category has been succesfully updated",
		Data:       response,
	}, nil
}

func (cs *categoryService) Delete(categoryId int) (*dto.DeleteCategoryByIdResponse, pkg.MessageErr) {
	category, err := cs.categoryRepo.CheckCategoryId(categoryId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequest("invalid user")
		}
		return nil, err
	}

	if category.Id != categoryId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	cs.categoryRepo.DeleteCategory(categoryId)

	response := &dto.DeleteCategoryByIdResponse{
		Message: "Category has been successfully deleted",
	}

	return response, nil
}
