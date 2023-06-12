package controller

import (
	models "be-api/features"
)

type BookingResponse struct {
	CustomerID   uint    `json:"customer_id,omitempty"`
	HomestayID   uint    `json:"homestay_id,omitempty"`
	CheckInDate  string  `json:"check_in_date,omitempty"`
	CheckOutDate string  `json:"check_out_date,omitempty"`
	Status       string  `json:"booking_status,omitempty"`
	Duration     uint    `json:"duration,omitempty"`
	TotalPrice   float64 `json:"total_price,omitempty"`
}

func BookingEntityToResponse(booking models.BookingEntity) BookingResponse {
	return BookingResponse{
		CustomerID:   booking.CustomerID,
		HomestayID:   booking.HomestayID,
		CheckInDate:  booking.CheckInDate,
		CheckOutDate: booking.CheckOutDate,
		Status:       booking.Status,
		Duration:     booking.Duration,
		TotalPrice:   booking.TotalPrice,
	}
}
