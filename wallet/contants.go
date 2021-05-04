package wallet

import (
	"math/rand"
	"strconv"
	"time"
)

const upperBytes = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowerBytes = "abcdefghijklmnopqrstuvwxyz"
const SENDMONEY_TYPE = 2
const ESCROW_PAYMENT = 3

func uuidgen() string {
	rand.Seed(time.Now().UnixNano())
	uuid := "Tx"
	for ii := 0; ii <= 15; ii += 1 {
		switch ii {
		case 4:
			uuid += "-"
		case 5:
			r := upperBytes[rand.Intn(len(upperBytes))]
			uuid += string(r)
		case 6:
			r := upperBytes[rand.Intn(len(upperBytes))]
			uuid += string(r)
		case 7:
			r := upperBytes[rand.Intn(len(upperBytes))]
			uuid += string(r)
		case 8:
			r := upperBytes[rand.Intn(len(upperBytes))]
			uuid += string(r)
		case 9:
			uuid += "-"
		case 12:
			r := lowerBytes[rand.Intn(len(lowerBytes))]
			uuid += string(r)
		case 14:
			uuid += "-"
		case 15:
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
