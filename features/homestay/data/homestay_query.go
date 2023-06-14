package data

import (
	models "be-api/features"
	homestayInterface "be-api/features/homestay"
	"errors"

	"gorm.io/gorm"
)

type homestayQuery struct {
	db *gorm.DB
}

// Delete implements homestay.HomestayRepository.
func (hq *homestayQuery) Delete(homestayID uint) error {
	deleteOpr := hq.db.Delete(&models.Homestay{}, homestayID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete homestay")
	}

	return nil
}

// Insert implements homestay.HomestayRepository.
func (hq *homestayQuery) Insert(homestay models.HomestayEntity) (uint, error) {
	homestayModel := models.HomestayEntityToModel(homestay)

	createOpr := hq.db.Create(&homestayModel)
	if createOpr.Error != nil {
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	return homestayModel.ID, nil
}

// Select implements homestay.HomestayRepository.
func (hq *homestayQuery) Select(homestayID uint) (models.HomestayEntity, error) {
	var homestay models.Homestay

	queryResult := hq.db.Preload("Bookings").Preload("Reviews").Preload("Images").First(&homestay, homestayID)
	if queryResult.Error != nil {
		return models.HomestayEntity{}, queryResult.Error
	}

	var rating float64
	if err := hq.db.Raw("SELECT AVG(ratings) FROM reviews WHERE homestay_id = ?", homestayID).Scan(&rating).Error; err != nil {
		return models.HomestayEntity{}, err
	}
	homestay.Rating = rating

	homestayEntity := models.HomestayModelToEntity(homestay)

	return homestayEntity, nil
}

// SelectAll implements homestay.HomestayRepository.
func (hq *homestayQuery) SelectAll() ([]models.HomestayEntity, error) {
	var homestays []models.Homestay

	queryResult := hq.db.Preload("Bookings").Preload("Reviews").Preload("Images").Find(&homestays)

	if queryResult.Error != nil {
		return []models.HomestayEntity{}, queryResult.Error
	}

	var homestayEntities []models.HomestayEntity
	for _, homestay := range homestays {
		var rating float64
		hq.db.Raw("SELECT AVG(ratings) FROM reviews WHERE homestay_id = ?", homestay.ID).Scan(&rating)
		homestay.Rating = rating
		homestayEntity := models.HomestayModelToEntity(homestay)
		homestayEntities = append(homestayEntities, homestayEntity)
	}

	return homestayEntities, nil
}

// Update implements homestay.HomestayRepository.
func (hq *homestayQuery) Update(homestayID uint, updatedHomestay models.HomestayEntity) error {
	var homestay models.Homestay

	queryResult := hq.db.First(&homestay, homestayID)
	if queryResult.Error != nil {
		return errors.New(queryResult.Error.Error() + ", failed to get homestay")
	}

	updatedHomestayModel := models.HomestayEntityToModel(updatedHomestay)
	updateOpr := hq.db.Model(&homestay).Updates(updatedHomestayModel)
	if updateOpr.Error != nil {
		return errors.New(updateOpr.Error.Error() + ", failed to update homestay")
	}

	return nil
}

func New(db *gorm.DB) homestayInterface.HomestayRepository {
	return &homestayQuery{
		db: db,
	}
}
