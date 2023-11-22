package auth

import (
	"net/http"
	"project-3/database"
	"project-3/model"
	"project-3/pkg"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		verifiedToken, err := pkg.VerifyToken(context)
		if err != nil {
			context.AbortWithStatusJSON(http.StatusUnauthorized, err.Error())
			return
		}

		context.Set("userData", verifiedToken)
		context.Next()
	}
}

func AdminAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		userData := context.MustGet("userData").(jwt.MapClaims)
		userRole := userData["role"].(string)

		if userRole != "admin" {
			err := pkg.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func CategoryAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		categoryId, err := pkg.GetIdParam(context, "categoryId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		db := database.GetDB()
		category := model.Category{}

		errMsg := db.First(&category, categoryId).Error
		if errMsg != nil {
			err := pkg.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func TaskAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		taskId, err := pkg.GetIdParam(context, "taskId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		userData := context.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))

		db := database.GetDB()
		task := model.Task{}

		errMsg := db.Select("user_id").First(&task, taskId).Error
		if errMsg != nil {
			err := pkg.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if task.UserID != userID {
			err := pkg.Unautorized("You are not allowed to access this data")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
