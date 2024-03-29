package firenotifier

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/CreamyMilk/agrobank/database/models"
	"google.golang.org/api/option"
)

const (
	fcmConfigarationPath = "bb.json"
)

var FCMMessanger *messaging.Client

func Init() {
	opt := option.WithCredentialsFile(fcmConfigarationPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	FCMMessanger, err = app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("error initializing meessages: %v\n", err)
	}
	log.Println("Notifications initalized 🎉")

}

func SuccessfulRegistrationNotif(checkoutid string) {
	log.Printf("Sending message to %s \n ", checkoutid)
	title := "🎉 Account Verified"
	body := "You can login to access your account"
	notifType := "registration"
	messo := &messaging.Message{
		Topic: checkoutid,
		Data: map[string]string{
			"type": notifType,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:                 title,
				Body:                  body,
				Color:                 "#ffffff",
				Priority:              messaging.PriorityMax,
				ChannelID:             "channelid",
				Icon:                  "ic_stat_sports_volleyball",
				VibrateTimingMillis:   []int64{100, 50, 100},
				DefaultVibrateTimings: false,
				//		ClickAction:           "FLUTTER_NOTIFICATION_CLICK", //Makes it to be open app or not
				ImageURL: "https://cdn.dribbble.com/users/414474/screenshots/16220082/media/3ae262821ac9096f55baca9d60a2f065.png?compress=1&resize=800x600",
			},
			CollapseKey: "ck",
			Data: map[string]string{
				"type": notifType,
			},
			Priority: "high",
		},
	}
	result, err := FCMMessanger.Send(context.Background(), messo)

	if err != nil {
		log.Fatalf("sending the meesage kinda failed %s", err.Error())
	}
	fmt.Printf("%s", result)
}

func ContactTheDevTeam(header, message string) {
	log.Printf("Sending Dev Team message to %s \n ", message)
	title := "🎉 Dev Team : " + header
	messo := &messaging.Message{
		Topic: "dev",
		Data: map[string]string{
			"type": "error",
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  message,
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:                 title,
				Body:                  message,
				Color:                 "#ffffff",
				Priority:              messaging.PriorityMax,
				ChannelID:             "channelid",
				Icon:                  "ic_stat_sports_volleyball",
				VibrateTimingMillis:   []int64{100, 50, 100},
				DefaultVibrateTimings: false,
				//		ClickAction:           "FLUTTER_NOTIFICATION_CLICK", //Makes it to be open app or not
				ImageURL: "https://i.pinimg.com/736x/10/7a/97/107a97ca5bd4a571edcebec54a66fc32.jpg",
			},
			CollapseKey: "ck",
			Data: map[string]string{
				"type": "dev",
			},
			Priority: "high",
		},
	}
	result, err := FCMMessanger.Send(context.Background(), messo)

	if err != nil {
		log.Fatalf("sending the meesage kinda failed %s", err.Error())
	}
	fmt.Printf("%s", result)
}

func SuccesfulDepoistNotif(message, walletAddress, amount string) {
	log.Printf("Sending Deposit message to %s \n ", message)
	title := "🎉 Depoist is Successful"
	notifType := "deposit"

	messo := &messaging.Message{
		Topic: walletAddress,
		Data: map[string]string{
			"type":   notifType,
			"amount": amount,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  message,
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:                 title,
				Body:                  message,
				Color:                 "#ffffff",
				Priority:              messaging.PriorityMax,
				ChannelID:             "channelid",
				Icon:                  "ic_stat_sports_volleyball",
				VibrateTimingMillis:   []int64{100, 50, 100},
				DefaultVibrateTimings: false,
				//		ClickAction:           "FLUTTER_NOTIFICATION_CLICK", //Makes it to be open app or not
				//ImageURL: "https://i.pinimg.com/736x/10/7a/97/107a97ca5bd4a571edcebec54a66fc32.jpg",
			},
			CollapseKey: "ck",
			Data: map[string]string{
				"type": notifType,
			},
			Priority: "high",
		},
	}
	result, err := FCMMessanger.Send(context.Background(), messo)

	if err != nil {
		log.Printf("sending the meesage kinda failed %s", err.Error())
	}
	fmt.Printf("%s", result)
}

