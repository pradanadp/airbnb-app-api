package service

import (
	models "be-api/features"
	paymentInterface "be-api/features/payment"
	"fmt"
)

type paymentService struct {
	paymentRepository paymentInterface.PaymentRepository
}

// CreatePayment implements payment.PaymentService.
func (ps *paymentService) CreatePayment(payment models.ResponMidtrans,BookingID uint) (uint, error) {
	id,err := ps.paymentRepository.Insert(payment,BookingID)
	if err != nil {
		return 0,fmt.Errorf("error: %v", err)
	}

	return id,nil
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
