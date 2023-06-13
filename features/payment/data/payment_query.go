package data

import (
	models "be-api/features"
	paymentInterface "be-api/features/payment"
	"errors"

	"gorm.io/gorm"
)

type paymentQuery struct {
	db *gorm.DB
}

// Delete implements payment.PaymentRepository.
func (pq *paymentQuery) Delete(paymentID uint) error {
	deleteOpr := pq.db.Delete(&models.Payment{}, paymentID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete payment")
	}

	return nil
}

// Insert implements payment.PaymentRepository.
func (pq *paymentQuery) Insert(payment models.PaymentEntity) (uint, error) {
	paymentModel := models.PaymentEntityToModel(payment)
	paymentCreateOpr := pq.db.Create(&paymentModel)
	if paymentCreateOpr.Error != nil {
		return 0, paymentCreateOpr.Error
	}

	return paymentModel.ID, nil
}

func New(db *gorm.DB) paymentInterface.PaymentRepository {
	return &paymentQuery{
		db: db,
	}
}
