package registration

import (
	"errors"
	"strconv"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/mpesa"
)

//RegistrationLimbo is the general type for first time registrations
type RegistrationLimbo struct {
	databaseID      int64
	Name            string `json:"name"`
	IdNumber        string `json:"idnumber"`
	PhotoUrl        string `json:"photourl"`
	PhoneNumber     string `json:"phone"`
	Email           string `json:"email"`
	FcmToken        string `json:"fcmtoken"`
	InformalAddress string `json:"informaladdress"`
	Xcordinates     string `json:"xcords"`
	Ycordinates     string `json:"ycords"`
	Role            string `json:"role"`
}

func GetTempByID(id string) *RegistrationLimbo {
	return nil
}

func (r *RegistrationLimbo) IsRegisterd() bool {
	return false
}

func (r *RegistrationLimbo) TempCreate() error {
	if r.IsRegisterd() {
		return errors.New("an Account has alreday been opened for your number")
	}
	values := []interface{}{r.IdNumber, r.PhoneNumber, r.FcmToken, "", r.PhotoUrl, r.Email, r.InformalAddress, r.Xcordinates, r.Ycordinates, r.Role}
	res, err := database.DB.Exec("INSERT registration_limbo (idnumber,phonenumber,fcmToken,stkPushid,photo_url,email,informal_address,xCords,yCords,role) VALUES (?,?,?,?,?,?,?,?,?,?)", values...)
	if err != nil {
		return (err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return (err)
	}
	r.databaseID = id
	err = r.sendPayment()
	if err != nil {
		return (err)
	}
	return nil
}

func (r *RegistrationLimbo) InsertPermanent() error {
	values := []interface{}{r.IdNumber, r.PhoneNumber, r.FcmToken, "", r.PhotoUrl, r.Email, r.InformalAddress, r.Xcordinates, r.Ycordinates, r.Role}
	_, err := database.DB.Exec("INSERT user_registration (idnumber,phonenumber,fcmToken,stkPushid,photo_url,email,informal_address,xCords,yCords,role) VALUES (?,?,?,?,?,?,?,?,?,?)", values...)
	if err != nil {
		return (err)
	}
	return nil
}

func (r *RegistrationLimbo) sendPayment() error {
	CheckoutRequestID, err := mpesa.SendSTK(r.PhoneNumber, strconv.Itoa(REGISTRATIONCOST), "JJJ", "ppp")
	if err != nil {
		return (err)
	}
	updatevalues := []interface{}{CheckoutRequestID, r.databaseID}
	_, err = database.DB.Exec("UPDATE registration_limbo SET stkPushid=? WHERE registerID=?", updatevalues...)
	if err != nil {
		return (err)
	}
	return nil
}

func (r *RegistrationLimbo) DeleteTempRegistraion() error {
	_, err := database.DB.Exec("DELETE FROM registration_limbo WHERE registerID = ?", r.databaseID)
	if err != nil {
		return err
	}
	return nil
}
