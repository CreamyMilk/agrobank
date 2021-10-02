package wallet

import (
	"github.com/CreamyMilk/agrobank/database"
	"github.com/CreamyMilk/agrobank/database/models"
)

func GetDepositAttemptByCheckId(CheckID string) *models.DepositAttempt {
	var attempt models.DepositAttempt
	database.DB.First(&attempt, "check_id = ? AND proccessed=?", CheckID, false)
	if attempt.CheckID == "" {
		return nil
	}
	return &attempt
}
