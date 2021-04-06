package mpesa

const (
	appKey    = "ZZZZZZZZZZZZZZZZZZZZZZZZZZ" // sandbox --> change to yours
	appSecret = "sissisisisisis"
	shortCode = 1000 // sandbox --> change to yours
)

func SendPaymentRequest(number string, amount int) {
	svc, err := NewPaymentService(appKey, appSecret, PRODUCTION)
	if err != nil {
		panic(err)
	}

	req := STKPush{
		BusinessShortCode: "12",
		Password:          "",
		Timestamp:         "",
		TransactionType:   "",
		Amount:            amount,
		PartyA:            number,
		PartyB:            "",
		PhoneNumber:       "",
		CallBackURL:       "",
		AccountReference:  "",
		TransactionDesc:   "",
	}
	svc.SendSTK(req)
}
