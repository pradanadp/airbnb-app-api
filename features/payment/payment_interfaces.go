package payment

import (
	models "be-api/features"
)

type PaymentRepository interface {
	Insert(payment models.ResponMidtrans,BookingID uint) (uint, error)
	Delete(paymentID uint) error
}

type PaymentService interface {
	CreatePayment(payment models.ResponMidtrans,BookingID uint) (uint, error)
	DeletePayment(paymentID uint) error
}
