package handler

import (
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PhotoHandler interface {
	AddPhoto(ctx *gin.Context)
	GetPhotos(ctx *gin.Context)
	UpdatePhoto(ctx *gin.Context)
	DeletePhoto(ctx *gin.Context)
}

type photoHandlerImpl struct {
	ps service.PhotoService
}

func NewPhotoHandler(photoService service.PhotoService) PhotoHandler {
	return &photoHandlerImpl{
		ps: photoService,
	}
}

func (p *photoHandlerImpl) AddPhoto(ctx *gin.Context) {
	user := ctx.MustGet("userData").(models.User)
	photoPayload := &dto.NewPhotoRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := pkg.NewUnprocessableModelsError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := p.ps.AddPhoto(user.Id, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (p *photoHandlerImpl) DeletePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	response, err := p.ps.DeletePhoto(photoId)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (p *photoHandlerImpl) GetPhotos(ctx *gin.Context) {
	response, err := p.ps.GetPhotos()

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}

func (p *photoHandlerImpl) UpdatePhoto(ctx *gin.Context) {
	photoId, _ := strconv.Atoi(ctx.Param("photoId"))

	photoPayload := &dto.PhotoUpdateRequest{}

	if err := ctx.ShouldBindJSON(photoPayload); err != nil {
		errBindJson := pkg.NewUnprocessableModelsError("invalid json body request")
		ctx.AbortWithStatusJSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := p.ps.UpdatePhoto(photoId, photoPayload)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(response.StatusCode, response)
}
