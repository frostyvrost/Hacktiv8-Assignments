package repo

import (
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
)

type UserRepository interface {
	CreateNewUser(userPayLoad *models.User) (*dto.NewUserResponse, pkg.MessageErr)
	GetUserByEmail(userEmail string) (*models.User, pkg.MessageErr)
	GetUserById(userId int) (*models.User, pkg.MessageErr)
	UpdateUser(userPayLoad *models.User) (*dto.UserUpdateResponse, pkg.MessageErr)
	DeleteUser(userId int) pkg.MessageErr
	Admin(userPayLoad *models.User) pkg.MessageErr
}
