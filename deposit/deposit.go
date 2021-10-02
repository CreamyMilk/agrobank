package deposit

// import (
// 	"github.com/CreamyMilk/agrobank/wallet"
// )

// type depositInvoice struct {
// 	DepositID         int64
// 	CheckoutRequestID string
// 	Wallet            wallet.Wallet
// 	Amount            string
// 	// ttype             int
// 	// method            string
// 	// mpesaID           string
// }

// func MakeInvoice(checkID string, recepeintWallet wallet.Wallet, depositAmount string) *depositInvoice {
// 	return &depositInvoice{CheckoutRequestID: checkID, Wallet: recepeintWallet, Amount: depositAmount}
// }

// func GetInvoiceByID(id string) *depositInvoice {
// 	r := depositInvoice{}
// 	walletName := ""
// 	getBalStm, err := database.DB.Prepare("SELECT did,checkoutRequestID,walletname,amount FROM deposit_attempts WHERE  mpesaID IS NULL AND checkoutRequestID = ?")
// 	if err != nil {
// 		return nil
// 	}
// 	err = getBalStm.QueryRow(id).Scan(&r.DepositID, &r.CheckoutRequestID, &walletName, &r.Amount)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}
// 	r.Wallet = *wallet.GetWalletByAddress(walletName)
// 	if err != nil {
// 		fmt.Print(err)
// 		return nil
// 	}
// 	return &r
// }

// //Create Adds New Depoist attempt to the database
// func (i *depositInvoice) Create() error {
// 	if i.CheckoutRequestID == "" {
// 		return fmt.Errorf("cannot create accouts without id %v", "oya")
// 	}
// 	_, err := database.DB.Exec(`INSERT INTO deposit_attempts
// 	(checkoutRequestID,walletname,amount,method)
// 	VALUES(?,?,?,?)`, i.CheckoutRequestID, i.Wallet.WalletName(), i.Amount, "MPESA")
// 	if err != nil {
// 		return (err)
// 	}
// 	return nil
// }

// func (i *depositInvoice) PayOut(receiptName string) error {
// 	d, err := strconv.Atoi(i.Amount)
// 	if err != nil {
// 		return err
// 	}
// 	success := i.Wallet.Deposit(int64(d))
// 	if !success {
// 		return errors.New("could Not deposit funds")
// 	}
// 	_, err = database.DB.Exec("UPDATE deposit_attempts SET mpesaID=? WHERE checkoutRequestID=?", receiptName, i.CheckoutRequestID)
// 	if err != nil {
// 		return err
// 	}
// 	_, err = notification.SendDepositNotifcation(i.Wallet.WalletName(), i.Amount)
// 	if err != nil {
// 		fmt.Printf("Failed to send notifcation because \n%v", err)
// 	}
// 	return nil
// }