func SuccesfulPurchaseNotif(p models.Product, sellerAddr, orderID string) {
	log.Printf("Sending Purchase Notif message to %s \n ", sellerAddr)
	title := fmt.Sprintf("🎉 An Order for '%s' has beed Placed", p.ProductName)
	message := "Click to view more details"
	notifType := "orderplace"
	messo := &messaging.Message{
		Topic: sellerAddr,
		Data: map[string]string{
			"type":     notifType,
			"orderid":  orderID,
			"prodname": p.ProductName,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  message,
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:                 title,
				Body:                  message,
				Color:                 "#ffffff",
				Priority:              messaging.PriorityMax,
				ChannelID:             "channelid",
				Icon:                  "ic_stat_sports_volleyball",
				VibrateTimingMillis:   []int64{100, 50, 100},
				DefaultVibrateTimings: false,
				//		ClickAction:           "FLUTTER_NOTIFICATION_CLICK", //Makes it to be open app or not
				//ImageURL: "https://i.pinimg.com/736x/10/7a/97/107a97ca5bd4a571edcebec54a66fc32.jpg",
			},
			CollapseKey: "ck",
			Data: map[string]string{
				"type":     notifType,
				"orderid":  orderID,
				"prodname": p.ProductName,
			},
			Priority: "high",
		},
	}
	result, err := FCMMessanger.Send(context.Background(), messo)

	if err != nil {
		log.Printf("sending the meesage kinda failed %s", err.Error())
	}
	fmt.Printf("%s", result)
}

func BuyerPurchaseComplete(buyerWalletAddr string) {
	log.Printf("Sending message to %s \n ", buyerWalletAddr)
	title := "🎉 Your order has been completed"
	body := "If not kindly contact support"
	notifType := "ordercomplete"
	messo := &messaging.Message{
		Topic: buyerWalletAddr,
		Data: map[string]string{
			"type": notifType,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:                 title,
				Body:                  body,
				Color:                 "#ffffff",
				Priority:              messaging.PriorityMax,
				ChannelID:             "channelid",
				Icon:                  "ic_stat_sports_volleyball",
				VibrateTimingMillis:   []int64{100, 50, 100},
				DefaultVibrateTimings: false,
				//		ClickAction:           "FLUTTER_NOTIFICATION_CLICK", //Makes it to be open app or not
				//ImageURL: "https://cdn.dribbble.com/users/414474/screenshots/16220082/media/3ae262821ac9096f55baca9d60a2f065.png?compress=1&resize=800x600",
			},
			CollapseKey: "ck",
			Data: map[string]string{
				"type": notifType,
			},
			Priority: "high",
		},
	}
	result, err := FCMMessanger.Send(context.Background(), messo)

	if err != nil {
		log.Fatalf("sending the meesage kinda failed %s", err.Error())
	}
	fmt.Printf("%s", result)
}

func SellerInvoiceComplete(sellerAddr string) {
	log.Printf("Sending message to %s \n ", sellerAddr)
	title := "🎉 Your have received you fund to the order"
	body := "If not kindly contact support"
	notifType := "sellerwin"
	messo := &messaging.Message{
		Topic: sellerAddr,
		Data: map[string]string{
			"type": notifType,
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:                 title,
				Body:                  body,
				Color:                 "#ffffff",
				Priority:              messaging.PriorityMax,
				ChannelID:             "channelid",
				Icon:                  "ic_stat_sports_volleyball",
				VibrateTimingMillis:   []int64{100, 50, 100},
				DefaultVibrateTimings: false,
				//		ClickAction:           "FLUTTER_NOTIFICATION_CLICK", //Makes it to be open app or not
				//ImageURL: "https://cdn.dribbble.com/users/414474/screenshots/16220082/media/3ae262821ac9096f55baca9d60a2f065.png?compress=1&resize=800x600",
			},
			CollapseKey: "ck",
			Data: map[string]string{
				"type": notifType,
			},
			Priority: "high",
		},
	}
	result, err := FCMMessanger.Send(context.Background(), messo)

	if err != nil {
		log.Fatalf("sending the meesage kinda failed %s", err.Error())
	}
	fmt.Printf("%s", result)
}
