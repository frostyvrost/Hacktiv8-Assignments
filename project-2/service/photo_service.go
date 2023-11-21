package service

import (
	"net/http"
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/repo"
)

type photoServiceMock struct {
}

type PhotoService interface {
	AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, pkg.Error)
	GetPhotos() (*dto.GetPhotoResponse, pkg.Error)
	UpdatePhoto(photoId int, photoPayload *dto.PhotoUpdateRequest) (*dto.GetPhotoResponse, pkg.Error)
	DeletePhoto(photoId int) (*dto.GetPhotoResponse, pkg.Error)
}

type photoServiceImpl struct {
	pr repo.PhotoRepository
}

func NewPhotoService(photoRepository repo.PhotoRepository) PhotoService {
	return &photoServiceImpl{
		pr: photoRepository,
	}
}

// AddPhoto implements PhotoService.
func (p *photoServiceImpl) AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, pkg.Error) {

	err := pkg.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	photo := &models.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
		UserId:   userId,
	}

	response, err := p.pr.AddPhoto(photo)

	if err != nil {
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusCreated,
		Message:    "new photo successfully added",
		Data:       response,
	}, nil
}

// GetPhotos implements PhotoService.
func (p *photoServiceImpl) GetPhotos() (*dto.GetPhotoResponse, pkg.Error) {

	result, err := p.pr.GetPhotos()

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, err
		}
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "photos successfully fetched",
		Data:       result,
	}, nil
}

// UpdatePhoto implements PhotoService.
func (p *photoServiceImpl) UpdatePhoto(photoId int, photoPayload *dto.PhotoUpdateRequest) (*dto.GetPhotoResponse, pkg.Error) {

	err := pkg.ValidateStruct(photoPayload)

	if err != nil {
		return nil, err
	}

	photo := &models.Photo{
		Title:    photoPayload.Title,
		Caption:  photoPayload.Caption,
		PhotoUrl: photoPayload.PhotoUrl,
	}

	response, err := p.pr.UpdatePhoto(photoId, photo)

	if err != nil {
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "photo has been successfully updated",
		Data:       response,
	}, nil
}

// DeletePhoto implements PhotoService.
func (p *photoServiceImpl) DeletePhoto(photoId int) (*dto.GetPhotoResponse, pkg.Error) {

	err := p.pr.DeletePhoto(photoId)

	if err != nil {
		return nil, err
	}

	return &dto.GetPhotoResponse{
		StatusCode: http.StatusOK,
		Message:    "Your photo has been successfully deleted",
		Data:       nil,
	}, nil
}

var (
	AddPhoto    func(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, pkg.Error)
	GetPhotos   func() (*dto.GetPhotoResponse, pkg.Error)
	UpdatePhoto func(photoId int, photoPayload *dto.PhotoUpdateRequest) (*dto.GetPhotoResponse, pkg.Error)
	DeletePhoto func(photoId int) (*dto.GetPhotoResponse, pkg.Error)
)

func NewPhotoServiceMock() PhotoService {
	return &photoServiceMock{}
}

// AddPhoto implements PhotoService.
func (psm *photoServiceMock) AddPhoto(userId int, photoPayload *dto.NewPhotoRequest) (*dto.GetPhotoResponse, pkg.Error) {
	return AddPhoto(userId, photoPayload)
}

// GetPhotos implements PhotoService.
func (psm *photoServiceMock) GetPhotos() (*dto.GetPhotoResponse, pkg.Error) {
	return GetPhotos()
}

// UpdatePhoto implements PhotoService.
func (psm *photoServiceMock) UpdatePhoto(photoId int, photoPayload *dto.PhotoUpdateRequest) (*dto.GetPhotoResponse, pkg.Error) {
	return UpdatePhoto(photoId, photoPayload)
}

// DeletePhoto implements PhotoService.
func (psm *photoServiceMock) DeletePhoto(photoId int) (*dto.GetPhotoResponse, pkg.Error) {
	return DeletePhoto(photoId)
}
