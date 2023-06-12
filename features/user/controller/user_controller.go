package controller

import (
	"be-api/app/middlewares"
	"be-api/features"
	"be-api/features/user"
	"be-api/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
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
		}
		if strings.Contains(err.Error(), "email tidak terdaftar") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("Email yang anda berikan tidak terdaftar", nil))
		}
		if strings.Contains(err.Error(), "password tidak cocok") {
			return c.JSON(http.StatusUnauthorized, utils.FailResponse("password yang anda berikan tidak valid", nil))
		}
	}

	accessToken, err := middlewares.CreateToken(userId)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, utils.SuccessResponse("successfully", accessToken))
}

func (handler *UserController) AddUser(c echo.Context) error {

	payload := features.UserEntity{}
	fmt.Println(payload)
	if err := c.Bind(&payload); err != nil {
		if err == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error bind payload " + err.Error(), nil))
		} 
	}

	err := handler.userService.AddUser(payload); if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error validation payload " + err.Error(), nil))
		} else if strings.Contains(err.Error(), "Duplicate entry") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("email tidak tersedia " + err.Error(), nil))
		}                   
	}
	return c.JSON(http.StatusOK, utils.SuccessWhitoutResponse("successfully"))
}

func (handler *UserController) GetUser(c echo.Context) error {
	id := middlewares.ExtracTokenUserId(c)
	user, err := handler.userService.GetUser(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("data tidak tersedia ", nil))
	}
	mapUser := EntityToResponse(user)
	return c.JSON(http.StatusOK, utils.SuccessResponse("successfully", mapUser))
}