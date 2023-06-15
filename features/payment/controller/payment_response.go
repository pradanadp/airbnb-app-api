package controller

import models "be-api/features"

func PaymentMidstransToModel(payment models.ResponMidtrans) models.Payment {
	var booking models.Booking
	return models.Payment{
		BookingID: 	booking.ID,
		Name:   	payment.Bank,
		Status: 	payment.TransactionStatus,
		OrderID:    payment.OrderId,
		VANumber:   payment.VANumber,
	}
}
