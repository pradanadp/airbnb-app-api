package service

import (
	"be-api/features"
	imageInterface "be-api/features/image"
	"errors"
	"fmt"
)

type imageService struct {
	imageRepository imageInterface.ImageRepository
}

// CreateImage implements image.ImageService.
func (is *imageService) CreateImage(image features.ImageEntity) (uint, error) {
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
	panic("unimplemented")
}

// GetImage implements image.ImageService.
func (is *imageService) GetImage(imageID uint) (features.ImageEntity, error) {
	panic("unimplemented")
}

func New(repo imageInterface.ImageRepository) imageInterface.ImageService {
	return &imageService{
		imageRepository: repo,
	}
}
