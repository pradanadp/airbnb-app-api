package controller

import (
	"be-api/app/middlewares"
	aws "be-api/aws"
	"be-api/features"
	"be-api/features/user"
	"be-api/utils"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	MaxFileSize = 1 << 20 // 1 MB
)

type UserController struct {
	userService user.UserServiceInterface
}

func New(service user.UserServiceInterface) *UserController {
	return &UserController{
		userService: service,
	}
}

func (handler *UserController) LoginUser(c echo.Context) error {
	var payload features.LoginUser
	if errBind := c.Bind(&payload); errBind != nil {
		if errBind == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("Login Failed. Email anda tidak terdaftar", nil))
		}
	}

	userId, err := handler.userService.LoginUser(payload.Email, payload.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusUnauthorized, utils.FailResponse("Input tidak valid, harap isi email dan password sesuai ketentuan", nil))
		} else if strings.Contains(err.Error(), "email tidak terdaftar") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("Email yang anda berikan tidak terdaftar", nil))
		} else {
			return c.JSON(http.StatusUnauthorized, utils.FailResponse("password yang anda berikan tidak valid", nil))
		}
	}
	user, err := handler.userService.GetUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data tidak tersedia ", nil))
	}
	mapUser := EntityToResponse(user)

	accessToken, err := middlewares.CreateToken(userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("successfully", map[string]any{
		"accessToken": accessToken,
		"user":        mapUser,
	}))
}

func (handler *UserController) AddUser(c echo.Context) error {

	payload := features.UserEntity{}
	if err := c.Bind(&payload); err != nil {
		if err == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error bind payload "+err.Error(), nil))
		}
	}

	id, err := handler.userService.AddUser(payload)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error validation payload "+err.Error(), nil))
		} else {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("email tidak tersedia "+err.Error(), nil))
		}
	}
	user, err := handler.userService.GetUser(int(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data tidak tersedia ", nil))
	}
	mapUser := EntityToResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("successfully", mapUser))
}

func (handler *UserController) GetUser(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	user, err := handler.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data tidak tersedia ", nil))
	}
	user.FullName = user.FirstName + " " + user.LastName
	mapUser := EntityToReadResponse(user)
	return c.JSON(http.StatusOK, utils.SuccessResponse("successfully", mapUser))
}

func (handler *UserController) DeleteUser(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	errId := handler.userService.GetId(id)
	if errId != nil {
		return c.JSON(http.StatusNotFound, utils.FailWithoutDataResponse("fail to id user not found"))
	}

	err := handler.userService.DeleteUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("delete Fail to Delete akun User", nil))
	}
	return c.JSON(http.StatusOK, utils.SuccessWhitoutResponse("Success delete akun User"))
}

func (handler *UserController) UpdateUser(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	errId := handler.userService.GetId(id)
	if errId != nil {
		return c.JSON(http.StatusNotFound, utils.FailWithoutDataResponse("fail to id user not found"))
	}

	update := features.UserEntity{}
	if err := c.Bind(&update); err != nil {
		if err == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error bind payload "+err.Error(), nil))
		}
	}

	errUpdate := handler.userService.Update(update, uint(id))
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	user, err := handler.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data tidak tersedia ", nil))
	}
	mapUser := EntityToResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("successfully", mapUser))

}

func (handler *UserController) UpgradeUser(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	errId := handler.userService.GetId(id)
	if errId != nil {
		return c.JSON(http.StatusNotFound, utils.FailWithoutDataResponse("fail to id user not found"))
	}

	upgrade := features.UserEntity{}
	if err := c.Bind(&upgrade); err != nil {
		if err == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error bind payload "+err.Error(), nil))
		}
	}

	errUpgrade := handler.userService.UpgradeUser(upgrade, uint(id))
	if errUpgrade != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error "+errUpgrade.Error(), nil))
	}

	user, err := handler.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data tidak tersedia ", nil))
	}
	mapUser := EntityToResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("successfully", mapUser))
}

func (handler *UserController) UploadProfilePicture(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	errId := handler.userService.GetId(id)
	if errId != nil {
		return c.JSON(http.StatusNotFound, utils.FailWithoutDataResponse("User ID not found"))
	}

	awsService := aws.InitS3()

	file, err := c.FormFile("profile_picture")
	if err != nil {
		return err
	}

	// Check file size before opening it
	fileSize := file.Size
	if fileSize > MaxFileSize {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("File size exceeds the limit of 1 MB", nil))
	}

	path := "profile-picture/" + file.Filename
	fileContent, err := file.Open()
	if err != nil {
		return err
	}
	defer fileContent.Close()

	err = awsService.UploadFile(path, fileContent)
	if err != nil {
		return err
	}

	var updatedUser features.UserEntity
	updatedUser.ProfilePicture = fmt.Sprintf("https://aws-airbnb-api.s3.ap-southeast-2.amazonaws.com/profile-picture/%s", filepath.Base(file.Filename))

	errUpdate := handler.userService.Update(updatedUser, uint(id))
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	user, err := handler.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data not found", nil))
	}
	userResponse := EntityToResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("profile picture updated successfully", userResponse))
}

func (handler *UserController) UploadHostDoc(c echo.Context) error {
	id := middlewares.ExtractTokenUserId(c)
	errId := handler.userService.GetId(id)
	if errId != nil {
		return c.JSON(http.StatusNotFound, utils.FailWithoutDataResponse("User ID not found"))
	}

	awsService := aws.InitS3()

	file, err := c.FormFile("host_document")
	if err != nil {
		return err
	}

	// Check file size before opening it
	fileSize := file.Size
	if fileSize > MaxFileSize {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("File size exceeds the limit of 1 MB", nil))
	}

	path := "host-doc/" + file.Filename
	fileContent, err := file.Open()
	if err != nil {
		return err
	}
	defer fileContent.Close()

	err = awsService.UploadFile(path, fileContent)
	if err != nil {
		return err
	}

	var updatedUser features.UserEntity
	updatedUser.HostDocument = fmt.Sprintf("https://aws-airbnb-api.s3.ap-southeast-2.amazonaws.com/host-doc/%s", filepath.Base(file.Filename))

	errUpdate := handler.userService.Update(updatedUser, uint(id))
	if errUpdate != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	user, err := handler.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data not found", nil))
	}
	userResponse := EntityToResponse(user)

	return c.JSON(http.StatusOK, utils.SuccessResponse("host document added successfully", userResponse))
}
