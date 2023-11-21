package repo

import (
	"project-4/database"
	"project-4/models"
	"project-4/pkg"
)

type userDomainRepo interface {
	Register(*models.User) (*models.User, pkg.Error)
	Login(*models.LoginCredential) (*models.User, pkg.Error)
	GetUserById(int) (*models.User, pkg.Error)
	UpdateBalance(*models.User, int) (int, pkg.Error)
}

type userRepo struct{}

var UserRepo userDomainRepo = &userRepo{}

func (u *userRepo) Register(user *models.User) (*models.User, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&user).Error

	if err != nil {
		return nil, pkg.InternalServerError("Something went wrong")
	}

	return user, nil
}

func (u *userRepo) Login(userLogin *models.LoginCredential) (*models.User, pkg.Error) {
	db := database.GetDB()

	var user models.User

	err := db.Where("email = ?", userLogin.Email).First(&user).Error

	if err != nil {
		return nil, pkg.Unauthorized("Invalid email/password")
	}

	return &user, nil
}

func (u *userRepo) GetUserById(userId int) (*models.User, pkg.Error) {
	db := database.GetDB()
	var user models.User

	err := db.First(&user, userId).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return &user, nil
}

func (u *userRepo) UpdateBalance(userWithUpdatedBalance *models.User, userId int) (int, pkg.Error) {
	db := database.GetDB()
	var user models.User

	err := db.First(&user, userId).Error

	if err != nil {
		return 0, pkg.ParseError(err)
	}

	db.Model(&user).Updates(userWithUpdatedBalance)

	return user.Balance, nil
}
