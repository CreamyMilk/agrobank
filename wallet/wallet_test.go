package wallet

import (
	"fmt"
	"testing"

	"github.com/CreamyMilk/agrobank/database"
)

func TestBalance(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	defer database.DB.Close()

	tt := []struct {
		name      string
		accountid string
		balance   int64
	}{
		{"Normal", "N001", 600},
		{"Negative", "Neg1", -100},
		{"Decimal Positive", "DP1", 290},
		{"Decimal Negative", "DN1", 90909},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := Wallet{name: tc.accountid, balance: tc.balance}
			e := w.Create()
			if e != nil {
				t.Errorf("Cannot create account because %v", e)
			}
			balance := w.GetBalance()
			if balance != tc.balance {
				t.Errorf("Expected Wallet Balance to be %v but got %v", tc.balance, balance)
			}
			if err := w.Delete(); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", w.name, err)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	defer database.DB.Close()
	tt := []struct {
		name         string
		accountid    string
		inital       int64
		deposit      int64
		possible     bool
		finalbalance int64
	}{
		{"Deposit", "A001", 100, 900, true, 1000},
		{"New Account", "New01", 0, 2000, true, 2000},
		{"Negative Deposits", "Negative", 3000, -234, false, 3000},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			w := Wallet{name: tc.accountid, balance: tc.inital}
			e := w.Create()
			if e != nil {
				t.Errorf("Cannot create account because %v", e)
			}
			//Attempt to deposit
			if w.Deposit(tc.deposit) == tc.possible {
				//Check New Balance
				newBalance := w.GetBalance()
				if newBalance != tc.finalbalance {
					t.Errorf("Inital Balance was %v Deposited Amount %v and expected Balance to be %v not -> %v", tc.inital, tc.deposit, tc.finalbalance, newBalance)
				}
			}
			if err := w.Delete(); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", w.name, err)
			}
		})
	}
}

func TestWitdrawals(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	defer database.DB.Close()
	tt := []struct {
		name         string
		accountid    string
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
			w := Wallet{name: tc.accountid, balance: tc.inital}
			w.Create()
			//Attempt to withdraw
			if w.Withdraw(tc.withdrawal) == tc.possible {
				//Check New Balance
				newBalance := w.GetBalance()
				if newBalance != tc.finalbalance {
					t.Errorf("Inital Balance was %v ite Amount %v and expected Balance to be %v not -> %v", tc.inital, tc.withdrawal, tc.finalbalance, newBalance)
				}
			}

			if err := w.Delete(); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", w.name, err)
			}
		})
	}
}

func TestSendMoney(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	defer database.DB.Close()
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
		{"SendingZero", "L001", 1000, "F001", 500, 0, false, 1000, 500},
		{"Too RichFor System", "BAZO001", 999999999999999999, "HUST0001", 999, 1000000000000000000, false, 999999999999999999, 999},
		{"Sending to self", "A001", 100, "A001", 100, 40, false, 100, 100},
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
			wA := Wallet{name: tc.wallA, balance: tc.aBalance}
			wB := Wallet{name: tc.wallB, balance: tc.bBalance}
			wA.Create()
			wB.Create()
			//Attempt to sendMoney
			_, possible := wA.SendMoney(tc.sendAmount, wB)
			if possible == tc.possible {
				//Check New Balance
				walletABal := wA.GetBalance()
				walletBBal := wB.GetBalance()
				if walletABal != tc.finalA {
					t.Errorf("Inital Balance of Wallet A was %v he sent Amount %v and expected Balance to be %v not -> %v", tc.aBalance, tc.sendAmount, tc.finalA, walletABal)
				}
				if walletBBal != tc.finalB {
					t.Errorf("Before Receiving amount %v from %v balance was %v but  after transaction expected Balance to be %v not -> %v", tc.sendAmount, tc.wallA, tc.aBalance, tc.finalB, walletBBal)
				}
			} else {
				t.Errorf("The Transaction btwn %s and %s of amount %v seems to have been classified Wrongly", wA.name, wB.name, tc.sendAmount)
			}

			if err := wA.Delete(); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", wA.name, err)
			}
			if err := wB.Delete(); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", wB.name, err)
			}
		})
	}
}
