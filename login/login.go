package login

import (
	"errors"
	"fmt"

	"github.com/CreamyMilk/agrobank/database"
	"golang.org/x/crypto/bcrypt"
)

type LoginDetails struct {
	Phonenumber string `json:"phonenumber"`
	Password    string `json:"password"`
	hash        string
}

type UserDetails struct {
	Name          string `json:"fullname"`
	Phonenumber   string `json:"phonenumber"`
	Walletname    string `json:"walletname"`
	WalletBalance string `json:"balance"`
	Role          string `json:"role"`
	Status        int    `json:"status"`
	fname         string
	mname         string
	lname         string
}

func (l LoginDetails) AttemptLogin() (*UserDetails, error) {
	database.DB.QueryRow("SELECT passwordHash FROM user_registration WHERE phonenumber=?", l.Phonenumber).Scan(&l.hash)
	err := bcrypt.CompareHashAndPassword([]byte(l.hash), []byte(l.Password))

	if err != nil {
		return nil, errors.New("username or password seems to be invalid")
	}
	u := new(UserDetails)
	database.DB.QueryRow("SELECT fname,mname,lname,phonenumber,role FROM user_registration WHERE phonenumber=?", l.Phonenumber).Scan(&u.fname, &u.mname, &u.lname, &u.Phonenumber, &u.Role)
	database.DB.QueryRow("SELECT wallet_name,balance FROM wallets_store WHERE wallet_name=?", l.Phonenumber).Scan(&u.Walletname, &u.WalletBalance)
	u.Name = fmt.Sprintf("%s %s %s", u.fname, u.mname, u.lname)
	u.Status = 0
	return u, nil
}
