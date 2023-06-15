package midtran

// import (
// 	appConfig "be-api/app/config"
//     "github.com/midtrans/midtrans-go"
//     "github.com/midtrans/midtrans-go/coreapi"
// )

// func MidtransTransaction(cfg *appConfig.AppConfig){
// 	// 1. Set you ServerKey with globally
// 	midtrans.ServerKey = cfg.KEY_SERVER_MIDTRANS
// 	authString := encodeAuthString(midtrans.ServerKey, "")
// 	fmt.Println("AUTH_STRING:", authString)
// 	midtrans.Environment = midtrans.Sandbox

// 	// 2. Initiate charge request
// 	chargeReq := &coreapi.ChargeReq{
// 		PaymentType: coreapi.PaymentTypeCreditCard,
// 		TransactionDetails: midtrans.TransactionDetails{
// 			OrderID:  "12345",
// 			GrossAmt: 200000,
// 		},
// 		CreditCard: &coreapi.CreditCardDetails{
// 			TokenID:        "YOUR-CC-TOKEN",
// 			Authentication: true,
// 		},
// 		Items: &[]midtrans.ItemDetails{
// 			{
// 				ID:    "ITEM1",
// 				Price: 200000,
// 				Qty:   1,
// 				Name:  "Someitem",
// 			},
// 		},
// 	}

// 	// 3. Request to Midtrans using global config
// 	coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
// 	fmt.Println("Response :", coreApiRes)
// }

// func encodeAuthString(username, password string) string {
// 	auth := username + ":" + password
// 	authBytes := []byte(auth)
// 	encodedAuth := base64.StdEncoding.EncodeToString(authBytes)
// 	return encodedAuth
// }