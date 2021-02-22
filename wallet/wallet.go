package wallet

//Wallet represents a users wallet in the system
type Wallet struct {
	id      string
	balance int64 //Range: +/- 9,223,372,036,854,775,807. nine quantillion
}

//GetBalance Fetches Wallet Balnce
func (w *Wallet) GetBalance() int64 {
	return w.balance
}

//Deposit adds Funds to wallet
func (w *Wallet) Deposit(amount int64) bool {
	if amount < 1 {
		return false
	}

	w.balance += amount
	return true
}

//Withdraw Initiates a withdrawal event
func (w *Wallet) Withdraw(amount int64) bool {

	if amount > w.GetBalance() {
		return false
	}

	w.balance -= amount
	return true
}
