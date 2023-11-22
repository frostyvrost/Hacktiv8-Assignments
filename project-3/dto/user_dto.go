package dto

import (
	"net/http"
	"project-3/model"
	"project-3/pkg"
	"project-3/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Register(context *gin.Context) {
	var user model.User

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.UserService.Register(&user)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"id":         result.ID,
		"full_name":  result.FullName,
		"email":      result.Email,
		"created_at": result.CreatedAt,
	})
}

func Login(context *gin.Context) {
	var user model.LoginCredential

	if err := context.ShouldBindJSON(&user); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	result, err := service.UserService.Login(&user)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"token": result})
}

func UpdateUser(context *gin.Context) {
	var update model.UserUpdate

	if err := context.ShouldBindJSON(&update); err != nil {
		errorHandler := pkg.UnprocessibleEntity("Invalid JSON body")

		context.JSON(errorHandler.Status(), errorHandler)
		return
	}

	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	result, err := service.UserService.UpdateUser(userID, &update)
	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"id":         result.ID,
		"email":      result.Email,
		"full_name":  result.FullName,
		"updated_at": result.UpdatedAt,
	})
}

func DeleteUser(context *gin.Context) {
	userData := context.MustGet("userData").(jwt.MapClaims)
	userID := uint(userData["id"].(float64))

	_, err := service.UserService.DeleteUser(userID)

	if err != nil {
		context.JSON(err.Status(), err)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Your Account has been successfully deleted",
	})
}
