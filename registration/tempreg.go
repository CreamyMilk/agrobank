package registration

import (
	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/mpesa"
)

type RegistrationLimbo struct {
	databaseID      int64
	idNumber        string
	photoUrl        string
	phoneNumber     string
	email           string
	fcmToken        string //fcmToken is used to store firebases messaging apitoken
	informalAddress string
	xcordinates     string
	ycordinates     string
	role            string
}

func (r *RegistrationLimbo) TempCreate() error {
	values := []interface{}{r.idNumber, r.phoneNumber, r.fcmToken, "", r.photoUrl, r.email, r.informalAddress, r.xcordinates, r.ycordinates, r.role}
	res, err := database.DB.Exec("INSERT registration_limbo (idnumber,phonenumber,fcmToken,stkPushid,photo_url,email,informal_ddress,xCords,yCords,role) VALUES (?,?,?,?,?,?,?,?,?,?)", values...)
	if err != nil {
		return (err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return (err)
	}
	r.databaseID = id
	return nil
}

func (r *RegistrationLimbo) SendPayment() error {
	mpesa.SendPaymentRequest(r.phoneNumber, REGISTRATIONCOST)
	return nil
}

func (r *RegistrationLimbo) DeleteTempRegistraion() error {
	_, err := database.DB.Exec("DELETE FROM registration_limbo WHERE registerID = ?", r.databaseID)
	if err != nil {
		return err
	}
	return nil
}
