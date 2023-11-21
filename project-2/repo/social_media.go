package repo

import (
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
)

type SocialMediaUserPhotoMapped struct {
	SocialMedia models.SocialMedia
	User        models.User
	Photo       models.Photo
}

type SocialMediaUserPhoto struct {
	SocialMedia models.SocialMedia
	User        models.User
	Photo       models.Photo
}

func (s *SocialMediaUserPhotoMapped) HandleMappingSocialMediaWithUserAndPhoto(socialMediaUserPhoto []SocialMediaUserPhoto) []*dto.GetSocialMedia {
	socialMediasWithUserAndPhoto := []*dto.GetSocialMedia{}

	for _, eachSocialMediaWithUserAndPhoto := range socialMediaUserPhoto {
		socialMediaWithUserAndPhoto := &dto.GetSocialMedia{
			Id:             eachSocialMediaWithUserAndPhoto.SocialMedia.Id,
			Name:           eachSocialMediaWithUserAndPhoto.SocialMedia.Name,
			SocialMediaUrl: eachSocialMediaWithUserAndPhoto.SocialMedia.SocialMediaUrl,
			UserId:         eachSocialMediaWithUserAndPhoto.SocialMedia.UserId,
			CreatedAt:      eachSocialMediaWithUserAndPhoto.SocialMedia.CreatedAt,
			UpdatedAt:      eachSocialMediaWithUserAndPhoto.SocialMedia.UpdatedAt,
			User: dto.SocialMediaUser{
				Id:              eachSocialMediaWithUserAndPhoto.User.Id,
				Username:        eachSocialMediaWithUserAndPhoto.User.Username,
				ProfileImageUrl: eachSocialMediaWithUserAndPhoto.Photo.PhotoUrl,
			},
		}

		socialMediasWithUserAndPhoto = append(socialMediasWithUserAndPhoto, socialMediaWithUserAndPhoto)
	}

	return socialMediasWithUserAndPhoto
}

func (s *SocialMediaUserPhotoMapped) HandleMappingSocialMediaWithUserAndPhotoById(socialMediaUserPhoto SocialMediaUserPhoto) *dto.GetSocialMedia {

	socialMediaWithUserAndPhoto := &dto.GetSocialMedia{
		Id:             socialMediaUserPhoto.SocialMedia.Id,
		Name:           socialMediaUserPhoto.SocialMedia.Name,
		SocialMediaUrl: socialMediaUserPhoto.SocialMedia.SocialMediaUrl,
		UserId:         socialMediaUserPhoto.SocialMedia.UserId,
		CreatedAt:      socialMediaUserPhoto.SocialMedia.CreatedAt,
		UpdatedAt:      socialMediaUserPhoto.SocialMedia.UpdatedAt,
		User: dto.SocialMediaUser{
			Id:              socialMediaUserPhoto.User.Id,
			Username:        socialMediaUserPhoto.User.Username,
			ProfileImageUrl: socialMediaUserPhoto.Photo.PhotoUrl,
		},
	}

	return socialMediaWithUserAndPhoto
}

type socialMediaMock struct {
}

var (
	AddSocialMedia     func(socialMediaPayload *models.SocialMedia) (*dto.NewSocialMediaResponse, pkg.Error)
	DeleteSocialMedia  func(socialMediaId int) pkg.Error
	UpdateSocialMedia  func(socialMediaId int, socialMediaPayload *models.SocialMedia) (*dto.SocialMediaUpdateResponse, pkg.Error)
	GetSocialMediaById func(socialMediaId int) (*dto.GetSocialMedia, pkg.Error)
	GetSocialMedias    func() ([]*dto.GetSocialMedia, pkg.Error)
)

func NewSocialMediaMock() SocialMediaRepository {
	return &socialMediaMock{}
}

func (s *socialMediaMock) AddSocialMedia(socialMediaPayload *models.SocialMedia) (*dto.NewSocialMediaResponse, pkg.Error) {
	return AddSocialMedia(socialMediaPayload)
}

func (s *socialMediaMock) DeleteSocialMedia(socialMediaId int) pkg.Error {
	return DeleteSocialMedia(socialMediaId)
}

func (s *socialMediaMock) UpdateSocialMedia(socialMediaId int, socialMediaPayload *models.SocialMedia) (*dto.SocialMediaUpdateResponse, pkg.Error) {
	return UpdateSocialMedia(socialMediaId, socialMediaPayload)
}

func (s *socialMediaMock) GetSocialMediaById(socialMediaId int) (*dto.GetSocialMedia, pkg.Error) {
	return GetSocialMediaById(socialMediaId)
}

func (s *socialMediaMock) GetSocialMedias() ([]*dto.GetSocialMedia, pkg.Error) {
	return GetSocialMedias()
}

type SocialMediaRepository interface {
	AddSocialMedia(socialMediaPayload *models.SocialMedia) (*dto.NewSocialMediaResponse, pkg.Error)
	UpdateSocialMedia(socialMediaId int, socialMediaPayload *models.SocialMedia) (*dto.SocialMediaUpdateResponse, pkg.Error)
	GetSocialMedias() ([]*dto.GetSocialMedia, pkg.Error)
	GetSocialMediaById(socialMediaId int) (*dto.GetSocialMedia, pkg.Error)
	DeleteSocialMedia(socialMediaId int) pkg.Error
}
