package dto

import (
	"net/http"
	"project-3/model"
	"project-3/pkg"
	"project-3/service"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func CreateTask(context *gin.Context) {
	var task model.Task

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.AbortWithStatusJSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	task.Status = false

	createdResponse, err := service.TaskService.CreateTask(&task, userID)

	if err != nil {
		context.AbortWithStatusJSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":          createdResponse.ID,
		"title":       createdResponse.Title,
		"status":      createdResponse.Status,
		"description": createdResponse.Description,
		"user_id":     createdResponse.UserID,
		"category_id": createdResponse.CategoryID,
		"created_at":  createdResponse.CreatedAt,
	})
}

func GetAllTasks(context *gin.Context) {

	results, err := service.TaskService.GetAllTasks()

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	tasks := make([]gin.H, 0, len(results))

	for _, result := range results {
		task := gin.H{
			"id":          result.ID,
			"title":       result.Title,
			"status":      result.Status,
			"description": result.Description,
			"user_id":     result.UserID,
			"category_id": result.CategoryID,
			"created_at":  result.CreatedAt,
			"User": gin.H{
				"id":        result.User.ID,
				"email":     result.User.Email,
				"full_name": result.User.FullName,
			},
		}

		tasks = append(tasks, task)
	}

	context.JSON(http.StatusOK, tasks)
}

func UpdateTask(context *gin.Context) {
	var task model.TaskUpdate

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	taskIDInt, _ := strconv.Atoi(context.Param("taskId"))
	taskID := uint(taskIDInt)

	result, err := service.TaskService.UpdateTask(&task, taskID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":          result.ID,
		"title":       result.Title,
		"description": result.Description,
		"status":      result.Status,
		"user_id":     result.UserID,
		"category_id": result.CategoryID,
		"updated_at":  result.CreatedAt,
	})
}

func UpdateStatusTask(context *gin.Context) {
	var task model.TaskStatusUpdate

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	taskIDInt, _ := strconv.Atoi(context.Param("taskId"))
	taskID := uint(taskIDInt)

	result, err := service.TaskService.UpdateStatusTask(&task, taskID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":          result.ID,
		"title":       result.Title,
		"description": result.Description,
		"status":      result.Status,
		"user_id":     result.UserID,
		"category_id": result.CategoryID,
		"updated_at":  result.CreatedAt,
	})
}

func UpdateCategoryIdTask(context *gin.Context) {
	var task model.TaskCategoryUpdate

	if err := context.ShouldBindJSON(&task); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")
		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	taskIDInt, _ := strconv.Atoi(context.Param("taskId"))
	taskID := uint(taskIDInt)

	result, err := service.TaskService.UpdateCategoryIdTask(&task, taskID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":          result.ID,
		"title":       result.Title,
		"description": result.Description,
		"status":      result.Status,
		"user_id":     result.UserID,
		"category_id": result.CategoryID,
		"updated_at":  result.CreatedAt,
	})
}

func DeleteTask(context *gin.Context) {
	id, _ := pkg.GetIdParam(context, "taskId")

	err := service.TaskService.DeleteTask(id)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Task has been successfully deleted",
	})
}
