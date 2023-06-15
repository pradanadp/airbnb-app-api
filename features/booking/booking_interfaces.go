package booking

import (
	models "be-api/features"
	"time"
)

type BookingRepository interface {
	Insert(booking models.BookingEntity) (uint, string, error)
	SelectAllByID(homestayID uint) ([]models.BookingEntity, error)
	Delete(bookingID uint) error
}

type BookingService interface {
	CreateBooking(booking models.BookingEntity) (uint, string, error)
	CheckAvailability(homestayID uint, checkInDate, checkOutDate time.Time) (bool, error)
	DeleteBooking(bookingID uint) error
}
