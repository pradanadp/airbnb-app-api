package data

import (
	models "be-api/features"
	imageInterface "be-api/features/image"

	"gorm.io/gorm"
)

type imageQuery struct {
	db *gorm.DB
}

// Delete implements image.ImageRepository.
func (iq *imageQuery) Delete(imageID uint) error {
	panic("unimplemented")
}

// Insert implements image.ImageRepository.
func (iq *imageQuery) Insert(image models.ImageEntity) (uint, error) {
	panic("unimplemented")
}

// Select implements image.ImageRepository.
func (iq *imageQuery) Select(imageID uint) (models.ImageEntity, error) {
	panic("unimplemented")
}

func New(db *gorm.DB) imageInterface.ImageRepository {
	return &imageQuery{
		db: db,
	}
}
