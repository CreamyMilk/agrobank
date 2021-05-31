package registration

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/mpesa"
	"github.com/CreamyMilk/agrobank/wallet"
	"golang.org/x/crypto/bcrypt"
)

//RegistrationLimbo is the general type for first time registrations
type RegistrationLimbo struct {
	databaseID        int64
	FirstName         string `json:"fname"`
	MiddleName        string `json:"mname"`
	LastName          string `json:"lname"`
	IdNumber          string `json:"idnumber"`
	PhotoUrl          string `json:"photourl"`
	PhoneNumber       string `json:"phone"`
	Email             string `json:"email"`
	FcmToken          string `json:"fcmtoken"`
	Password          string `json:"password"`
	passwordHash      string
	checkoutRequestID string
	InformalAddress   string `json:"informaladdress"`
	Xcordinates       string `json:"xcords"`
	Ycordinates       string `json:"ycords"`
	Role              string `json:"role"`
}

func GetTempByID(id string) *RegistrationLimbo {
	r := RegistrationLimbo{}
	getBalStm, err := database.DB.Prepare("SELECT registerID,idnumber,phonenumber,fname,mname,lname,checkoutRequestID,fcmToken,photo_url,email,passwordHash,informal_address,xCords,yCords,role FROM registration_limbo WHERE checkoutRequestID = ?")
	getBalStm.QueryRow(id).Scan(&r.databaseID, &r.IdNumber, &r.PhoneNumber, &r.FirstName, &r.MiddleName, &r.LastName, &r.checkoutRequestID, &r.FcmToken, &r.PhotoUrl, &r.Email, &r.passwordHash, &r.InformalAddress, &r.Xcordinates, &r.Ycordinates, &r.Role)
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return &r
}

func (r *RegistrationLimbo) IsRegisterd() bool {
	id := 0
	database.DB.QueryRow("SELECT userid FROM user_registration WHERE phonenumber=?", r.PhoneNumber).Scan(&id)
	return id != 0
}

func (r *RegistrationLimbo) TempCreate() (error, string) {
	if r.IsRegisterd() {
		return errors.New("an Account has alreday been opened for your number"), ""
	}
	hash, _ := bcrypt.GenerateFromPassword([]byte(r.Password), 4)
	r.passwordHash = string(hash)
	values := []interface{}{r.IdNumber, r.PhoneNumber, r.FirstName, r.MiddleName, r.LastName, r.FcmToken, "", r.PhotoUrl, r.Email, r.passwordHash, r.InformalAddress, r.Xcordinates, r.Ycordinates, r.Role}
	res, err := database.DB.Exec("INSERT registration_limbo (idnumber,phonenumber,fname,mname,lname,fcmToken,checkoutRequestID,photo_url,email,passwordHash,informal_address,xCords,yCords,role) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)", values...)
	if err != nil {
		return errors.New("400"), ""
	}
	id, err := res.LastInsertId()
	if err != nil {
		return errors.New("402"), ""
	}
	r.databaseID = id
	err, topicName := r.sendPayment()
	if err != nil {
		return errors.New("300"), ""
	}

	return nil, topicName
}

func (r *RegistrationLimbo) InsertPermanent() error {
	values := []interface{}{r.IdNumber, r.PhoneNumber, r.FirstName, r.MiddleName, r.LastName, r.checkoutRequestID, r.PhotoUrl, r.Email, r.passwordHash, r.InformalAddress, r.Xcordinates, r.Ycordinates, r.Role}
	_, err := database.DB.Exec("INSERT user_registration (idnumber,phonenumber,fname,mname,lname,checkoutRequestID,photo_url,email,passwordHash,informal_address,xCords,yCords,role) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)", values...)
	if err != nil {
		return (err)
	}
	w := wallet.MakeWallet(r.PhoneNumber, INITALWALLETDEPOSIT)
	err = w.Create()
	if err != nil {
		return err
	}
	return nil
}

func (r *RegistrationLimbo) sendPayment() (error, string) {
	CheckoutRequestID, err := mpesa.SendSTK(r.PhoneNumber, strconv.Itoa(REGISTRATIONCOST), "JJJ", "ppp")
	r.checkoutRequestID = CheckoutRequestID
	if err != nil {
		return err, ""
	}
	updatevalues := []interface{}{r.checkoutRequestID, r.databaseID}
	_, err = database.DB.Exec("UPDATE registration_limbo SET checkoutRequestID=? WHERE registerID=?", updatevalues...)
	if err != nil {
		return err, ""
	}
	return nil, CheckoutRequestID
}

func (r *RegistrationLimbo) DeleteTempRegistraion() error {
	_, err := database.DB.Exec("DELETE FROM registration_limbo WHERE registerID = ?", r.databaseID)
	if err != nil {
		return err
	}
	return nil
}
