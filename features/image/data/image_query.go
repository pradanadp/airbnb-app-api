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
	deleteOpr := iq.db.Delete(&models.Image{}, imageID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete image")
	}

	return nil
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
func (iq *imageQuery) SelectAll(homestayID uint) ([]models.ImageEntity, error) {
	var images []models.Image
	queryResult := iq.db.Preload("Users").Where("homestay_id = ?", homestayID).Find(&images)
	if queryResult.Error != nil {
		return []models.ImageEntity{}, queryResult.Error
	}

	var imageEntities []models.ImageEntity
	for _, image := range images {
		imageEntity := models.ImageModelToEntity(image)
		imageEntities = append(imageEntities, imageEntity)
	}

	return imageEntities, nil
}

func New(db *gorm.DB) imageInterface.ImageRepository {
	return &imageQuery{
		db: db,
	}
}
