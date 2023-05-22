package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func initDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Zedtech0340@@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return db
}

func pullDB(db *sql.DB) {

}

func pushDB(db *sql.DB) {

}
