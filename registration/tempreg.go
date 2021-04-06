package registration

import (
	"github.com/CreamyMilk/agrobank/database"
)

type RegistrationLimbo struct {
	idNumber        string
	photoUrl        string
	phoneNumber     string
	email           string
	fcmToken        string
	informalAddress string
	xcordinates     string
	ycordinates     string
	role            string
}

func (r *RegistrationLimbo) TempCreate() error {
	values := []interface{}{r.idNumber, r.phoneNumber, r.fcmToken, "", r.photoUrl, r.email, r.informalAddress, r.xcordinates, r.ycordinates, r.role}
	_, err := database.DB.Query("INSERT registration_limbo (idnumber,phonenumber,fcmToken,stkPushid,photo_url,email,informal_ddress,xCords,yCords,role) VALUES (?,?,?,?,?,?,?,?,?)", values...)
	if err != nil {
		return (err)
	}
	return nil
}

func (r *RegistrationLimbo) SendPaymentRequest() error {
	//mpesa.SendSTK(r.phoneNumber, REGISTRATIONCOST)
	return nil
}
