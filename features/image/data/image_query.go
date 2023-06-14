package data

import (
	models "be-api/features"
	imageInterface "be-api/features/image"
	"errors"

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
	imageModel := models.ImageEntityToModel(image)

	createOpr := iq.db.Create(&imageModel)
	if createOpr.Error != nil {
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	return imageModel.ID, nil
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
