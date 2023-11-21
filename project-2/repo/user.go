package repo

import (
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
)

type UserRepository interface {
	Create(userPayload *models.User) (*dto.UserResponse, pkg.Error)
	Fetch(email string) (*models.User, pkg.Error)
	FetchById(userId int) (*models.User, pkg.Error)
	Update(userPayload *models.User) (*dto.UserUpdateResponse, pkg.Error)
	Delete(userId int) pkg.Error
}

var (
	Create    func(userPayload *models.User) (*dto.UserResponse, pkg.Error)
	Fetch     func(email string) (*models.User, pkg.Error)
	Update    func(userPayload *models.User) (*dto.UserUpdateResponse, pkg.Error)
	FetchById func(userId int) (*models.User, pkg.Error)
	Delete    func(userId int) pkg.Error
)

type userRepositoryMock struct {
}

func NewUserRepositoryMock() UserRepository {
	return &userRepositoryMock{}
}

func (urm *userRepositoryMock) Create(userPayload *models.User) (*dto.UserResponse, pkg.Error) {
	return Create(userPayload)
}

func (urm *userRepositoryMock) Fetch(email string) (*models.User, pkg.Error) {
	return Fetch(email)
}

func (urm *userRepositoryMock) Update(userPayload *models.User) (*dto.UserUpdateResponse, pkg.Error) {
	return Update((userPayload))
}

func (urm *userRepositoryMock) FetchById(userId int) (*models.User, pkg.Error) {
	return FetchById(userId)
}

func (urm *userRepositoryMock) Delete(userId int) pkg.Error {
	return Delete(userId)
}