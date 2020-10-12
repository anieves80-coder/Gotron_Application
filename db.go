package main

import (
	"database/sql"
	// "fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type dataInfo struct {
	rma     int
	sn1     string
	sn2     string
	date    string
	comment string
	err     string
}

var results []dataInfo

var db *sql.DB

func init() {

	var err error
	db, err = sql.Open("sqlite3", "./db/data")

	if err != nil {
		log.Fatal("Cant connect to db")
	}
}

func (d dataInfo) getAll() []dataInfo {

	rows, err := db.Query("SELECT * FROM rmaData")

	if err != nil {
		d.err = "Error running query."
		results = append(results, d)
		return results
	}

	if rows.Next() {
		err = rows.Scan(&d.rma, &d.sn1, &d.sn2, &d.date, &d.comment)
		if err != nil {
			d.err = "Error reading data from query."
			results = append(results, d)
			return results
		}
		results = append(results, d)
	}

	rows.Close()
	return results

}
