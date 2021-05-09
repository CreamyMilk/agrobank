package wallet

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/notification"
	"github.com/go-sql-driver/mysql"
)

//Wallet represents a users wallet in the system
type Wallet struct {
	name    string
	balance int64 //Range: +/- 9,223,372,036,854,775,807. nine quantillion
}

type Transaction struct {
	TransactionID string `json:"transactionid"`
	From          string `json:"from"`
	FromName      string `json:"fromName"`
	To            string `json:"to"`
	Amount        int64  `json:"amount"`
	Charge        int64  `json:"charge"`
	TypeID        int    `json:"typeid"`
	TypeName      string `json:"typename"`
	Timestamp     int64  `json:"timestamp"`
}

type Transactions struct {
	Transactions []Transaction `json:"transactions"`
	StatusCode   int           `json:"status"`
}

func GetWalletByName(name string) *Wallet {
	tempWall := new(Wallet)
	err := database.DB.QueryRow("SELECT wallet_name,balance FROM wallets_store WHERE wallet_name=? ", name).Scan(&tempWall.name, &tempWall.balance)
	if err != nil {
		return nil
	}
	return tempWall
}

func MakeWallet(name string, amount int64) Wallet {
	return Wallet{name: name, balance: amount}
}

func GetTransactionPrice(amount int64) (int64, error) {
	transactionCost := 0
	err := database.DB.QueryRow("SELECT cost FROM transaction_costs WHERE upper_limit >= ? LIMIT 1", amount).Scan(&transactionCost)
	if err != nil {
		return 0, errors.New("transaction cost coult not be determined")
	}
	return int64(transactionCost), nil
}

func (w *Wallet) WalletName() string {
	return w.name
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

func (w *Wallet) GetTransactions() (*Transactions, error) {
	result := new(Transactions)
	rows, err := database.DB.Query(`
	
	SELECT CONCAT(fname,' ',mname,' ',lname)as senderName,
    transuuid,sender_name,
    receiver_name,amount,
    charge,ttype,
    transactions_type.name as transactionName,
	UNIX_TIMESTAMP(createdAt) as timestamp
	FROM transactions_list 
    INNER JOIN user_registration 
    ON user_registration.phonenumber=sender_name
	LEFT JOIN transactions_type 
    ON transactions_list.ttype = transactions_type.type
    WHERE (sender_name=? OR receiver_name=?) 
    ORDER BY timestamp DESC LIMIT 15
	`, w.name, w.name)

	if err != nil {
		result.StatusCode = -500
		return result, err
	}

	for rows.Next() {
		singleTransaction := Transaction{}
		if err := rows.Scan(
			&singleTransaction.FromName,
			&singleTransaction.TransactionID,
			&singleTransaction.From,
			&singleTransaction.To,
			&singleTransaction.Amount,
			&singleTransaction.Charge,
			&singleTransaction.TypeID,
			&singleTransaction.TypeName,
			&singleTransaction.Timestamp); err != nil {
			result.StatusCode = -501
			return result, err
		}
		result.Transactions = append(result.Transactions, singleTransaction)
	}
	if err != nil {
		result.StatusCode = -502
		return result, err
	}
	if result.Transactions == nil {
		result.StatusCode = -503
	}
	defer rows.Close()
	return result, nil
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
		//fmt.Printf("-------------->%v", err)
	}
	w.balance = int64(tempBalance)
	tx.Commit()
	return true
}

//SendMoney is used to move money from account a to account b
func (w *Wallet) SendMoney(amountToSend int64, recipientW Wallet) (string, bool) {
	tx, err := database.DB.Begin()
	errorMessage := ""
	if err != nil {
		tx.Rollback()
		return errorMessage, false
	}
	if amountToSend <= 0 {
		errorMessage = "You cannot send negative values."
		tx.Rollback()
		return errorMessage, false
	}
	//You cannot send to self
	if w.name == recipientW.name {
		errorMessage = "Cannot send funds to self"
		tx.Rollback()
		return errorMessage, false
	}
	//get current balance
	err = tx.QueryRow("SELECT balance FROM wallets_store WHERE wallet_name = ?", w.name).Scan(&w.balance)
	if err != nil {
		errorMessage = "Could not retrieve your balance."
		tx.Rollback()
		return errorMessage, false
	}

	transactionCost := 0
	err = tx.QueryRow("SELECT cost FROM transaction_costs WHERE upper_limit >= ? LIMIT 1", amountToSend).Scan(&transactionCost)
	if err != nil {
		errorMessage = "Could not retrieve your balance."
		tx.Rollback()
		return errorMessage, false
	}
	//Transaction is over system limit
	if transactionCost == 0 {
		errorMessage = "Too much funds requested."
		tx.Rollback()
		return errorMessage, false
	}

	if amountToSend+int64(transactionCost) >= w.balance {
		errorMessage = "Does not have transaction cost."
		tx.Rollback()
		return errorMessage, false
	}

	if recipientW.name == "" {
		errorMessage = "The Receipient seems to be invalid."
		tx.Rollback()
		return errorMessage, false
	}

	newSenderBalance := w.balance - (amountToSend + int64(transactionCost))
	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newSenderBalance, w.name)
	if err != nil {
		errorMessage = "No update to sender balance"
		tx.Rollback()
		return errorMessage, false
	}

	err = tx.QueryRow("SELECT balance FROM wallets_store WHERE wallet_name = ?", recipientW.name).Scan(&recipientW.balance)
	if err != nil {
		errorMessage = "Could not get current balance"
		tx.Rollback()
		return errorMessage, false
	}
	newRecipientBalance := recipientW.balance + amountToSend

	_, err = tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newRecipientBalance, recipientW.name)
	if err != nil {
		errorMessage = "No update to receiver balance"
		tx.Rollback()
		return errorMessage, false
	}

	//generate transaction of what had occured
	genereatedId := MakeTransactionCode()
	_, err = tx.Exec("INSERT INTO transactions_list (transuuid,sender_name,receiver_name,amount,charge,ttype) VALUES (?,?,?,?,?,?)", genereatedId, w.name, recipientW.name, amountToSend, transactionCost, SENDMONEY_TYPE)
	if err != nil {
		errorMessage = ""
		tx.Rollback()
		return errorMessage, false
	}
	//Applys Changes to the database
	tx.Commit()

	_, err = notification.SendNotification(w.name, notification.SENDING_MONEY, amountToSend)
	if err != nil {
		fmt.Printf("Failed to send notifcation because %v", err)
	}
	_, err = notification.SendNotification(recipientW.name, notification.RECEVIEING_MONEY, amountToSend)
	if err != nil {
		fmt.Printf("Failed to send notifcation because %v", err)
	}
	return "", true
}

func MakeTransactionCode() string {
	//Add check to ensure generated uuuids here
	return uuidgen()
}

func (w *Wallet) PayEscrow(tx *sql.Tx, productShortName string, productCode string, amountPayable int64) error {

	newRecipientBalance := w.GetBalance() - amountPayable

	if newRecipientBalance < 0 {
		tx.Rollback()
		return errors.New("sadly you dont have enough funds to complete the trasaction")
	}
	_, err := tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newRecipientBalance, w.name)
	if err != nil {
		tx.Rollback()
		return errors.New("sadly we were unable to charge you account kindly try again later")
	}
	//look into a way to charge to charge users if they place an order
	transactionCost := 0
	_, err = tx.Exec(`INSERT INTO transactions_list 
	(transuuid,sender_name,receiver_name,amount,charge,ttype) 
	VALUES (?,?,?,?,?,?)`, productCode, w.name, productShortName, amountPayable, transactionCost, ESCROW_PAYMENT)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
