package wallet

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/go-sql-driver/mysql"
)

//Wallet represents a users wallet in the system
type Wallet struct {
	name    string
	balance int64 //Range: +/- 9,223,372,036,854,775,807. nine quantillion
}

//Create Adds New wallet into db
func (w *Wallet) Create() error {
	_, err := database.DB.Query("INSERT INTO wallets_store (wallet_name,balance) VALUES (?, ?)", w.name, w.balance)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 { //Duplicate errors should be ignored
			return nil
		}
		return (err)
	}
	return nil
}

//Delete Destorys Wallets after testing
func (w *Wallet) Delete() error {
	_, err := database.DB.Query("DELETE FROM wallets_store WHERE wallet_name = ?", w.name)
	if err != nil {
		return err
	}
	return nil
}

//GetBalance Fetches Wallet Balnce
func (w *Wallet) GetBalance() int64 {
	tempBalance := 0
	err := database.DB.QueryRow("SELECT balance FROM wallets_store WHERE wallet_name = ?", w.name).Scan(&tempBalance)
	if err != nil {
		fmt.Printf("Unable to get balance coz of error %v", err)
	}
	w.balance = int64(tempBalance)
	return w.balance
}

//Deposit adds Funds to wallet
func (w *Wallet) Deposit(amount int64) bool {
	newBalance := w.GetBalance() + amount
	if amount < 1 {
		return false
	}
	_, err := database.DB.Query("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newBalance, w.name)
	if err != nil {
		fmt.Printf("-------------->%v", err)
	}
	w.balance = w.GetBalance()
	return true
}

//Withdraw Initiates a withdrawal event
func (w *Wallet) Withdraw(amount int64) bool {

	currentBalance := w.GetBalance()
	if amount <= 0 {
		return false
	}
	if amount > currentBalance {
		return false
	}

	newBalance := currentBalance - amount
	_, err := database.DB.Query("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newBalance, w.name)
	if err != nil {
		fmt.Printf("-------------->%v", err)
	}
	w.balance = w.GetBalance()
	return true
}
