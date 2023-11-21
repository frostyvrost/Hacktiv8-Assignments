package main

import (
	"project-3/config"
	"project-3/db"
	"project-3/handler"
	"project-3/repo/pg"
	"project-3/service"

	"github.com/gin-gonic/gin"

	swaggoFile "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Kanban Board
// @version 1.0
// @description Final Project 3 Kampus Merdeka

// @contact.name GLNG-KS07 - Vormes Gema Merdeka

func StartApp() {

	config.LoadEnv()

	db.InitiliazeDatabase()
	db := db.GetDatabaseInstance()

	//Dependency Injection
	userRepo := pg.NewUserPG(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	taskRepo := pg.NewTaskRepo(db)
	categoryRepo := pg.NewCategoryRepo(db)

	taskService := service.NewTaskService(taskRepo, categoryRepo, userRepo)
	categoryService := service.NewCategorySevice(categoryRepo, taskRepo)

	categoryHandler := handler.NewCategoryHandler(categoryService)
	taskHandler := handler.NewTaskHandler(taskService)

	authService := service.NewAuthService(userRepo, taskRepo, categoryRepo)

	route := gin.Default()

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggoFile.Handler))

	userRoute := route.Group("/users")
	{
		userRoute.POST("/register", userHandler.Register)
		userRoute.POST("/login", userHandler.Login)
		userRoute.PUT("/update-account", authService.Authentication(), userHandler.Update)
		userRoute.DELETE("/delete-account", authService.Authentication(), userHandler.Delete)
		userRoute.POST("/admin", userHandler.Admin)
	}

	userRoute = route.Group("/categories")
	{
		userRoute.POST("", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.Create)
		userRoute.GET("", authService.Authentication(), categoryHandler.Get)
		userRoute.PATCH("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.Update)
		userRoute.DELETE("/:categoryId", authService.Authentication(), authService.AdminAuthorization(), categoryHandler.Delete)
	}

	userRoute = route.Group("/tasks")
	{
		userRoute.POST("", authService.Authentication(), taskHandler.Create)
		userRoute.GET("", authService.Authentication(), taskHandler.Get)
		userRoute.PUT("/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.Update)
		userRoute.PATCH("/update-status/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateByStatus)
		userRoute.PATCH("/update-category/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.UpdateByCategoryId)
		userRoute.DELETE("/:taskId", authService.Authentication(), authService.TaskAuthorization(), taskHandler.Delete)
	}

	route.Run(":" + config.AppConfig().Port)
}
