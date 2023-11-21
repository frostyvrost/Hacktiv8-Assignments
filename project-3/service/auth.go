package service

import (
	"project-3/models"
	"project-3/pkg"
	"project-3/repo"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AdminAuthorization() gin.HandlerFunc
	TaskAuthorization() gin.HandlerFunc
	CategoryAuthorization() gin.HandlerFunc
}

type authService struct {
	userRepo     repo.UserRepository
	taskRepo     repo.TaskRepository
	categoryRepo repo.CategoryRepository
}

// , taskRepo task_repo.Repository
func NewAuthService(userRepo repo.UserRepository, taskRepo repo.TaskRepository, categoryRepo repo.CategoryRepository) AuthService {
	return &authService{
		userRepo:     userRepo,
		taskRepo:     taskRepo,
		categoryRepo: categoryRepo,
	}
}

func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = pkg.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		var user models.User

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.userRepo.GetUserById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}

func (a *authService) AdminAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(models.User)
		if !ok {
			newError := pkg.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}
		if userData.Role != "admin" {
			newError := pkg.NewUnauthorizedError("You're not authorized to access this endpoint")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) TaskAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(models.User)
		if !ok {
			newError := pkg.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		taskId, _ := strconv.Atoi(ctx.Param("taskId"))

		task, err := a.taskRepo.GetTaskById(taskId)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if task.UserId != userData.Id {
			newError := pkg.NewUnauthorizedError("You're not authorized to modify this task")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}

func (a *authService) CategoryAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData, ok := ctx.MustGet("userData").(models.User)
		if !ok {
			newError := pkg.NewBadRequest("Failed to get user data")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		categoryId, _ := strconv.Atoi(ctx.Param("categoryId"))

		task, err := a.taskRepo.GetTaskById(categoryId)
		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if task.UserId != userData.Id {
			newError := pkg.NewUnauthorizedError("You're not authorized to modify this task")
			ctx.AbortWithStatusJSON(newError.Status(), newError)
			return
		}

		ctx.Next()
	}
}
