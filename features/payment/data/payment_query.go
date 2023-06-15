package data

import (
	"be-api/features"
	models "be-api/features"
	paymentInterface "be-api/features/payment"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strconv"

	"be-api/app/config"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"

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

func (pq *paymentQuery) Insert(payment models.ResponMidtrans) (features.PaymentEntity, error) {

	cfg := config.InitConfig()
	midtrans.ServerKey = cfg.KEY_SERVER_MIDTRANS
	authString := encodeAuthString(midtrans.ServerKey, "")
	fmt.Println("AUTH_STRING:", authString)
	midtrans.Environment = midtrans.Sandbox

	if payment.GrossAmount == "" {
		return features.PaymentEntity{}, errors.New("GrossAmount is empty")
	}

	num, err := strconv.ParseInt(payment.GrossAmount, 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return features.PaymentEntity{},err
	}

	bankTransferReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  payment.OrderId,
			GrossAmt: num,
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: "bca",
		},
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
	errCreate := pq.db.Create(&result)
	if errCreate != nil{
		return features.PaymentEntity{},errCreate.Error
	}

	data :=features.PaymentModelToEntity(result)
	
	return data,nil

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
