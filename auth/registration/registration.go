package registration

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
	"github.com/CreamyMilk/agrobank/firenotifier"
	"github.com/CreamyMilk/agrobank/mpesa"
	"github.com/CreamyMilk/agrobank/wallet"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var (
	errAccountExists       = errors.New("an Account has already been opened for your number")
	errTechnicalProblems   = errors.New("we are currently experiencing some techinical problems")
	errCouldNotSendPayment = errors.New("could Not send payment reqeust")
	//errRegistationAmountErr    = errors.New("the supplied amount is invalid")
	errRegistraionDataNotFound = errors.New("could not find the past registration attempt")
	errAccountCreationFailed   = errors.New("could not create a valid account")
)

const (
	REGISTRATIONCOST    = 50
	INITALWALLETDEPOSIT = 5
	BcryptRounds        = 4
)

//RegistrationLimbo is the general type for first time registrations
type TempRegistrationReq struct {
	Fullname        string `json:"fullname"`
	StoreName       string `json:"storename"`
	IdNumber        string `json:"idnumber"`
	DOB             string `json:"dateofbirth"`
	PhoneNumber     string `json:"phone"`
	Email           string `json:"email"`
	ProfilePhoto    string `json:"profile"`
	IDPhoto         string `json:"idphoto"`
	Pin             string `json:"pin"`
	Role            string `json:"role"`
	FcmToken        string `json:"fcmtoken"`
	InformalAddress string `json:"informaladdress"`
	Xcordinates     string `json:"xcords"`
	Ycordinates     string `json:"ycords"`
}

func (r *TempRegistrationReq) IsNotRegisterd() bool {
	var singleRegistration models.Verifieduser
	result := database.DB.First(&singleRegistration, "phone_number=?", r.PhoneNumber)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func (r *TempRegistrationReq) TempCreate() (string, error) {
	if !r.IsNotRegisterd() {
		return "", errAccountExists
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(r.Pin), BcryptRounds)
	computedHash := string(hash)

	checkID, err := mpesa.SendSTK(r.PhoneNumber, strconv.Itoa(REGISTRATIONCOST), "Register", r.FcmToken, mpesa.DepositTypeSTK)
	if err != nil {
		log.Println(err)
		return "", errCouldNotSendPayment
	}

	newRegistration := &models.Registation{
		DOB:          r.DOB,
		Role:         r.Role,
		Email:        r.Email,
		CheckoutID:   checkID,
		IDPhoto:      r.IDPhoto,
		IDNumber:     r.IdNumber,
		Fullname:     r.Fullname,
		StoreName:    r.StoreName,
		PasswordHash: computedHash,
		PhoneNumber:  r.PhoneNumber,
		ProfilePhoto: r.ProfilePhoto,
	}

	res := database.DB.Create(newRegistration)
	if res.Error != nil {
		log.Println(res.Error)
		return "", errTechnicalProblems
	}

	return checkID, nil
}

func GetUserDetailsByMobile(mobile string) *models.Verifieduser {
	var tempUser models.Verifieduser
	//Add is verified
	database.DB.First(&tempUser, "phone_number=?", mobile)
	if tempUser.ID == 0 {
		return nil
	}
	return &tempUser
}

func ValidateUser(checkoutID string, amount float64) error {
	// if amount != REGISTRATIONCOST {
	// 	log.Println(errRegistationAmountErr)
	// }

	var pastReg models.Registation

	err := database.DB.First(&pastReg, "checkout_id=?", checkoutID).Error
	if err != nil {
		log.Println(err)
		go firenotifier.ContactTheDevTeam("Unrecorgnized Callback", err.Error())
		return errRegistraionDataNotFound
	}
	fmt.Printf("Some fake stuff %+v", pastReg)

	//Create valid Profile Page
	var newUser models.Verifieduser
	newUser.Fullname = pastReg.Fullname
	newUser.StoreName = pastReg.StoreName
	newUser.IDNumber = pastReg.IDNumber
	newUser.DOB = pastReg.DOB
	newUser.PhoneNumber = pastReg.PhoneNumber
	newUser.Email = pastReg.Email
	newUser.IDPhoto = pastReg.IDPhoto
	newUser.ProfilePhoto = pastReg.ProfilePhoto
	newUser.PasswordHash = pastReg.PasswordHash
	newUser.Role = pastReg.Role

	err = database.DB.Create(&newUser).Error
	if err != nil {
		log.Println(err)
		go firenotifier.ContactTheDevTeam("Record Creation Failed", err.Error())
		return errAccountCreationFailed
	}

	//Create wallet
	//Send notification
	_, err = wallet.CreateWallet(newUser.ID, newUser.PhoneNumber, newUser.PasswordHash)
	if err != nil {
		go firenotifier.ContactTheDevTeam("Wallet creation Failed", err.Error())
		return nil
	}

	go firenotifier.SuccessfulRegistrationNotif(checkoutID)
	return nil
}
