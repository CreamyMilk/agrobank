package mpesa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	appKey       = "HMHVHRMFqLgCAwVVG2AMcQhIxTEj0CGc" // sandbox --> change to yours
	appSecret    = "3hX4Y98isZvf7mAS"
	shortCode    = "174379" // sandbox --> change to yours
	passKey      = "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"
	baseMpesaURL = "https://sandbox.safaricom.co.ke/"
)

type MPayments struct {
	MerchID    string `json:"id,omitempty" bson:"_id,omitempty"`
	CheckoutID string `json:"CheckoutRequestID"`
	Token      string `json:"token"`
}

func generatePasswordAndTimeStamp(shortCode, passkey string) (string, string) {
	timestamp := time.Now().Format("20060102030405")
	str := fmt.Sprintf("%s%s%s", shortCode, passkey, timestamp)
	return base64.StdEncoding.EncodeToString([]byte(str)), timestamp
}

func getToken() string {
	url := baseMpesaURL + "oauth/v1/generate?grant_type=client_credentials"
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ""
	}
	req.SetBasicAuth(appKey, appSecret)
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")
	req.Header.Add("Connection", "keep-alive")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	if res != nil {
		defer res.Body.Close()
	}
	var authResp authResponse
	err = json.NewDecoder(res.Body).Decode(&authResp)
	if err != nil {
		return ""
	}
	accessToken := authResp.AccessToken
	return accessToken
}

func SendSTK(phonenumber, amount, accountNo, notifToken string) (string, error) {
	transaction := new(MPayments)
	sendSTKUrl := baseMpesaURL + "/mpesa/stkpush/v1/processrequest"
	transType := "CustomerPayBillOnline"
	password, timestamp := generatePasswordAndTimeStamp(shortCode, passKey)
	//fmt.Println(token)
	jsonData := map[string]string{
		"BusinessShortCode": shortCode,
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   transType, //CustomerPayBillOnline
		"Amount":            amount,
		"PartyA":            phonenumber,
		"PartyB":            shortCode,
		"PhoneNumber":       phonenumber,
		"CallBackURL":       "https://68b59ad9a2fa.ngrok.io/stkcall", //Add ourcallback url here,
		"AccountReference":  accountNo,
		"TransactionDesc":   "detail",
	}
	// fmt.Print(jsonData)
	jsonValue, _ := json.Marshal(jsonData)
	//fmt.Println(jsonData)
	request, _ := http.NewRequest("POST", sendSTKUrl, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+getToken())
	request.Header.Set("cache-control", "no-cache")
	//fmt.Printf("%v", request)
	client := &http.Client{}
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)
	var tempSTK map[string]string
	json.Unmarshal([]byte(body), &tempSTK)
	//fmt.Printf("%v", tempSTK)
	if tempSTK["ResponseCode"] == "0" {
		transaction.CheckoutID = string(tempSTK["CheckoutRequestID"])
		transaction.MerchID = string(tempSTK["MerchantRequestID"])
		transaction.Token = notifToken
		return transaction.CheckoutID, nil
	}
	return "", fmt.Errorf("could not send STK push because %v", tempSTK)
}
