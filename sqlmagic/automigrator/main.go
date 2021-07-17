package main

import (
	"database/sql"
	"fmt"

	"io/ioutil"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = "3306" // Default port
	user     = "root"
	password = "test_pass"
	dbname   = "sys"
)

// Connect to db
func connect() error {
	var err error

	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}

func runQuery(query string) error {
	parsedQ := fmt.Sprintf("%s", query)
	_, err := DB.Exec(parsedQ)
  println(parsedQ)
	return err
}

func runAllSqlMigrations(){
  validSqlFiles := []string{
     "wallet.sql",
     "registration.sql",
     "store.sql",
     "machine.sql",
     "escrow.sql",
  }
  for _,file := range validSqlFiles{
    path := fmt.Sprintf("../%s",file)
    content, err := ioutil.ReadFile(path)
    filedata := strings.Split(string(content), ";")
    if err != nil {
      log.Fatal(err)
    }

    for _, data := range filedata[:len(filedata)-1] {
      fmt.Print("+")
      qerr := runQuery(data)
      if qerr != nil {
        println(path)
        println(data)
        log.Fatal(qerr)
      }
    }
  }
}


func main() {
	err := connect()
  if err != nil {
    log.Fatal(err)
  }
	fmt.Println("Running Database Migrations")
  runAllSqlMigrations()
	fmt.Printf("\nDone\n")
}

