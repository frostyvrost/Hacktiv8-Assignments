package repo

import (
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"time"
)

type PhotoUser struct {
	Photo models.Photo
	User  models.User
}

type User struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type PhotoUserMapped struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      User      `json:"user"`
}

func (pum *PhotoUserMapped) HandleMappingPhotoWithUser(photoUser []PhotoUser) []PhotoUserMapped {
	photosUserMapped := []PhotoUserMapped{}

	for _, eachPhotoUser := range photoUser {
		photoUserMapped := PhotoUserMapped{
			Id:        eachPhotoUser.Photo.Id,
			Title:     eachPhotoUser.Photo.Title,
			Caption:   eachPhotoUser.Photo.Caption,
			PhotoUrl:  eachPhotoUser.Photo.PhotoUrl,
			UserId:    eachPhotoUser.Photo.UserId,
			CreatedAt: eachPhotoUser.Photo.CreatedAt,
			UpdatedAt: eachPhotoUser.Photo.UpdatedAt,
			User: User{
				Email:    eachPhotoUser.User.Email,
				Username: eachPhotoUser.User.Username,
			},
		}

		photosUserMapped = append(photosUserMapped, photoUserMapped)
	}

	return photosUserMapped
}

func (pum *PhotoUserMapped) HandleMappingPhotoWithUserByPhotoId(photoUser PhotoUser) *PhotoUserMapped {
	return &PhotoUserMapped{
		Id:        photoUser.Photo.Id,
		Title:     photoUser.Photo.Title,
		Caption:   photoUser.Photo.Caption,
		PhotoUrl:  photoUser.Photo.PhotoUrl,
		UserId:    photoUser.Photo.UserId,
		CreatedAt: photoUser.Photo.CreatedAt,
		UpdatedAt: photoUser.Photo.UpdatedAt,
		User: User{
			Email:    photoUser.User.Email,
			Username: photoUser.User.Username,
		},
	}
}

var (
	AddPhoto    func(photoPayload *models.Photo) (*dto.PhotoResponse, pkg.Error)
	GetPhotos   func() ([]PhotoUserMapped, pkg.Error)
	GetPhotoId  func(photoId int) (*PhotoUserMapped, pkg.Error)
	UpdatePhoto func(photoId int, photoPayload *models.Photo) (*dto.PhotoUpdateResponse, pkg.Error)
	DeletePhoto func(photoId int) pkg.Error
)

type photoRepositoryMock struct {
}

func NewPhotoRepositoryMock() PhotoRepository {
	return &photoRepositoryMock{}
}

func (prm *photoRepositoryMock) AddPhoto(photoPayload *models.Photo) (*dto.PhotoResponse, pkg.Error) {
	return AddPhoto(photoPayload)
}

func (prm *photoRepositoryMock) GetPhotos() ([]PhotoUserMapped, pkg.Error) {
	return GetPhotos()
}

func (prm *photoRepositoryMock) GetPhotoId(photoId int) (*PhotoUserMapped, pkg.Error) {
	return GetPhotoId(photoId)
}

func (prm *photoRepositoryMock) UpdatePhoto(photoId int, photoPayload *models.Photo) (*dto.PhotoUpdateResponse, pkg.Error) {
	return UpdatePhoto(photoId, photoPayload)
}

func (prm *photoRepositoryMock) DeletePhoto(photoId int) pkg.Error {
	return DeletePhoto(photoId)
}

type PhotoRepository interface {
	AddPhoto(photoPayload *models.Photo) (*dto.PhotoResponse, pkg.Error)
	GetPhotos() ([]PhotoUserMapped, pkg.Error)
	GetPhotoId(photoId int) (*PhotoUserMapped, pkg.Error)
	UpdatePhoto(photoId int, photoPayload *models.Photo) (*dto.PhotoUpdateResponse, pkg.Error)
	DeletePhoto(photoId int) pkg.Error
}
