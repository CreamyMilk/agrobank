package wallet

import (
	"errors"
	"fmt"
	"testing"

	"github.com/CreamyMilk/agrobank/database"
)

func TestBalance(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}

	tt := []struct {
		name           string
		ownerid        string
		deposoitAmount int64
		expectedErr    error
	}{
		{"Standard Wallet", "N001", 600, nil},
		{"Decimal Negative", "DN1", 90909, nil},
		{"Negative Wallet", "Neg-1", -2920, errNegativeDeposit},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w, e := CreateWallet(1, tc.ownerid, "")
			if e != nil {
				t.Errorf("Cannot create account because %v", e)
				return
			}
			er := DepositToAddress(w.WalletAddress, tc.deposoitAmount)

			if er != nil {
				if !errors.Is(er, tc.expectedErr) {
					t.Errorf("Expected error to be %s but got %s %+v", tc.expectedErr, er.Error(), errors.Is(er, tc.expectedErr))
					return
				}
				DeleteWallet(w.WalletAddress, w.ID)
				return
			}
			newWallet := GetWalletByAddress(w.WalletAddress)
			if newWallet.Balance != tc.deposoitAmount {
				t.Errorf("Expected Wallet Balance to be %v but got %v", tc.deposoitAmount, newWallet.Balance)
			}

			if err := DeleteWallet(w.WalletAddress, w.ID); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", w.WalletAddress, err)
			}
		})
	}
}

func TestWitdrawals(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	tt := []struct {
		name         string
		address      string
		inital       int64
		withdrawal   int64
		possible     bool
		finalbalance int64
	}{
		{"Complete", "A001", 900, 900, true, 0},
		{"Partial Withdrawal", "P001", 1000, 800, true, 200},
		{"Over Withdrawal", "O001", 3000, 10000, false, 3000},
		{"Negative Witdrawal", "N001", 4000, -200, false, 4000},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//Get Wallet
			w, e := CreateWallet(1, tc.name, "")
			if e != nil {
				t.Errorf("Cannot create account because %v", e)
				return
			}
			DepositToAddress(w.WalletAddress, tc.inital)
			//Attempt to withdraw
			if WithdrawFromWallet(w.WalletAddress, tc.withdrawal) == tc.possible {
				//Check New Balance
				newBalance := GetWalletBalance(w.WalletAddress)
				if newBalance != tc.finalbalance {
					t.Errorf("Inital Balance was %v withdrew Amount %v and expected Balance to be %v not -> %v", tc.inital, tc.withdrawal, tc.finalbalance, newBalance)
				}
			}

			if err := DeleteWallet(w.WalletAddress, w.ID); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", w.WalletAddress, err)
			}
		})
	}
}

func TestSendMoney(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	tt := []struct {
		name       string
		wallA      string
		aBalance   int64
		wallB      string
		bBalance   int64
		sendAmount int64
		possible   bool
		finalA     int64
		finalB     int64
	}{
		//{"Ghost Sending", "L001", 1000, "", 0, 10, false, 1000, 0},
		//{"Sending to self", "A001", 100, "A001", 100, 40, false, 100, 100},
		{"SendingZero", "L001", 1000, "F001", 500, 0, false, 1000, 500},
		{"Too RichFor System", "BAZO001", 999999999999999999, "HUST0001", 999, 1000000000000000000, false, 999999999999999999, 999},
		{"Confirm Charges", "HUSTER001", 150, "NDIZI", 5000, 148, false, 150, 5000},
		{"Lacks Transaction Cost", "L001", 1000, "F001", 500, 999, false, 1000, 500},
		{"NegativeAmount", "A002", 1000, "B002", 500, -23, false, 1000, 500},
		{"Sender Has Less", "L001", 1000, "F001", 300, 1500, false, 1000, 300},
		{"Sender Everything", "L001", 1000, "F001", 500, 1000, false, 1000, 500},
		{"Receiver Balance is 0", "C001", 1000, "D001", 0, 500, true, 495, 500},
		{"Ideal Transaction", "J0T1", 7000, "B007", 500, 6300, true, 690, 6800},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			//Get Wallet
			wA, _ := CreateWallet(1, tc.wallA, "")
			wB, _ := CreateWallet(2, tc.wallB, "")
			DepositToAddress(wA.WalletAddress, tc.aBalance)
			DepositToAddress(wB.WalletAddress, tc.bBalance)
			//Attempt to sendMoney
			possible, _ := SendMoney(tc.sendAmount, wA.WalletAddress, wB.WalletAddress)
			if possible == tc.possible {
				walletABal := GetWalletBalance(wA.WalletAddress)
				walletBBal := GetWalletBalance(wB.WalletAddress)
				if walletABal != tc.finalA {
					t.Errorf("Inital Balance of Wallet A was %v he sent Amount %v and expected Balance to be %v not -> %v", tc.aBalance, tc.sendAmount, tc.finalA, walletABal)
				}
				if walletBBal != tc.finalB {
					t.Errorf("Before Receiving amount %v from %v balance was %v but  after transaction expected Balance to be %v not -> %v", tc.sendAmount, tc.wallA, tc.aBalance, tc.finalB, walletBBal)
				}
			} else {
				t.Errorf("The Transaction btwn %s and %s of amount %v seems to have been classified Wrongly", wA.WalletAddress, wB.WalletAddress, tc.sendAmount)
			}

			if err := DeleteWallet(wA.WalletAddress, wA.ID); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", wA.WalletAddress, err)
			}

			if err := DeleteWallet(wB.WalletAddress, wB.ID); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", wB.WalletAddress, err)
			}
		})
	}
}
