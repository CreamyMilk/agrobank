package callback

import (
	"fmt"
	"log"

	"github.com/CreamyMilk/agrobank/auth/registration"
	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/firenotifier"
	"github.com/CreamyMilk/agrobank/wallet"
)

type StkpushCallbackResponse struct {
	ID   string `json:"_id"`
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
}

func (resp *StkpushCallbackResponse) ParseRegistrations() {
	if resp.Body.StkCallback.ResultCode == 0 {
		check := resp.Body.StkCallback.CheckoutRequestID
		amount := resp.Body.StkCallback.CallbackMetadata.Item[0].Value.(float64)

		registration.ValidateUser(check, amount)
	}
}
func (resp *StkpushCallbackResponse) ParseDeposits() {
	if resp.Body.StkCallback.ResultCode == 0 {
		check := resp.Body.StkCallback.CheckoutRequestID
		amount := resp.Body.StkCallback.CallbackMetadata.Item[0].Value.(float64)

		//getDeposit
		attempt := wallet.GetDepositAttemptByCheckId(check)
		if attempt == nil {
			go firenotifier.ContactTheDevTeam("Replay Deposits Detected", "Refer to logs for more info")
			return
		}

		err := wallet.DepositToAddress(attempt.WalletAddress, int64(amount))
		if err != nil {
			go firenotifier.ContactTheDevTeam("Deposit Issues", attempt.WalletAddress)
			return
		}

		attempt.Proccessed = true
		database.DB.Save(attempt)
		message := fmt.Sprintf("You wallet has beed debited with Ksh.%s", attempt.Amount)

		go firenotifier.SuccesfulDepoistNotif(message, attempt.WalletAddress, attempt.Amount)
		log.Println("Deposits have been processed well")
	}
}
