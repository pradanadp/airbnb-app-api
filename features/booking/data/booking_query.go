package data

import (
	models "be-api/features"
	bookingInterface "be-api/features/booking"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type bookingQuery struct {
	db *gorm.DB
}

// SelectAll implements booking.BookingRepository.
func (bq *bookingQuery) SelectAllByID(homestayID uint) ([]models.BookingEntity, error) {
	var bookings []models.Booking
	if err := bq.db.Where("homestay_id = ?", homestayID).Find(&bookings).Error; err != nil {
		return []models.BookingEntity{}, errors.Wrap(err, "failed to retrieve bookings for the homestay")
	}

	var bookingEntities []models.BookingEntity
	for _, booking := range bookings {
		bookingEntities = append(bookingEntities, models.BookingModelToEntity(booking))
	}

	return bookingEntities, nil
}

// Delete implements booking.BookingRepository.
func (bq *bookingQuery) Delete(bookingID uint) error {
	deleteOpr := bq.db.Delete(&models.Booking{}, bookingID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete booking")
	}

	return nil
}

// Insert implements booking.BookingRepository.
func (bq *bookingQuery) Insert(booking models.BookingEntity) (uint, error) {
	bookingModel := models.BookingEntityToModel(booking)

	// Insert the booking and the associated payment
	bookingCreateOpr := bq.db.Create(&bookingModel)
	if bookingCreateOpr.Error != nil {
		return 0, bookingCreateOpr.Error
	}

	payment := models.Payment{
		BookingID: bookingModel.ID,
		Status:    "pending",
	}

	// Insert the payment
	paymentCreateOpr := bq.db.Create(&payment)
	if paymentCreateOpr.Error != nil {
		return 0, paymentCreateOpr.Error
	}

	return bookingModel.ID, nil
}

func New(db *gorm.DB) bookingInterface.BookingRepository {
	return &bookingQuery{
		db: db,
	}
}
