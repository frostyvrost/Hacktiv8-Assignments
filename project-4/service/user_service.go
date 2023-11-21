package service

import (
	"project-4/models"
	"project-4/pkg"
	"project-4/repo"

	"github.com/asaskevich/govalidator"
)

type userServiceRepo interface {
	Register(*models.User) (*models.User, pkg.Error)
	Login(*models.LoginCredential) (string, pkg.Error)
	UpdateBalance(*models.BalanceUpdate, int) (int, pkg.Error)
}

type userService struct{}

var UserService userServiceRepo = &userService{}

func (u *userService) Register(user *models.User) (*models.User, pkg.Error) {
	if _, err := govalidator.ValidateStruct(user); err != nil {
		return nil, pkg.BadRequest(err.Error())
	}

	password, err := pkg.HashPassword(user.Password)

	if err != nil {
		return nil, err
	}

	user.Password = password

	userResponse, err := repo.UserRepo.Register(user)

	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (u *userService) Login(userLogin *models.LoginCredential) (string, pkg.Error) {
	if _, err := govalidator.ValidateStruct(userLogin); err != nil {
		return "", pkg.BadRequest(err.Error())
	}

	user, err := repo.UserRepo.Login(userLogin)

	if err != nil {
		return "", err
	}

	if isPasswordCorrect := pkg.ComparePassword(userLogin.Password, user.Password); !isPasswordCorrect {
		return "", pkg.Unauthorized("Invalid email/password")
	}

	token, err := pkg.GenerateToken(user.ID, user.Email, user.Role)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *userService) UpdateBalance(balance *models.BalanceUpdate, userId int) (int, pkg.Error) {
	if _, err := govalidator.ValidateStruct(balance); err != nil {
		return 0, pkg.BadRequest(err.Error())
	}

	user, err := repo.UserRepo.GetUserById(userId)

	if err != nil {
		return 0, err
	}

	user.Balance += balance.Balance

	updatedBalance, err := repo.UserRepo.UpdateBalance(user, userId)

	if err != nil {
		return 0, err
	}

	return updatedBalance, nil
}
