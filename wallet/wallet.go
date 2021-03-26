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
	if w.name == "" {
		return fmt.Errorf("cannot create accouts without id %v", "oya")
	}
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
	_, err := database.DB.Exec("DELETE FROM wallets_store WHERE wallet_name = ?", w.name)
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

//Withdraw Initiates a withdrawal event
func (w *Wallet) Withdraw(amount int64) bool {
	currentBalance := w.GetBalance()
	if amount < 0 {
		return false
	}
	if amount > currentBalance {
		return false
	}

	newBalance := currentBalance - amount
	_, err := database.DB.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newBalance, w.name)
	if err != nil {
		fmt.Printf("-------------->%v", err)
	}
	w.balance = w.GetBalance()
	return true
}

//Deposit is a transactional representaition of old deposit
func (w *Wallet) Deposit(amount int64) bool {
	tempBalance := 0
	tx, err := database.DB.Begin()
	if err != nil {
		tx.Rollback()
		return false
	}

	getBalStm, err := tx.Prepare("SELECT balance FROM wallets_store WHERE wallet_name = ?")
	getBalStm.QueryRow(w.name).Scan(&tempBalance)
	defer getBalStm.Close()
	if err != nil {
		return false
	}
	newBalance := int64(tempBalance) + amount
	if amount < 1 {
		tx.Rollback()
		return false
	}
	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newBalance, w.name)
	if err != nil {
		tx.Rollback()
		fmt.Printf("-------------->%v", err)
	}
	w.balance = int64(tempBalance)
	tx.Commit()
	return true
}

//SendMoney is used to move money from account a to account b
func (w *Wallet) SendMoney(amountToSend int64, recipientW Wallet) bool {
	tx, err := database.DB.Begin()
	if err != nil {
		tx.Rollback()
		return false
	}
	if amountToSend <= 0 {
		tx.Rollback()
		return false
	}
	//get current balance
	err = tx.QueryRow("SELECT balance FROM wallets_store WHERE wallet_name = ?", w.name).Scan(&w.balance)
	if err != nil {
		tx.Rollback()
		return false
	}

	transactionCost := 0
	err = tx.QueryRow("SELECT cost FROM transaction_costs WHERE upper_limit >= ? LIMIT 1", amountToSend).Scan(&transactionCost)
	if err != nil {
		tx.Rollback()
		return false
	}
	//Transaction is over system limit
	if transactionCost == 0 {
		tx.Rollback()
		return false
	}

	if amountToSend+int64(transactionCost) >= w.balance {
		tx.Rollback()
		return false
	}

	if recipientW.name == "" {
		tx.Rollback()
		return false
	}

	newSenderBalance := w.balance - (amountToSend + int64(transactionCost))
	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newSenderBalance, w.name)
	if err != nil {
		tx.Rollback()
		fmt.Printf("-------------->%v", err)
		return false
	}

	err = tx.QueryRow("SELECT balance FROM wallets_store WHERE wallet_name = ?", recipientW.name).Scan(&recipientW.balance)
	if err != nil {
		tx.Rollback()
		return false
	}
	newRecipientBalance := recipientW.balance + amountToSend

	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newRecipientBalance, recipientW.name)
	if err != nil {
		tx.Rollback()
		fmt.Printf("-------------->%v", err)
		return false
	}

	tx.Commit()
	return true
}
