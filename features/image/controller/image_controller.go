package controller

import (
	aws "be-api/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

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
