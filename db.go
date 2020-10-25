package main

import (
	"database/sql"
	"encoding/json"	
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

//DataInfo struct
type DataInfo struct {
	Rma     int    `json:"rma"`
	Sn1     string `json:"sn1"`
	Sn2     string `json:"sn2"`
	Date    string `json:"date"`
	Comment string `json:"comment"`
	Err     string `json:"err"`
}

// Initializes the connection to the SQLite database.
func init() {
	var err error
	db, err = sql.Open("sqlite3", "./db/data")
	if err != nil {
		log.Fatal("Cant connect to db")
	}
}

// Runs the query to add a new record to the database.
func (d DataInfo) addRec() {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO rmaData (RMA, SN1, SN2, DATE, COMMENTS) VALUES (?, ?, ?, ?, ?)")
	stmt.Exec(d.Rma, d.Sn1, d.Sn2, d.Date, d.Comment)
	tx.Commit()
}

// Updates an existing record in the db.
func (d DataInfo) update(p int) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("UPDATE rmaData SET RMA=?, SN1=?, SN2=?, DATE=?, COMMENTS=? WHERE RMA=?")
	stmt.Exec(d.Rma, d.Sn1, d.Sn2, d.Date, d.Comment, p)
	tx.Commit()
}

// Performs a search query.
func (d DataInfo) search(q string) ([]string, string) {
	var results []string
	rows, err := db.Query(q)
	if err != nil {
		d.Err = "Error running query."
		return results, d.Err
	}
	defer rows.Close()
	r, error := returnResults(d, rows)
	if error != "" {
		return results, d.Err
	}
	return r, d.Err
}

// Gets the search results and adds all the individual records to a string slice.
func returnResults(d DataInfo, rows *sql.Rows) ([]string, string) {
	var results []string
	for rows.Next() {
		err := rows.Scan(&d.Rma, &d.Sn1, &d.Sn2, &d.Date, &d.Comment)
		if err != nil {
			d.Err = "Error reading data from query."
			return results, d.Err
		}
		res, _ := json.Marshal(d)
		results = append(results, string(res))
	}
	return results, d.Err
}
