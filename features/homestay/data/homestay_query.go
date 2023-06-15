package data

import (
	models "be-api/features"
	homestayInterface "be-api/features/homestay"
	"database/sql"
	"errors"

	"gorm.io/gorm"
)

type homestayQuery struct {
	db *gorm.DB
}

// Delete implements homestay.HomestayRepository.
func (hq *homestayQuery) Delete(homestayID uint) error {
	homestay, err := hq.Select(homestayID)
	if err != nil {
		return err
	}

	// Retrieve the user from the database
	user := models.User{}
	err = hq.db.Model(&user).Where("id = ?", homestay.HostID).First(&user).Error
	if err != nil {
		return errors.New("error to retrieve user data")
	}

	// Increment the hosting_count
	user.HostingCount--

	// Update the hosting_count column in the users table
	err = hq.db.Model(&user).Where("id = ?", homestay.HostID).UpdateColumn("hosting_count", user.HostingCount).Error
	if err != nil {
		return errors.New("failed to update user hosting count")
	}

	// Delete homestay data from database
	deleteOpr := hq.db.Delete(&models.Homestay{}, homestayID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete homestay")
	}

	return nil
}

// Insert implements homestay.HomestayRepository.
func (hq *homestayQuery) Insert(homestay models.HomestayEntity) (uint, error) {
	homestayModel := models.HomestayEntityToModel(homestay)

	// Insert homestay data to database
	createOpr := hq.db.Create(&homestayModel)
	if createOpr.Error != nil {
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	// Retrieve the user from the database
	user := models.User{}
	err := hq.db.Model(&user).Where("id = ?", homestayModel.HostID).First(&user).Error
	if err != nil {
		return 0, errors.New("error to retrieve user data")
	}

	// Increment the hosting_count
	user.HostingCount++

	// Update the hosting_count column in the users table
	err = hq.db.Model(&user).Where("id = ?", homestayModel.HostID).UpdateColumn("hosting_count", user.HostingCount).Error
	if err != nil {
		return 0, errors.New("failed to update user hosting count")
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

	var rating sql.NullFloat64
	if err := hq.db.Raw("SELECT AVG(ratings) FROM reviews WHERE homestay_id = ?", homestayID).Scan(&rating).Error; err != nil {
		return models.HomestayEntity{}, err
	}

	averageRating := 0.0
	if rating.Valid {
		averageRating = rating.Float64
	}

	homestay.Rating = averageRating

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
		var rating sql.NullFloat64

		if err := hq.db.Raw("SELECT AVG(ratings) FROM reviews WHERE homestay_id = ?", homestay.ID).Scan(&rating).Error; err != nil {
			return []models.HomestayEntity{}, err
		}

		averageRating := 0.0
		if rating.Valid {
			averageRating = rating.Float64
		}

		homestay.Rating = averageRating
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
