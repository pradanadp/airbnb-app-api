package payment

import (
	"be-api/features"
	models "be-api/features"
)

type PaymentRepository interface {
	Insert(payment models.ResponMidtrans) (features.PaymentEntity, error)
	Delete(paymentID uint) error
}

type PaymentService interface {
	CreatePayment(payment models.ResponMidtrans) (features.PaymentEntity, error)
	DeletePayment(paymentID uint) error
}
