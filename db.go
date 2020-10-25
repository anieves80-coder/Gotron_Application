package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
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

func init() {

	var err error
	db, err = sql.Open("sqlite3", "./db/data")

	if err != nil {
		log.Fatal("Cant connect to db")
	}
}

func (d DataInfo) addUser() {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("INSERT INTO rmaData (RMA, SN1, SN2, DATE, COMMENTS) VALUES (?, ?, ?, ?, ?)")
	_, err := stmt.Exec(d.Rma, d.Sn1, d.Sn2, d.Date, d.Comment)
	fmt.Println(err)
	tx.Commit()
}

func (d DataInfo) update(p int) {		
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("UPDATE rmaData SET RMA=?, SN1=?, SN2=?, DATE=?, COMMENTS=? WHERE RMA=?")
	_, err := stmt.Exec(d.Rma, d.Sn1, d.Sn2, d.Date, d.Comment, p)
	fmt.Println(err)
	tx.Commit()
}

func (d DataInfo) search(q string) []string {
	var results []string
	rows, err := db.Query(q)
	if err != nil {
		d.Err = "Error running query."
		res, _ := json.Marshal(d)
		results = append(results, string(res))
		return results
	}
	defer rows.Close()
	return returnResults(d, rows)
}

func returnResults(d DataInfo, rows *sql.Rows) []string {

	var results []string
	for rows.Next() {		
		err := rows.Scan(&d.Rma, &d.Sn1, &d.Sn2, &d.Date, &d.Comment)
		if err != nil {
			d.Err = "Error reading data from query."
			res, _ := json.Marshal(d)
			results = append(results, string(res))
			return results
		}		
		res, _ := json.Marshal(d)
		results = append(results, string(res))
	}
	return results
}
