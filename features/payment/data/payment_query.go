package data

import (
	"be-api/features"
	models "be-api/features"
	paymentInterface "be-api/features/payment"
	"encoding/base64"
	"errors"
	"fmt"
	"log"

	"be-api/app/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"

	"gorm.io/gorm"
)

type paymentQuery struct {
	db *gorm.DB
}

// Select implements payment.PaymentRepository
func (pq *paymentQuery) Select(UserId uint) (features.PaymentEntity, error) {
	var booking models.Booking

	queryBooking := pq.db.Where("customer_id = ?", UserId).First(&booking)
	if queryBooking.Error != nil {
		return features.PaymentEntity{}, queryBooking.Error
	}
	
	var payment models.Payment
	queryPayment := pq.db.Preload("Bookings").Where("booking_id = ?", booking.ID).First(&payment)
	if queryPayment.Error != nil {
		return features.PaymentEntity{}, queryPayment.Error
	}

	data := features.PaymentModelToEntity(payment)
	return data, nil
	
}

// Delete implements payment.PaymentRepository.
func (pq *paymentQuery) Delete(paymentID uint) error {
	deleteOpr := pq.db.Delete(&models.Payment{}, paymentID)
	if deleteOpr.Error != nil {
		return errors.New(deleteOpr.Error.Error() + ", failed to delete payment")
	}

	return nil
}

func (pq *paymentQuery) Insert(payment models.ResponMidtrans,UserId uint,booking_id uint) (features.PaymentEntity, error) {
	
    var booking models.Booking
    query := pq.db.Where("id = ?",booking_id).Last(&booking)
    if query.Error != nil {
        return models.PaymentEntity{}, query.Error
    }


	cfg := config.InitConfig()
	midtrans.ServerKey = cfg.KEY_SERVER_MIDTRANS
	authString := encodeAuthString(midtrans.ServerKey, "")
	fmt.Println("AUTH_STRING:", authString)
	midtrans.Environment = midtrans.Sandbox

	bankTransferReq := &coreapi.ChargeReq{
		PaymentType:        coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{OrderID: booking.OrderID, GrossAmt: int64(booking.TotalPrice)},
		BankTransfer:       &coreapi.BankTransferDetails{Bank: "bca"},
		Metadata:           nil,
	}

	coreApiRes, errCore := coreapi.ChargeTransaction(bankTransferReq)
	if errCore != nil {
		log.Fatal("Failed to charge transaction:", errCore)
	}
	fmt.Println("Response :", coreApiRes)

	var result features.Payment
	result.Name = coreApiRes.Bank
	result.OrderID = coreApiRes.OrderID
	result.Status = coreApiRes.TransactionStatus
	result.VANumber = coreApiRes.VaNumbers[0].VANumber

	result.BookingID = booking.ID 
	result.Name ="bca"
	errCreate := pq.db.Create(&result)
	if errCreate != nil {
		return features.PaymentEntity{}, errCreate.Error
	}

	data := features.PaymentModelToEntity(result)

	return data, nil

}

func encodeAuthString(username, password string) string {
	auth := username + ":" + password
	authBytes := []byte(auth)
	encodedAuth := base64.StdEncoding.EncodeToString(authBytes)
	return encodedAuth
}

func New(db *gorm.DB) paymentInterface.PaymentRepository {
	return &paymentQuery{
		db: db,
	}
}
