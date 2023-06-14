package image

import (
	models "be-api/features"
)

type ImageRepository interface {
	Insert(image models.ImageEntity) (uint, error)
	Select(imageID uint) (models.ImageEntity, error)
	Delete(imageID uint) error
}

type ImageService interface {
	CreateImage(image models.ImageEntity) (uint, error)
	GetImage(imageID uint) (models.ImageEntity, error)
	DeleteImage(imageID uint) error
}
