package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func query1(db *sql.DB, id int) {

	var result = ""

	err := db.QueryRow("INSERT INTO users (id, user_name) VALUES (?, 'tekitou')", id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result1:" + result)
}

func query2(db *sql.DB) {

	var result = ""

	fmt.Println("query2")
	err := db.QueryRow("select * from users").Scan(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result2:" + result)
}

func main() {
	db, dberr := sql.Open("mysql", "USER:PASSWORD@/DBNAME")
	if dberr != nil {
		log.Fatal(dberr)
	}
	defer db.Close()

	id := 105

	query1(db, id)
	query2(db)
}
