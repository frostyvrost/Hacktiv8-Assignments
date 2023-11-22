package repo

import (
	"project-3/database"
	"project-3/model"
	"project-3/pkg"
)

type userModelRepo interface {
	Register(*model.User) (*model.User, pkg.Error)
	Login(*model.LoginCredential) (*model.User, pkg.Error)
	UpdateUser(userID uint, update *model.UserUpdate) (*model.User, pkg.Error)
	DeleteUser(userID uint) (*model.User, pkg.Error)
}

type userModel struct{}

var UserModel userModelRepo = &userModel{}

func (t *userModel) Register(user *model.User) (*model.User, pkg.Error) {
	db := database.GetDB()

	err := db.Create(&user).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	return user, nil
}

func (t *userModel) Login(login *model.LoginCredential) (*model.User, pkg.Error) {
	db := database.GetDB()

	var user model.User
	err := db.Where("email = ?", login.Email).First(&user).Error

	if err != nil {
		return nil, pkg.Unautorized("Invalid email/password")
	}

	return &user, nil
}

func (t *userModel) UpdateUser(userID uint, update *model.UserUpdate) (*model.User, pkg.Error) {
	db := database.GetDB()

	var user model.User
	err := db.First(&user, userID).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Model(&user).Updates(update)

	return &user, nil
}

func (t *userModel) DeleteUser(userID uint) (*model.User, pkg.Error) {
	db := database.GetDB()

	var user model.User
	err := db.First(&user, userID).Error

	if err != nil {
		return nil, pkg.ParseError(err)
	}

	db.Delete(&user)

	return &user, nil
}
