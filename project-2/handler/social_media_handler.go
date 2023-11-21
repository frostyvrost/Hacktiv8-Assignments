package handler

import (
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SocialMediasHandler interface {
	AddSocialMedia(ctx *gin.Context)
	GetSocialMedias(ctx *gin.Context)
	UpdateSocialMedia(ctx *gin.Context)
	DeleteSocialMedia(ctx *gin.Context)
}

type socialMediaHndlerImpl struct {
	ss service.SocialMediaService
}

func NewSocialMediasHandler(socialMediaService service.SocialMediaService) SocialMediasHandler {
	return &socialMediaHndlerImpl{
		ss: socialMediaService,
	}
}

func (s *socialMediaHndlerImpl) AddSocialMedia(ctx *gin.Context) {

	socialMediaPayload := &dto.NewSocialMediaRequest{}
	user := ctx.MustGet("userData").(models.User)

	if err := ctx.ShouldBindJSON(socialMediaPayload); err != nil {
		errBindJson := pkg.NewUnprocessableModelsError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := s.ss.AddSocialMedia(user.Id, socialMediaPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (s *socialMediaHndlerImpl) DeleteSocialMedia(ctx *gin.Context) {
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	response, err := s.ss.DeleteSocialMedia(socialMediaId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (s *socialMediaHndlerImpl) GetSocialMedias(ctx *gin.Context) {
	response, err := s.ss.GetSocialMedias()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (s *socialMediaHndlerImpl) UpdateSocialMedia(ctx *gin.Context) {

	socialMediaPayload := &dto.UpdateSocialMediaRequest{}
	socialMediaId, _ := strconv.Atoi(ctx.Param("socialMediaId"))

	if err := ctx.ShouldBindJSON(socialMediaPayload); err != nil {
		errBindJson := pkg.NewUnprocessableModelsError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := s.ss.UpdateSocialMedia(socialMediaId, socialMediaPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
