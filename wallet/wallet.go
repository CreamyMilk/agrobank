package wallet

import (
	"errors"
	"log"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	errCouldNotGetTransactionCost = errors.New("could not get transaction cost")
	errCouldNotDeleteWallet       = errors.New("could not delete the stated wallet")
	errCouldNotGetWallet          = errors.New("could not get wallet")
	errNegativeDeposit            = errors.New("cant handle negative deposits")
	errDepoistUpdateFailed        = errors.New("could not update the wallet balance after deposit")
	errCouldNotPersitTransaction  = errors.New("could not pesist transaction")
	errSendingToSelf              = errors.New("you cannot send funds to self")
	errTooMuchMoney               = errors.New("the stated amount is way higher than expected")
	errNoTransactionCost          = errors.New("does not have transaction cost")
)

type TransactionsList struct {
	Transactions []models.Transaction `json:"transactions"`
	StatusCode   int                  `json:"status"`
}

func GetWalletByAddress(address string) *models.Wallet {
	var tempWallet models.Wallet
	database.DB.First(&tempWallet, "wallet_address=?", address)
	if tempWallet.WalletAddress == "" {
		return nil
	}
	return &tempWallet
}

func GetWalletByPhone(mobile string) *models.Wallet {
	var tempWallet models.Wallet
	database.DB.First(&tempWallet, "phone_number=?", mobile)
	if tempWallet.WalletAddress == "" {
		return nil
	}
	return &tempWallet
}

func GetWalletByUserID(userid int64) *models.Wallet {
	var tempWallet models.Wallet
	database.DB.First(&tempWallet, "user_id=?", userid)
	if tempWallet.WalletAddress == "" {
		return nil
	}
	return &tempWallet
}
func GetTransactionPrice(amount int64) (int64, error) {
	var tempCost models.TransactionCost
	err := database.DB.First(&tempCost, "upper_limit>=?", amount).Error
	if err != nil {
		return 0, errCouldNotGetTransactionCost
	}
	return tempCost.Charge, nil
}

//CreateWallet Adds New wallet into db
func CreateWallet(Userid uint, mobileNo string, passwordHash string) (*models.Wallet, error) {
	var tempWallet models.Wallet

	tempWallet.WalletAddress = uuid.New().String()
	tempWallet.PhoneNumber = mobileNo
	tempWallet.UserID = Userid
	tempWallet.Balance = 0
	tempWallet.WalletHash = passwordHash
	err := database.DB.Create(&tempWallet).Error

	if err != nil {
		return nil, err
	}
	return &tempWallet, nil
}

//Delete Destorys Wallets after testing
func DeleteWallet(address string, walletid uint) error {
	r := database.DB.Where("wallet_address=? AND id=?", address, walletid).Delete(&models.Wallet{})
	if r.Error != nil {
		return errCouldNotDeleteWallet
	}
	return nil
}

//GetBalance Fetches Wallet Balnce
func GetWalletBalance(address string) int64 {
	var tempWallet models.Wallet
	database.DB.First(&tempWallet, "wallet_address=?", address)
	return tempWallet.Balance
}

func GetWalletTransactions(wallAddress string) (*TransactionsList, error) {
	var list TransactionsList
	err := database.DB.Where("transactions.to = ? OR transactions.from=?", wallAddress, wallAddress).Order("created_at desc").Find(&list.Transactions).Error
	if err != nil {
		list.StatusCode = -1
		return &list, err
	}
	return &list, nil
}

//Withdraw Initiates a withdrawal event
//Todo do this in a transaction explisitly
func WithdrawFromWallet(address string, amount int64) bool {
	w := GetWalletByAddress(address)
	if w == nil {
		return false
	}
	currentBalance := w.Balance
	if amount < 0 {
		return false
	}
	if amount > currentBalance {
		return false
	}
	newBalance := currentBalance - amount

	w.Balance = newBalance
	database.DB.Save(w)

	return true
}

//Deposit is a transactional representaition of old deposit
func DepositToAddress(address string, amount int64) error {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if amount < 1 {
			return errNegativeDeposit
		}
		var currentWallet models.Wallet
		res := tx.First(&currentWallet, "wallet_address=?", address)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return errCouldNotGetWallet
			}
			return res.Error
		}
		newBalance := currentWallet.Balance + amount
		currentWallet.Balance = newBalance
		finalUpdate := tx.Save(&currentWallet)
		if finalUpdate.Error != nil {
			log.Println("Deposits Service is crazy")
			return errDepoistUpdateFailed
		}
		if finalUpdate.RowsAffected == 0 {
			log.Println("Deposits didn't reach wallet")
			return errDepoistUpdateFailed
		}
		var transs models.Transaction
		transs.Amount = amount
		transs.Charge = 0
		transs.From = "Deposit Wallet"
		transs.To = address
		transs.TypeName = models.DepositType
		transs.TrackID = uuidgen()
		persistErr := tx.Create(&transs).Error
		if persistErr != nil {
			log.Println("Transation persist failed")
			log.Println(persistErr)
			return errCouldNotPersitTransaction
		}
		return nil
	})
	if err != nil {
		log.Println("Transaction Creation freaking failed")
		log.Println(err)
		return err
	}
	return nil
}

