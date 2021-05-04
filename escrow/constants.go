package escrow

import (
	"math/rand"
	"strconv"
	"time"
)

const STANDARD_DELIVERY_DURATION = 48
const upperBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func ReconciliationCodeGen() string {
	rand.Seed(time.Now().UnixNano())
	uuid := "Px"
	for ii := 1; ii <= 14; ii += 1 {
		switch ii {
		case 1, 6, 11:
			uuid += "-"

		case 2, 7, 8, 9, 12:
			r := upperBytes[rand.Intn(len(upperBytes))]
			uuid += string(r)

		default:
			r := strconv.Itoa(rand.Intn(9))
			uuid += r
		}
	}
	//Check uniqueness
	return uuid
}
