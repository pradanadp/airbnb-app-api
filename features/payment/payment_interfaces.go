package payment

import (
	"be-api/features"
	models "be-api/features"
)

type PaymentRepository interface {
	Insert(payment models.ResponMidtrans,id uint) (features.PaymentEntity, error)
	Delete(paymentID uint) error
	Select(UserId uint) (features.PaymentEntity, error)
}

type PaymentService interface {
	CreatePayment(payment models.ResponMidtrans, id uint) (features.PaymentEntity, error)
	DeletePayment(paymentID uint) error
	GetPayment(UserId uint) (features.PaymentEntity, error)
}
