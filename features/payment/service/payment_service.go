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

// GetPayment implements payment.PaymentService
func (ps *paymentService) GetPayment(id uint) (features.PaymentEntity,error) {
	data,err:=ps.paymentRepository.Select(id)
	if err != nil {
		return features.PaymentEntity{}, fmt.Errorf("error: %v", err)
	}

	return data, nil

}

// CreatePayment implements payment.PaymentService.
func (ps *paymentService) CreatePayment(payment models.ResponMidtrans, booking_id uint) (uint, error) {
	Orderid, err := ps.paymentRepository.Insert(payment,booking_id)
	if err != nil {
		return 0, fmt.Errorf("error: %v", err)
	}

	return Orderid, nil
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
