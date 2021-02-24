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
					t.Errorf("Inital Balance was %v Withdrew Amount %v and expected Balance to be %v not -> %v", tc.inital, tc.withdrawal, tc.finalbalance, newBalance)
				}
			}

			if err := w.Delete(); err != nil {
				t.Errorf("Cannot delete wallet %s because : %v", w.name, err)
			}
		})
	}
}
