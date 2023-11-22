package dto

import (
	"net/http"
	"project-3/model"
	"project-3/pkg"
	"project-3/service"

	"github.com/gin-gonic/gin"
)

func CreateCategory(context *gin.Context) {
	var category model.Category

	if err := context.ShouldBindJSON(&category); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.CategoryService.CreateCategory(&category)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         result.ID,
		"type":       result.Type,
		"created_at": result.CreatedAt,
	})
}

func UpdateCategory(context *gin.Context) {
	var update model.CategoryUpdate

	if err := context.ShouldBindJSON(&update); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	categoryId, _ := pkg.GetIdParam(context, "categoryId")

	result, err := service.CategoryService.UpdateCategory(&update, categoryId)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         result.ID,
		"type":       result.Type,
		"updated_at": result.UpdatedAt,
	})
}

func GetAllCategories(context *gin.Context) {
	categories, err := service.CategoryService.GetAllCategories()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	var categoriesMaps []map[string]interface{}

	for _, category := range categories {
		categoryMap := createCategoryMap(category)
		categoriesMaps = append(categoriesMaps, categoryMap)
	}

	context.JSON(http.StatusOK, categoriesMaps)
}

func createCategoryMap(category *model.Category) map[string]interface{} {
	var taskMaps []map[string]interface{}

	for _, task := range category.Tasks {
		taskMap := createTaskMap(&task)
		taskMaps = append(taskMaps, taskMap)
	}

	return map[string]interface{}{
		"id":         category.ID,
		"type":       category.Type,
		"created_at": category.CreatedAt,
		"updated_at": category.UpdatedAt,
		"Tasks":      taskMaps,
	}
}

func createTaskMap(task *model.Task) map[string]interface{} {
	return map[string]interface{}{
		"id":          task.ID,
		"title":       task.Title,
		"description": task.Description,
		"user_id":     task.UserID,
		"category_id": task.CategoryID,
		"created_at":  task.CreatedAt,
		"updated_at":  task.UpdatedAt,
	}
}

func DeleteCategory(context *gin.Context) {
	id, _ := pkg.GetIdParam(context, "categoryId")

	err := service.CategoryService.DeleteCategory(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Category has been successfully deleted",
	})
}
