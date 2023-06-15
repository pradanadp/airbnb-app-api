package controller

import (
	aws "be-api/aws"
	models "be-api/features"
	imageInterface "be-api/features/image"
	"be-api/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/labstack/echo/v4"
)

const (
	MaxFileSize = 1 << 20 // 1 MB
)

type imageController struct {
	imageService imageInterface.ImageService
}

func New(service imageInterface.ImageService) *imageController {
	return &imageController{
		imageService: service,
	}
}

func (ic *imageController) UploadHomestayPhotos(c echo.Context) error {
	idParam := c.Param("homestay_id")
	homestayID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid homestay ID", nil))
	}

	awsService := aws.InitS3()
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	for _, file := range files {
		fileSize := file.Size
		if fileSize > MaxFileSize {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("File size exceeds the limit of 1 MB", nil))
		}

		fileType := file.Header.Get("Content-Type")

		path := "homestay-photos/" + file.Filename
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		err = awsService.UploadFile(path, fileType, src)
		if err != nil {
			return err
		}

		image := models.ImageEntity{
			HomestayID: uint(homestayID),
			Link: fmt.Sprintf(
				"https://aws-airbnb-api.s3.ap-southeast-2.amazonaws.com/homestay-photos/%s", filepath.Base(file.Filename),
			),
		}

		_, err = ic.imageService.CreateImage(image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to create image", nil))
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success upload files",
	})
}

func (ic *imageController) UploadHomestayPhotosLocal(c echo.Context) error {
	idParam := c.Param("homestay_id")
	homestayID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid homestay ID", nil))
	}

	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	files := form.File["files"]
	for _, file := range files {
		fileSize := file.Size
		if fileSize > MaxFileSize {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("File size exceeds the limit of 1 MB", nil))
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dstPath := filepath.Join("homestay-photos", file.Filename)

		dst, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		image := models.ImageEntity{
			HomestayID: uint(homestayID),
			Link: fmt.Sprintf(
				"https://aws-airbnb-api.s3.ap-southeast-2.amazonaws.com/%s", filepath.Base(dstPath),
			),
		}

		err = c.Bind(&image)
		if err != nil {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind image data", nil))
		}

		_, err = ic.imageService.CreateImage(image)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to create image", nil))
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success upload files",
	})
}

func (ic *imageController) DeleteImage(c echo.Context) error {
	homestayID, err := strconv.Atoi(c.Param("homestay_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid homestay ID", nil))
	}

	imageID, err := strconv.Atoi(c.Param("image_id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid image ID", nil))
	}

	// Delete image
	err = ic.imageService.DeleteImage(uint(imageID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	// Get All images
	images, err := ic.imageService.GetImage(uint(homestayID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("images not found", nil))
	}

	var imageEntities []ImageResponse
	for _, image := range images {
		imageEntities = append(imageEntities, ImageEntityToResponse(image))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("image deleted successfully", imageEntities))
}
