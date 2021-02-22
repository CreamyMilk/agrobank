package wallet

import (
	"testing"
)

func TestBalance(t *testing.T) {
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
			w := Wallet{id: tc.accountid, balance: tc.balance}
			balance := w.GetBalance()
			if balance != tc.balance {
				t.Errorf("Expected Wallet Balance to be %v but got %v", tc.balance, balance)
			}
		})
	}

}

func TestDeposit(t *testing.T) {
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
			//Get Wallet
			w := Wallet{id: tc.accountid, balance: tc.inital}
			//Attempt to deposit
			if w.Deposit(tc.deposit) == tc.possible {
				//Check New Balance
				newBalance := w.GetBalance()
				if newBalance != tc.finalbalance {
					t.Errorf("Inital Balance was %v Deposited Amount %v and expected Balance to be %v not -> %v", tc.inital, tc.deposit, tc.finalbalance, newBalance)
				}
			}
		})
	}
}

func TestWitdrawals(t *testing.T) {
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
			w := Wallet{id: tc.accountid, balance: tc.inital}
			//Attempt to withdraw
			if w.Withdraw(tc.withdrawal) == tc.possible {
				//Check New Balance
				newBalance := w.GetBalance()
				if newBalance != tc.finalbalance {
					t.Errorf("Inital Balance was %v Deposited Amount %v and expected Balance to be %v not -> %v", tc.inital, tc.withdrawal, tc.finalbalance, newBalance)
				}
			}
		})
	}
}
