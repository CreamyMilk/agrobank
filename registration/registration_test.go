package registration

import (
	"fmt"
	"testing"

	"github.com/CreamyMilk/agrobank/database"
)

func TestNormalRegistration(t *testing.T) {
	if err := database.Connect(); err != nil {
		fmt.Printf("DB ERROR %v", err)
	}
	defer database.DB.Close()

	tt := []struct {
		name            string
		idNumber        string
		photoUrl        string
		phoneNumber     string
		email           string
		fcmToken        string
		informalAddress string
		xcordinates     string
		ycordinates     string
		role            string
	}{
		{"Ideal", "14489829", "phtourl", "254797678252", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "34234322", "phtourl", "254677234534", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "34342555", "phtourl", "254214690431", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "88888856", "phtourl", "254203040202", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "45354345", "phtourl", "254799029029", "email", "fcm", "informal", "xcord", "ycord", "teach"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rl := RegistrationLimbo{
				IdNumber:        tc.idNumber,
				PhoneNumber:     tc.phoneNumber,
				PhotoUrl:        tc.photoUrl,
				Email:           tc.email,
				FcmToken:        tc.fcmToken,
				InformalAddress: tc.informalAddress,
				Xcordinates:     tc.xcordinates,
				Ycordinates:     tc.ycordinates,
				Role:            tc.role}
			err := rl.TempCreate()
			if err != nil {
				t.Errorf("Cannot create account because %v", err)
			}

			err = rl.DeleteTempRegistraion()
			if err != nil {
				t.Errorf("Failed to delete the account because %v", err)
			}

		})
	}
}
