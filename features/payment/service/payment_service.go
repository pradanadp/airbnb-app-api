package service

import (
	"be-api/features"
	models "be-api/features"
	paymentInterface "be-api/features/payment"
	"fmt"
)

type paymentService struct {
	paymentRepository paymentInterface.PaymentRepository
}

// CreatePayment implements payment.PaymentService.
func (ps *paymentService) CreatePayment(payment models.ResponMidtrans) (features.PaymentEntity, error) {
	Orderid,err := ps.paymentRepository.Insert(payment)
	if err != nil {
		return features.PaymentEntity{},fmt.Errorf("error: %v", err)
	}

	return Orderid,nil
}

// DeletePayment implements payment.PaymentService.
func (ps *paymentService) DeletePayment(paymentID uint) error {
	err := ps.paymentRepository.Delete(paymentID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func New(repo paymentInterface.PaymentRepository) paymentInterface.PaymentService {
	return &paymentService{
		paymentRepository: repo,
	}
}
