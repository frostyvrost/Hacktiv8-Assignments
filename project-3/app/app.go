package app

import (
	"project-3/auth"
	"project-3/dto"

	"github.com/gin-gonic/gin"
)

var PORT = ":8089"

func StartServer() {
	router := gin.Default()

	userRouter := router.Group("/users")
	{
		userRouter.POST("/register", dto.Register)
		userRouter.POST("/login", dto.Login)
		userRouter.PUT("/update-account", dto.UpdateUser)
		userRouter.DELETE("/delete-account", dto.DeleteUser)
	}

	categoryRouter := router.Group("/categories")
	{
		categoryRouter.Use(auth.Authentication())
		categoryRouter.POST("/", auth.AdminAuthorization(), dto.CreateCategory)
		categoryRouter.GET("/", dto.GetAllCategories)
		categoryRouter.PATCH("/:categoryId", auth.AdminAuthorization(), auth.CategoryAuthorization(), dto.UpdateCategory)
		categoryRouter.DELETE("/:categoryId", auth.AdminAuthorization(), auth.CategoryAuthorization(), dto.DeleteCategory)
	}

	taskRouter := router.Group("/tasks")
	{
		taskRouter.Use(auth.Authentication())
		taskRouter.POST("/", dto.CreateTask)
		taskRouter.GET("/", dto.GetAllTasks)
		taskRouter.PUT("/:taskId", auth.TaskAuthorization(), dto.UpdateTask)
		taskRouter.PATCH("/update-status/:taskId", auth.TaskAuthorization(), dto.UpdateStatusTask)
		taskRouter.PATCH("/update-category/:taskId", auth.TaskAuthorization(), dto.UpdateCategoryIdTask)
		taskRouter.DELETE("/:taskId", auth.TaskAuthorization(), dto.DeleteTask)
	}

	router.Run(PORT)
}
