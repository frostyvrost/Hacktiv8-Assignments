package service

import (
	"net/http"
	"project-2/dto"
	"project-2/models"
	"project-2/pkg"
	"project-2/repo"
)

type userServiceMock struct {
}

type UserService interface {
	Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, pkg.Error)
	Get(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, pkg.Error)
	Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, pkg.Error)
	Remove(userId int) (*dto.GetUserResponse, pkg.Error)
}

type userServiceImpl struct {
	ur repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &userServiceImpl{
		ur: userRepo,
	}
}

// Add implements UserService.
func (u *userServiceImpl) Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, pkg.Error) {

	err := pkg.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: userPayload.Username,
		Email:    userPayload.Email,
		Age:      userPayload.Age,
		Password: userPayload.Password,
	}

	user.HashPassword()

	response, err := u.ur.Create(user)

	if err != nil {
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusCreated,
		Message:    "create new user successfully",
		Data:       response,
	}, nil
}

// Get implements UserService.
func (us *userServiceImpl) Get(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, pkg.Error) {

	err := pkg.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user, err := us.ur.Fetch(userPayload.Email)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequestError("invalid email/password")
		}
		return nil, err
	}

	isValidPassword := user.ComparePassword(userPayload.Password)

	if !isValidPassword {
		return nil, pkg.NewBadRequestError("invalid email/password")
	}
	token := user.GenerateToken()

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "successfully loged in",
		Data: dto.TokenResponse{
			Token: token,
		},
	}, nil
}

// Edit implements UserService.
func (u *userServiceImpl) Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, pkg.Error) {

	err := pkg.ValidateStruct(userPayload)

	if err != nil {
		return nil, err
	}

	user, err := u.ur.FetchById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequestError("invalid user")
		}
		return nil, err
	}

	if user.Id != userId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	usr := &models.User{
		Id:       userId,
		Email:    userPayload.Email,
		Username: userPayload.Username,
	}

	response, err := u.ur.Update(usr)

	if err != nil {
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "user has been successfully updated",
		Data:       response,
	}, nil
}

// Remove implements UserService.
func (u *userServiceImpl) Remove(userId int) (*dto.GetUserResponse, pkg.Error) {

	user, err := u.ur.FetchById(userId)

	if err != nil {
		if err.Status() == http.StatusNotFound {
			return nil, pkg.NewBadRequestError("invalid user")
		}
		return nil, err
	}

	if user.Id != userId {
		return nil, pkg.NewNotFoundError("invalid user")
	}

	err = u.ur.Delete(userId)

	if err != nil {
		return nil, err
	}

	return &dto.GetUserResponse{
		StatusCode: http.StatusOK,
		Message:    "Your account has been successfully deleted",
		Data:       nil,
	}, nil
}

var (
	Add    func(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, pkg.Error)
	Get    func(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, pkg.Error)
	Edit   func(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, pkg.Error)
	Remove func(userId int) (*dto.GetUserResponse, pkg.Error)
)

func NewUserServiceMock() UserService {
	return &userServiceMock{}
}

// Add implements UserService.
func (usm *userServiceMock) Add(userPayload *dto.NewUserRequest) (*dto.GetUserResponse, pkg.Error) {
	return Add(userPayload)
}

// Get implements UserService.
func (usm *userServiceMock) Get(userPayload *dto.UserLoginRequest) (*dto.GetUserResponse, pkg.Error) {
	return Get(userPayload)
}

// Edit implements UserService.
func (usm *userServiceMock) Edit(userId int, userPayload *dto.UserUpdateRequest) (*dto.GetUserResponse, pkg.Error) {
	return Edit(userId, userPayload)
}

// Remove implements UserService.
func (usm *userServiceMock) Remove(userId int) (*dto.GetUserResponse, pkg.Error) {
	return Remove(userId)
}
