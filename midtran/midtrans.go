package midtran

import (

	// "be-api/features"

	"be-api/app/config"
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/midtrans/midtrans-go"
)


func ChargeTransaction(payload []byte) ([]byte,error){
	cfg := config.InitConfig()
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

	response, errRequest := sendRequest(url, "POST", headers, payload)
	if errRequest != nil {	
		return nil, errRequest
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
