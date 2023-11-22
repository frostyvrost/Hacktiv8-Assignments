package service

import (
	"project-3/model"
	"project-3/pkg"
	"project-3/repo"

	"github.com/asaskevich/govalidator"
)

type userServiceRepo interface {
	Register(*model.User) (*model.User, pkg.Error)
	Login(*model.LoginCredential) (string, pkg.Error)
	UpdateUser(userID uint, update *model.UserUpdate) (*model.User, pkg.Error)
	DeleteUser(userID uint) (*model.User, pkg.Error)
}

type userService struct{}

var UserService userServiceRepo = &userService{}

func (t *userService) Register(user *model.User) (*model.User, pkg.Error) {
	user.Role = "member"
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	password, err := pkg.HashPass(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = password

	result, err := repo.UserModel.Register(user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (t *userService) Login(login *model.LoginCredential) (string, pkg.Error) {
	if _, err := govalidator.ValidateStruct(login); err != nil {
		return "", pkg.BadRequest(err.Error())
	}

	user, err := repo.UserModel.Login(login)
	if err != nil {
		return "", err
	}

	if isPasswordCorrect := pkg.ComparePass(user.Password, login.Password); !isPasswordCorrect {
		return "", pkg.Unautorized("Invalid email/password")
	}

	token, err := pkg.GenerateToken(user.ID, user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (t *userService) UpdateUser(userID uint, update *model.UserUpdate) (*model.User, pkg.Error) {
	if _, err := govalidator.ValidateStruct(update); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	updatedUser, err := repo.UserModel.UpdateUser(userID, update)

	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (t *userService) DeleteUser(userID uint) (*model.User, pkg.Error) {
	result, err := repo.UserModel.DeleteUser(userID)

	if err != nil {
		return nil, err
	}

	return result, nil
}
