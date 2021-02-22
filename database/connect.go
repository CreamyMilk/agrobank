package database

import (
	"database/sql"
	"fmt"

	//Used to provide postgress driver for database/sql package
	_ "github.com/lib/pq"
)

//DB holds global database object
var DB *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 5432 // Default port
	user     = "postgres"
	password = "password"
	dbname   = "ams_demo"
)

// Connect to db
func Connect() error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
