package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	notificationaURL = "http://localhost:8081/notifytopic"
)

type NotificationRequest struct {
	Topic       string `json:"topic"`
	Title       string `json:"title"`
	Extra       string `json:"extra"`
	MessageType int    `json:"type"`
}

func createMessage(walletName string, notificationType int, amount int64) (string, string) {
	var title string
	var extra string

	switch notificationType {
	case SENDING_MONEY:
		title = "Transaction was successful"
		extra = "Click to view receipt"
	case RECEVIEING_MONEY:
		title = "Funds received"
		extra = fmt.Sprintf("You have received Ksh.%v click to view receipt", amount)
	}

	return title, extra
}

func SendNotification(walletName string, typeofnotifcation int, amount int64) (string, error) {
	newnotif := new(NotificationRequest)
	newnotif.MessageType = typeofnotifcation
	newnotif.Topic = walletName
	newnotif.Title, newnotif.Extra = createMessage(walletName, typeofnotifcation, amount)

	jsonValue, _ := json.Marshal(newnotif)
	request, _ := http.NewRequest("POST", notificationaURL, bytes.NewBuffer(jsonValue))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("cache-control", "no-cache")

	client := &http.Client{}
	response, _ := client.Do(request)
	body, _ := ioutil.ReadAll(response.Body)

	var notifResponse map[string]string
	json.Unmarshal([]byte(body), &notifResponse)
	//fmt.Printf("%v", tempSTK)
	if notifResponse["status"] == "0" {

		return "Sent Succesfully", nil
	}
	return "Error has occured", fmt.Errorf("could not send STK push because %v", notifResponse)
}
