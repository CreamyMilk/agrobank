package database

import (
	"database/sql"
	"fmt"

	//Used to provide mysql driver for database/sql package
	_ "github.com/go-sql-driver/mysql"
)

//DB holds global database object
var DB *sql.DB

// Database settings
const (
	host     = "localhost"
  port     = "3306" // Default port
	user     = "root"
	password = "test_pass"
	dbname   = "agrodb"
)

// Connect to db
func Connect() error {
	var err error

  DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host,port,dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
