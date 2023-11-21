package service

import (
	"project-2/repo"
	"project-2/pkg"
	"project-2/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	AuthorizationPhoto() gin.HandlerFunc
	AuthorizationComment() gin.HandlerFunc
	AuthorizationSocialMedia() gin.HandlerFunc
}

type authServiceImpl struct {
	ur repo.UserRepository
	pr repo.PhotoRepository
	cr repo.CommentRepository
	sr repo.SocialMediaRepository
}

func NewAuthService(userRepo repo.UserRepository, photoRepo repo.PhotoRepository, commentRepo repo.CommentRepository, socialMediaRepo repo.SocialMediaRepository) AuthService {
	return &authServiceImpl{
		ur: userRepo,
		pr: photoRepo,
		cr: commentRepo,
		sr: socialMediaRepo,
	}
}

// Authentication implements a.
func (a *authServiceImpl) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		invalidToken := pkg.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")

		user := models.User{}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		_, err = a.ur.FetchById(user.Id)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidToken.Status(), invalidToken)
			return
		}

		ctx.Set("userData", user)
		ctx.Next()
	}
}

// AuthorizationPhoto implements a.
func (a *authServiceImpl) AuthorizationPhoto() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := ctx.MustGet("userData").(models.User)

		photoId, _ := strconv.Atoi(ctx.Param("photoId"))

		photo, err := a.pr.GetPhotoId(photoId)

		if err != nil {
			if err.Status() == http.StatusNotFound {
				errBadRequest := pkg.NewBadRequestError("photo not found")
				ctx.AbortWithStatusJSON(errBadRequest.Status(), errBadRequest)
				return
			}
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if photo.UserId != user.Id {
			errUnathorized := pkg.NewUnathorizedError("you are not authorized to modify the photo")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}

// AuthorizationComment implements a.
func (a *authServiceImpl) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(models.User)

		commentId, _ := strconv.Atoi(ctx.Param("commentId"))

		comment, err := a.cr.GetCommentById(commentId)

		if err != nil {
			if err.Status() == http.StatusNotFound {
				errBadRequest := pkg.NewBadRequestError("comment not found")
				ctx.AbortWithStatusJSON(errBadRequest.Status(), errBadRequest)
				return
			}
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if comment.UserId != user.Id {
			errUnathorized := pkg.NewUnathorizedError("you are not authorized to modify the comment")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}

// AuthorizationSocialMedia implements AuthService.
func (a *authServiceImpl) AuthorizationSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(models.User)

		socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

		socialMedia, err := a.sr.GetSocialMediaById(socialMediaId)

		if err != nil {
			if err.Status() == http.StatusNotFound {
				errBadRequest := pkg.NewBadRequestError("social media not found")
				ctx.AbortWithStatusJSON(errBadRequest.Status(), errBadRequest)
				return
			}
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		if socialMedia.UserId != user.Id {
			errUnathorized := pkg.NewUnathorizedError("you are not authorized to modify the comment")
			ctx.AbortWithStatusJSON(errUnathorized.Status(), errUnathorized)
		}

		ctx.Next()
	}
}
