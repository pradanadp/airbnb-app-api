package payment

import (
	models "be-api/features"
)

type PaymentRepository interface {
	Insert(payment models.PaymentEntity) (uint, error)
	Delete(paymentID uint) error
}

type PaymentService interface {
	CreatePayment(payment models.PaymentEntity) (uint, error)
	DeletePayment(paymentID uint) error
}