//SendMoney is used to move money from account a to account b
func SendMoney(amountToSend int64, senderAddress string, recipientAddr string) (bool, error) {

	if amountToSend <= 0 {
		return false, errNegativeDeposit
	}

	if senderAddress == recipientAddr {
		return false, errSendingToSelf
	}
	transactionCost, err := GetTransactionPrice(amountToSend)
	if err != nil {
		return false, err
	}
	if transactionCost == 0 {
		return false, errTooMuchMoney
	}
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		var senderWallet models.Wallet
		res := tx.First(&senderWallet, "wallet_address=?", senderAddress)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return errCouldNotGetWallet
			}
			return res.Error
		}

		var receiverWallet models.Wallet
		res = tx.First(&receiverWallet, "wallet_address=?", recipientAddr)
		if res.Error != nil {
			if errors.Is(res.Error, gorm.ErrRecordNotFound) {
				return errCouldNotGetWallet
			}
			return res.Error
		}

		if amountToSend+transactionCost >= senderWallet.Balance {
			return errNoTransactionCost
		}

		newSenderBalance := senderWallet.Balance - (amountToSend + (transactionCost))
		senderWallet.Balance = newSenderBalance

		finalUpdate := tx.Save(&senderWallet)
		if finalUpdate.Error != nil {
			log.Println("Could not persist sender wallet stuff")
			return errDepoistUpdateFailed
		}
		if finalUpdate.RowsAffected == 0 {
			log.Println("Deposits didn't reach sender wallet")
			return errDepoistUpdateFailed
		}

		newRecipientBalance := receiverWallet.Balance + amountToSend
		receiverWallet.Balance = newRecipientBalance

		rr := tx.Save(&receiverWallet)
		if rr.Error != nil {
			log.Println("Could not persist sender wallet stuff")
			return errDepoistUpdateFailed
		}
		if rr.RowsAffected == 0 {
			log.Println("Deposits didn't reach recipient wallet")
			return errDepoistUpdateFailed
		}

		var transs models.Transaction
		transs.Amount = amountToSend
		transs.Charge = transactionCost
		transs.From = senderWallet.WalletAddress
		transs.To = receiverWallet.WalletAddress
		transs.TypeName = models.SendMoneyType
		transs.TrackID = MakeTransactionCode()
		persistErr := tx.Create(&transs).Error
		if persistErr != nil {
			log.Println("Transation persist failed")
			log.Println(persistErr)
			return errCouldNotPersitTransaction
		}
		return nil
	})

	if err != nil {
		log.Println("didn't catch all the values")
		return false, err
	}
	// 	//Applys Changes to the database
	// 	_, err = notification.SendNotification(w.name, notification.SENDING_MONEY, amountToSend)
	// 	if err != nil {
	// 		fmt.Printf("Failed to send notifcation because %v", err)
	// 	}
	// 	_, err = notification.SendNotification(recipientW.name, notification.RECEVIEING_MONEY, amountToSend)
	// 	if err != nil {
	// 		fmt.Printf("Failed to send notifcation because %v", err)
	// 	}
	return true, nil
}
func MakeTransactionCode() string {
	return uuidgen()
}

// func (w *Wallet) PayEscrow(tx *sql.Tx, productShortName string, productCode string, amountPayable int64) error {
// 	newRecipientBalance := w.GetBalance() - amountPayable
// 	if newRecipientBalance < 0 {
// 		tx.Rollback()
// 		return errors.New("sadly you dont have enough funds to complete the trasaction")
// 	}
// 	_, err := tx.Exec("UPDATE wallets_store SET balance=? WHERE wallet_name=?", newRecipientBalance, w.name)
// 	if err != nil {
// 		tx.Rollback()
// 		return errors.New("sadly we were unable to charge you account kindly try again later")
// 	}
// 	//look into a way to charge to charge users if they place an order
// 	transactionCost := 0
// 	_, err = tx.Exec(`INSERT INTO transactions_list
//  	(transuuid,sender_name,receiver_name,amount,charge,ttype)
//  	VALUES (?,?,?,?,?,?)`, productCode, w.name, productShortName, amountPayable, transactionCost, ESCROW_PAYMENT)
// 	if err != nil {
// 		tx.Rollback()
// 		return err
// 	}
// 	return nil
// }
