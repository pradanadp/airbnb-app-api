package service

import (
	models "be-api/features"

	imageInterface "be-api/features/image"
	"errors"
	"fmt"
)

type imageService struct {
	imageRepository imageInterface.ImageRepository
}

// CreateImage implements image.ImageService.
func (is *imageService) CreateImage(image models.ImageEntity) (uint, error) {
	switch {
	case image.HomestayID == 0:
		return 0, errors.New("error, homestay ID is required")
	case image.Link == "":
		return 0, errors.New("error, image link is required")
	}

	imageID, err := is.imageRepository.Insert(image)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return imageID, nil
}

// DeleteImage implements image.ImageService.
func (is *imageService) DeleteImage(imageID uint) error {
	err := is.imageRepository.Delete(imageID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

// GetImage implements image.ImageService.
func (is *imageService) GetImage(homestayID uint) ([]models.ImageEntity, error) {
	imageEntity, err := is.imageRepository.SelectAll(homestayID)
	if err != nil {
		return []models.ImageEntity{}, fmt.Errorf("error: %v", err)
	}

	return imageEntity, nil
}

func New(repo imageInterface.ImageRepository) imageInterface.ImageService {
	return &imageService{
		imageRepository: repo,
	}
}
