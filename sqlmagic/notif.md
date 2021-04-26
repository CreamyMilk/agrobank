logging to show us whats going on. The main.go file now looks like this:

package main

import (
	"log"
	"time"
)

const iCalDateFormat = "20060102"

var db *pgDb

func main() {
	log.Println("Connecting to database")
	var err error
	db, err = initDb()
	if err != nil {
		log.Fatalf("Error initializing database: %v\n", err)
	}

	err = db.createTablesIfNotExist()
	if err != nil {
		log.Fatalf("Error creating database tables: %v\n", err)
	}

	log.Println("Starting event update goroutine")
	updateTicker := time.NewTicker(time.Second * 5)
	go func() {
		for range updateTicker.C {
			log.Println("Updating events")
			updateEvents()
		}
	}()

	log.Println("Starting http server")
	serveHttp()
}
