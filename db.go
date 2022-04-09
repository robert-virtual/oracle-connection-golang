package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	go_ora "github.com/sijms/go-ora/v2"
)

var db *sql.DB

func connectDB() {
	databaseUrl := go_ora.BuildUrl("137.184.92.195", 1539, "XE", os.Getenv("DBUSER"), os.Getenv("DBPASS"), nil)
	var openErr error
	db, openErr = sql.Open("oracle", databaseUrl)
	if openErr != nil {
		log.Fatal("openErr:", openErr)
	}
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal("pingErr:", pingErr)
	}
	fmt.Println("conectado a  db...")
}

func closeDB() {
	fmt.Println("Cerrando db...")
	closeErr := db.Close()
	if closeErr != nil {
		log.Fatal("closeErr:", closeErr)
	}
}
