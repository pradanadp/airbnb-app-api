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
func (ps *paymentService) GetPayment(UserId uint) (features.PaymentEntity,error) {
	data,err:=ps.paymentRepository.Select(UserId)
	if err != nil {
		return features.PaymentEntity{}, fmt.Errorf("error: %v", err)
	}

	return data, nil

}

// CreatePayment implements payment.PaymentService.
func (ps *paymentService) CreatePayment(payment models.ResponMidtrans, id uint) (features.PaymentEntity, error) {
	Orderid, err := ps.paymentRepository.Insert(payment,id)
	if err != nil {
		return features.PaymentEntity{}, fmt.Errorf("error: %v", err)
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
