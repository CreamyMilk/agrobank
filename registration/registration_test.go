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
		{"Normal", "14489829", "phtourl", "photonumber", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "34234322", "phtourl", "photonumber", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "34342555", "phtourl", "photonumber", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "88888856", "phtourl", "photonumber", "email", "fcm", "informal", "xcord", "ycord", "teach"},
		{"Normal", "45354345", "phtourl", "photonumber", "email", "fcm", "informal", "xcord", "ycord", "teach"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			rl := RegistrationLimbo{
				idNumber:        tc.idNumber,
				phoneNumber:     tc.phoneNumber,
				photoUrl:        tc.photoUrl,
				email:           tc.email,
				fcmToken:        tc.fcmToken,
				informalAddress: tc.informalAddress,
				xcordinates:     tc.xcordinates,
				ycordinates:     tc.ycordinates,
				role:            tc.role}
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
