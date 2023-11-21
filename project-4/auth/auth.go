package auth

import (
	"project-4/database"
	"project-4/models"
	"project-4/pkg"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		verifiedToken, err := pkg.VerifyToken(context)

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
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
			err := pkg.Unauthorized("You are not allowed to access this data")
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
		category := models.Category{}

		errMsg := db.First(&category, categoryId).Error
		if errMsg != nil {
			err := pkg.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}

func ProductAuthorization() gin.HandlerFunc {
	return func(context *gin.Context) {
		productId, err := pkg.GetIdParam(context, "productId")

		if err != nil {
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		db := database.GetDB()
		product := models.Product{}

		errMsg := db.First(&product, productId).Error
		if errMsg != nil {
			err := pkg.NotFound("Data not found")
			context.AbortWithStatusJSON(err.Status(), err)
			return
		}

		context.Next()
	}
}
