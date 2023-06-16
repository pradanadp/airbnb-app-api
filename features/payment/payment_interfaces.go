package payment

import (
	"be-api/features"
	models "be-api/features"
)

type PaymentRepository interface {
	Insert(payment models.ResponMidtrans,booking_id uint) (uint, error)
	Delete(paymentID uint) error
	Select(id uint) (features.PaymentEntity, error)
}

type PaymentService interface {
	CreatePayment(payment models.ResponMidtrans, booking_id uint) (uint, error)
	DeletePayment(paymentID uint) error
	GetPayment(id uint) (features.PaymentEntity, error)
}
