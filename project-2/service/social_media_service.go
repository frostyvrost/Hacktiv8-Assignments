package service

import (
	"net/http"
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/repo"
)

type socialMediaMock struct {
}

type SocialMediaService interface {
	AddSocialMedia(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error)
	GetSocialMedias() (*dto.GetSocialMediaHttpResponse, pkg.Error)
	UpdateSocialMedia(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error)
	DeleteSocialMedia(socialMediaId int) (*dto.GetSocialMediaResponse, pkg.Error)
}

type socialMediaServiceImpl struct {
	sr repo.SocialMediaRepository
}

func NewSocialMediaService(socialMediaRepo repo.SocialMediaRepository) SocialMediaService {
	return &socialMediaServiceImpl{
		sr: socialMediaRepo,
	}
}

// AddSocialMedia implements SocialMediaService.
func (s *socialMediaServiceImpl) AddSocialMedia(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error) {

	err := pkg.ValidateStruct(socialMediaPayload)

	if err != nil {
		return nil, err
	}

	socialMedia := &models.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaUrl: socialMediaPayload.SocialMediaUrl,
		UserId:         userId,
	}

	data, err := s.sr.AddSocialMedia(socialMedia)

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaResponse{
		StatusCode: http.StatusCreated,
		Message:    "new social media successfully added",
		Data:       data,
	}, nil
}

// DeleteSocialMedia implements SocialMediaService.
func (s *socialMediaServiceImpl) DeleteSocialMedia(socialMediaId int) (*dto.GetSocialMediaResponse, pkg.Error) {

	err := s.sr.DeleteSocialMedia(socialMediaId)

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaResponse{
		StatusCode: http.StatusOK,
		Message:    "Your social media has been successfully deleted",
		Data:       nil,
	}, nil
}

// GetSocialMedias implements SocialMediaService.
func (s *socialMediaServiceImpl) GetSocialMedias() (*dto.GetSocialMediaHttpResponse, pkg.Error) {

	socialMedia, err := s.sr.GetSocialMedias()

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaHttpResponse{
		StatusCode:  http.StatusOK,
		Message:     "social medias successfully fetched",
		SocialMedia: socialMedia,
	}, nil
}

// UpdateSocialMedia implements SocialMediaService.
func (s *socialMediaServiceImpl) UpdateSocialMedia(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error) {

	err := pkg.ValidateStruct(socialMediaPayload)

	if err != nil {
		return nil, err
	}

	socialMedia := &models.SocialMedia{
		Name:           socialMediaPayload.Name,
		SocialMediaUrl: socialMediaPayload.SocialMediaUrl,
	}

	data, err := s.sr.UpdateSocialMedia(socialMediaId, socialMedia)

	if err != nil {
		return nil, err
	}

	return &dto.GetSocialMediaResponse{
		StatusCode: http.StatusOK,
		Message:    "social media successfully updated",
		Data:       data,
	}, nil
}

var (
	AddSocialMedia    func(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error)
	DeleteSocialMedia func(socialMediaId int) (*dto.GetSocialMediaResponse, pkg.Error)
	GetSocialMedias   func() (*dto.GetSocialMediaHttpResponse, pkg.Error)
	UpdateSocialMedia func(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error)
)

func NewSocialMediaMock() SocialMediaService {
	return &socialMediaMock{}
}

// AddSocialMedia implements SocialMediaService.
func (s *socialMediaMock) AddSocialMedia(userId int, socialMediaPayload *dto.NewSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error) {
	return AddSocialMedia(userId, socialMediaPayload)
}

// DeleteSocialMedia implements SocialMediaService.
func (s *socialMediaMock) DeleteSocialMedia(socialMediaId int) (*dto.GetSocialMediaResponse, pkg.Error) {
	return DeleteSocialMedia(socialMediaId)
}

// GetSocialMedias implements SocialMediaService.
func (s *socialMediaMock) GetSocialMedias() (*dto.GetSocialMediaHttpResponse, pkg.Error) {
	return GetSocialMedias()
}

// UpdateSocialMedia implements SocialMediaService.
func (s *socialMediaMock) UpdateSocialMedia(socialMediaId int, socialMediaPayload *dto.UpdateSocialMediaRequest) (*dto.GetSocialMediaResponse, pkg.Error) {
	return UpdateSocialMedia(socialMediaId, socialMediaPayload)
}
