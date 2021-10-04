package mpesa

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/CreamyMilk/agrobank/utils"
)

type STKCallbackType string

const (
	DepositTypeSTK      STKCallbackType = "depositType"
	RegistrationTypeSTK STKCallbackType = "registrationType"
)

func getMpesaUrl() string {
	if os.Getenv("MPESA_PRODUCTION") != "" {
		return "https://sandbox.safaricom.co.ke"
	}
	return "https://api.safaricom.co.ke"
}

var (
	baseMpesaURL      = getMpesaUrl()
	defaultApiTimeout = time.Minute
)

type MPesaStkRequest struct {
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
	var (
		appKey    = os.Getenv("MPESA_APP_KEY")
		appSecret = os.Getenv("MPESA_APP_SECRET")
	)
	url := baseMpesaURL + "/oauth/v1/generate?grant_type=client_credentials"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ""
	}
	fmt.Printf("key :%s, secret: %s \n\n", appKey, appSecret)
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

func getCallBackURl(t STKCallbackType) string {
	switch t {
	case DepositTypeSTK:
		return os.Getenv("DEPOSITS_CALLBACK")

	case RegistrationTypeSTK:
		return os.Getenv("REGISTRATIONS_CALLBACK")
	}

	return ""

}
func SendSTK(phonenumber, amount, accountNo, notifToken string, paymentType STKCallbackType) (string, error) {
	var (
		shortCode = os.Getenv("MPESA_SHORT_CODE")
		passKey   = os.Getenv("MPESA_PASS_KEY")
		transType = os.Getenv("MPESA_TYPE")
	)
	fmt.Printf("ShortCode %s Pass Keys : %s \n", shortCode, passKey)
	callbackUrl := getCallBackURl(paymentType)
	transaction := new(MPesaStkRequest)
	sendSTKUrl := baseMpesaURL + "/mpesa/stkpush/v1/processrequest"
	password, timestamp := generatePasswordAndTimeStamp(shortCode, passKey)
	jsonData := map[string]string{
		"BusinessShortCode": shortCode,
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   transType,
		"Amount":            amount,
		"PartyA":            phonenumber,
		"PartyB":            shortCode,
		"PhoneNumber":       utils.ConvertTo254(phonenumber),
		"CallBackURL":       callbackUrl,
		"AccountReference":  accountNo,
		"TransactionDesc":   "detail",
	}
	jsonValue, _ := json.Marshal(jsonData)
	request, _ := http.NewRequest("POST", sendSTKUrl, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+getToken())
	request.Header.Set("cache-control", "no-cache")
	//fmt.Printf("%v", request)
	client := &http.Client{
		Timeout: defaultApiTimeout,
	}

	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var tempSTK map[string]string
	json.Unmarshal([]byte(body), &tempSTK)
	if tempSTK["ResponseCode"] == "0" {
		transaction.CheckoutID = string(tempSTK["CheckoutRequestID"])
		transaction.MerchID = string(tempSTK["MerchantRequestID"])
		transaction.Token = notifToken
		return transaction.CheckoutID, nil
	}
	return "", fmt.Errorf("could not send STK push because %v", tempSTK)
}
