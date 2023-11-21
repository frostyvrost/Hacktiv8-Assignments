package handler

import (
	"net/http"
	"project-3/dto"
	"project-3/models"
	"project-3/pkg"
	"project-3/service"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

// Register implements UserHandler
// Register godoc
// @Summary User register
// @Description User register
// @Tags Users
// @Accept json
// @Produce json
// @Param RequestBody body dto.NewUserRequest true "body request for user register"
// @Success 201 {object} dto.NewUserResponse
// @Router /users/register [post]
func (uh *userHandler) Register(ctx *gin.Context) {
	newUserRequest := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := pkg.NewUnprocessibleModelsError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Register(newUserRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

// Login implements UserHandler.
// Login godoc
// @Summary User login
// @Description User login
// @Tags Users
// @Accept json
// @Produce json
// @Param RequestBody body dto.UserLoginRequest true "body request for user login"
// @Success 200 {object} dto.UserLoginResponse
// @Router /users/login [post]
func (uh *userHandler) Login(ctx *gin.Context) {
	userLoginRequest := &dto.UserLoginRequest{}

	if err := ctx.ShouldBindJSON(&userLoginRequest); err != nil {
		errBindJson := pkg.NewUnprocessibleModelsError("invalid request body")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Login(userLoginRequest)

	if err != nil {
		ctx.JSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Updateements UserHandler.
// Update godoc
// @Summary User Update
// @Description User Update
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Param RequestBody body dto.UserUpdateRequest true "body request for user login"
// @Success 200 {object} dto.UserUpdateResponse
// @Router /users/update-account [put]
func (uh *userHandler) Update(ctx *gin.Context) {
	userData, ok := ctx.MustGet("userData").(models.User)

	var userUpdateRequest = &dto.UserUpdateRequest{}

	if !ok {
		errData := pkg.NewBadRequest("Failed get user data!!")
		ctx.AbortWithStatusJSON(errData.Status(), errData)
		return
	}
	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		errData := pkg.NewUnprocessibleModelsError(err.Error())
		ctx.AbortWithStatusJSON(errData.Status(), errData)
		return
	}

	response, err := uh.userService.Update(userData.Id, userUpdateRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// Delete implements UserHandler.
// Delete godoc
// @Summary Delete User
// @Description Delete Users
// @Tags Users
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} dto.DeleteResponse
// @Router /users/delete-account [delete]
func (uh *userHandler) Delete(ctx *gin.Context) {
	user, ok := ctx.MustGet("userData").(models.User)

	if !ok {
		errData := pkg.NewBadRequest("Failed get user data!!")
		ctx.AbortWithStatusJSON(errData.Status(), errData)
		return
	}
	response, err := uh.userService.Delete(user.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uh *userHandler) Admin(ctx *gin.Context) {
	newUserRequest := &dto.NewUserRequest{}

	if err := ctx.ShouldBindJSON(&newUserRequest); err != nil {
		errBindJson := pkg.NewUnprocessibleModelsError("invalid request")

		ctx.JSON(errBindJson.Status(), errBindJson)
		return
	}

	response, err := uh.userService.Admin(newUserRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(err.Status(), err)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
