package login

import (
	"errors"
	"log"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/CreamyMilk/agrobank/wallet"
	"golang.org/x/crypto/bcrypt"
)

var (
	errUserAccountNotFound = errors.New("accounts not found")
	errFailedToLogin       = errors.New("username or password seems to be invalid")
)

type LoginReq struct {
	Phonenumber string `json:"phone"`
	Password    string `json:"password"`
}
type UserDetails struct {
	UserID        uint   `json:"userid"`
	Name          string `json:"fullname"`
	Phonenumber   string `json:"phonenumber"`
	Walletname    string `json:"walletname"`
	WalletBalance int64  `json:"balance"`
	Role          string `json:"role"`
	Status        int    `json:"status"`
}

func (l *LoginReq) AttemptLogin() (*UserDetails, error) {
	var legitUser models.Verifieduser
	err := database.DB.First(&legitUser, "phone_number=?", l.Phonenumber).Error
	if err != nil {
		log.Println(err)
		return nil, errUserAccountNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(legitUser.PasswordHash), []byte(l.Password))
	if err != nil {
		log.Println(err)
		return nil, errFailedToLogin
	}

	wall := wallet.GetWalletByUserID(int64(legitUser.ID))
	if wall == nil {
		log.Println("Could not get wallet details")
		return nil, errFailedToLogin
	}
	u := UserDetails{
		UserID:        legitUser.ID,
		Name:          legitUser.Fullname,
		Phonenumber:   legitUser.PhoneNumber,
		Walletname:    wall.WalletAddress,
		WalletBalance: wall.Balance,
		Role:          legitUser.Role,
		Status:        0,
	}

	return &u, nil
}
