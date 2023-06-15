package data

import (
	"be-api/features"
	models "be-api/features"
	paymentInterface "be-api/features/payment"
	"be-api/features/payment/controller"
	"be-api/midtran"
	"encoding/json"
	"errors"
	"fmt"
	"log"

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
func (pq *paymentQuery) Insert(payment models.ResponMidtrans, BookingID uint) (uint, error) {

	var order features.Booking

	errFound :=pq.db.First(order, BookingID)
	if errFound != nil{
		return 0,errFound.Error
	}
	payload := map[string]interface{}{
		"payment_type": "bank_transfer",
		"transaction_details": map[string]interface{}{
			"order_id": order.OrderID,
			"gross_amount": order.TotalPrice,
		},
		"bank_transfer": map[string]interface{}{
			"bank": "bca",
		},
	}
	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		
		return 0,err
	}

    response, errMidtrans := midtran.ChargeTransaction(payloadJSON)
    if errMidtrans != nil {
        return 0, errMidtrans
    }
    fmt.Println("response:", response)

    var midtransResp models.ResponMidtrans
    errMarshal := json.Unmarshal(response, &midtransResp)
    if errMarshal != nil {
		log.Printf("Error sending request: %s", errMarshal.Error())
        return 0, errMarshal
    }
	fmt.Println("responseMashal:", midtransResp)

    midtrans := controller.PaymentMidstransToModel(midtransResp)
	midtrans.BookingID = 6
    errCreate := pq.db.Create(&midtrans)
    if errCreate != nil {
        return 0, errCreate.Error
    }

	fmt.Println("data status:",midtrans.Status)
	fmt.Println("data order:",midtrans.OrderID)
	fmt.Println("data kartu:",midtrans.VANumber)

    dataRespons := features.PaymentModelToEntity(midtrans)
    fmt.Println("dataRespons.ID:", dataRespons.ID)
    return dataRespons.ID, nil

}

func New(db *gorm.DB) paymentInterface.PaymentRepository {
	return &paymentQuery{
		db: db,
	}
}
