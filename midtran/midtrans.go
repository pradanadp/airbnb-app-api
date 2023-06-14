package midtran

import (
	appConfig "be-api/app/config"
	// "be-api/features"
	paymentInterface "be-api/features/payment"
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/midtrans/midtrans-go"
)

func MitransPayment(cfg *appConfig.AppConfig) ([]byte,error){

	midtrans.ServerKey = cfg.KEY_SERVER_MIDTRANS
	authString := encodeAuthString(midtrans.ServerKey, "")
	fmt.Println("AUTH_STRING:", authString)
	midtrans.Environment = midtrans.Sandbox
	
	url := "https://api.sandbox.midtrans.com/v2/charge"

	headers := map[string]string{
		"Accept":        "application/json",
		"Authorization": authString,
		"Content-Type":  "application/json",
	}
	// var order features.Booking

	// payload := map[string]interface{}{
	// 	"payment_type": "bank_transfer",
	// 	"transaction_details": map[string]interface{}{
	// 		// "order_id": order.OrderID,
	// 		// "gross_amount": order.TotalPrice,
	// 	},
	// 	"bank_transfer": map[string]interface{}{
	// 		"bank": "bca",
	// 	},
	// }

	payload := []byte(`{
		"payment_type": "bank_transfer",
		"transaction_details": {
			"order_id": "order-6171",
			"gross_amount": 44000
		},
		"bank_transfer": {
			"bank": "bca"
		}
	}`)
	
	// payloadJSON, err := json.Marshal(payload)
	// if err != nil {
		
	// 	return nil,err
	// }

	response, err := sendRequest(url, "POST", headers, payload)
	if err != nil {
		
		return nil, err
	}
	return response,nil
}

func sendRequest(url, method string, headers map[string]string, payload []byte) ([]byte, error) {

	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	responseBody := new(bytes.Buffer)
	_, err = responseBody.ReadFrom(res.Body)
	if err != nil {
		return nil, err
	}

	return responseBody.Bytes(), nil
}

func encodeAuthString(username, password string) string {
	auth := username + ":" + password
	authBytes := []byte(auth)
	encodedAuth := base64.StdEncoding.EncodeToString(authBytes)
	return encodedAuth
}

type paymentController struct {
	paymentService paymentInterface.PaymentService
}

func New(service paymentInterface.PaymentService) *paymentController {
	return &paymentController{
		paymentService: service,
	}
}