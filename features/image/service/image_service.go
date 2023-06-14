package service

import (
	"be-api/features"
	imageInterface "be-api/features/image"
)

type imageService struct {
	imageRepository imageInterface.ImageRepository
}

// CreateImage implements image.ImageService.
func (is *imageService) CreateImage(image features.ImageEntity) (uint, error) {
	panic("unimplemented")
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
