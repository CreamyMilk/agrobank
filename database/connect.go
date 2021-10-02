package database

import (
	"fmt"

	"github.com/CreamyMilk/agrobank/database/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//DB holds global database object
var DB *gorm.DB

// Database settings
const (
	host     = "localhost"
	port     = "3306" // Default port
	user     = "root"
	password = "test_pass"
	dbname   = "AGRODB"
)

// Connect to db
func Connect() error {
	var err error
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	DB, err = gorm.Open(mysql.Open(dbDSN), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func SetupModels(models ...interface{}) {
	err := DB.AutoMigrate(models...)
	if err != nil {
		panic(err)
	}
}

func SeedTransactionCosts() {
	dtbl := []struct {
		upperLimit int64
		charge     int64
	}{
		{100, 1},
		{1000, 5},
		{10000, 10},
		{100000, 15},
		{1000000, 20},
		{10000000, 25},
		{100000000, 100},
		{999999999999, 200},
	}
	for _, t := range dtbl {
		singleCost := models.TransactionCost{
			Upper_limit: t.upperLimit,
			Charge:      t.charge,
		}
		DB.Create(&singleCost)
	}
}

// 	tt := []struct {
// 		name      string
// 		accountid string
// 		balance   int64
// 	}{
// 		{"Normal", "N001", 600},
// 		{"Negative", "Neg1", -100},
// 		{"Decimal Positive", "DP1", 290},
// 		{"Decimal Negative", "DN1", 90909},
// 	}
