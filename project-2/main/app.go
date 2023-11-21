package main

import (
	"net/http"
	"project-2/database"
	"project-2/handler"
	"project-2/repo/pg"
	"project-2/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func StartApplication() {

	database.LoadEnv()

	database.InitializeDatabase()
	db := database.GetInstanceDatabaseConnection()

	userRepo := pg.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	photoRepo := pg.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepo)
	photoHandler := handler.NewPhotoHandler(photoService)

	commentRepo := pg.NewCommentRepository(db)
	commentService := service.NewCommentService(commentRepo, photoRepo)
	commentHandler := handler.NewCommentHandler(commentService)

	socialMediaRepo := pg.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepo)
	socialMediaHandler := handler.NewSocialMediasHandler(socialMediaService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialMediaRepo)

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{
			http.MethodPost,
			http.MethodGet,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
		},
	}))

	users := app.Group("users")

	{
		users.POST("/register", userHandler.Register)
		users.POST("/login", userHandler.Login)
		users.PUT("", authService.Authentication(), userHandler.Update)
		users.DELETE("", authService.Authentication(), userHandler.Delete)
	}

	photos := app.Group("photos")

	{
		photos.POST("", authService.Authentication(), photoHandler.AddPhoto)
		photos.GET("", authService.Authentication(), photoHandler.GetPhotos)
		photos.PUT("/:photoId", authService.Authentication(), authService.AuthorizationPhoto(), photoHandler.UpdatePhoto)
		photos.DELETE("/:photoId", authService.Authentication(), authService.AuthorizationPhoto(), photoHandler.DeletePhoto)
	}

	comments := app.Group("comments")

	{
		comments.POST("", authService.Authentication(), commentHandler.AddComment)
		comments.GET("", authService.Authentication(), commentHandler.GetComments)
		comments.PUT("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.UpdateComment)
		comments.DELETE("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.DeleteComment)
	}

	socialMedias := app.Group("socialmedias")

	{
		socialMedias.POST("", authService.Authentication(), socialMediaHandler.AddSocialMedia)
		socialMedias.GET("", authService.Authentication(), socialMediaHandler.GetSocialMedias)
		socialMedias.PUT("/:socialMediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialMediaHandler.UpdateSocialMedia)
		socialMedias.DELETE("/:socialMediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialMediaHandler.DeleteSocialMedia)
	}

	app.Run(":" + database.AppConfig().Port)
}
