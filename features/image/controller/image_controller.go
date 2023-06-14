package controller

import (
	aws "be-api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UploadProfilePicture(c echo.Context) error {
	awsService := aws.InitS3()

	file, err := c.FormFile("profile_picture")
	if err != nil {
		return err
	}

	path := "profile-picture/" + file.Filename
	err = awsService.UploadFile(path, file.Filename)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "profile picture updated successfully",
	})
}

func UploadHostDoc(c echo.Context) error {
	awsService := aws.InitS3()

	file, err := c.FormFile("host_document")
	if err != nil {
		return err
	}

	path := "host-doc/" + file.Filename
	err = awsService.UploadFile(path, file.Filename)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success upload file",
	})
}

func UploadHomestayPhotos(c echo.Context) error {
	awsService := aws.InitS3()

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["link"]
	for _, file := range files {
		path := "homestay-photos/" + file.Filename
		err = awsService.UploadFile(path, file.Filename)
		if err != nil {
			return err
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success upload files",
	})
}
