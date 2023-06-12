package service

import (
	models "be-api/features"
	bookingInterface "be-api/features/booking"
	"fmt"
	"time"

	"github.com/pkg/errors"
)

type bookingService struct {
	bookingRepository bookingInterface.BookingRepository
}

// CheckAvailability implements booking.BookingService.
func (bs *bookingService) CheckAvailability(homestayID uint, checkInDate time.Time) (bool, error) {
	bookings, err := bs.bookingRepository.SelectAllByID(homestayID)
	if err != nil {
		return false, err
	}

	for _, booking := range bookings {
		bookingCheckInDate, err := time.Parse("2006-01-02", booking.CheckInDate)
		if err != nil {
			return false, errors.Wrap(err, "failed to parse booking check-in date")
		}

		bookingCheckOutDate, err := time.Parse("2006-01-02", booking.CheckOutDate)
		if err != nil {
			return false, errors.Wrap(err, "failed to parse booking check-out date")
		}

		// Check for overlap
		if checkInDate.After(bookingCheckInDate) && checkInDate.Before(bookingCheckOutDate) {
			return false, errors.New("homestay is not available for the provided check-in date")
		}
	}

	return true, nil
}

// CreateBooking implements booking.BookingService.
func (bs *bookingService) CreateBooking(booking models.BookingEntity) (uint, error) {
	checkInDate, err := time.Parse("2006-01-02", booking.CheckInDate)
	if err != nil {
		return 0, errors.Wrap(err, "failed to parse booking check-in date")
	}
	isAvailable, err := bs.CheckAvailability(booking.HomestayID, checkInDate)

	if !isAvailable {
		if err != nil {
			return 0, err
		}
	}

	switch {
	case booking.CustomerID == 0:
		return 0, errors.New("error, customer id is required")
	case booking.HomestayID == 0:
		return 0, errors.New("error, homestay id is required")
	case booking.CheckInDate == "":
		return 0, errors.New("error, check in date is required")
	case booking.CheckOutDate == "":
		return 0, errors.New("error, check out date is required")
	case booking.Duration == 0:
		return 0, errors.New("error, duration is required")
	case booking.TotalPrice == 0:
		return 0, errors.New("error, total price is required")
	}

	booking.Status = "reserved"
	bookingID, err := bs.bookingRepository.Insert(booking)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return bookingID, nil
}

// DeleteBooking implements booking.BookingService.
func (bs *bookingService) DeleteBooking(bookingID uint) error {
	err := bs.bookingRepository.Delete(bookingID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func New(repo bookingInterface.BookingRepository) bookingInterface.BookingService {
	return &bookingService{
		bookingRepository: repo,
	}
}
