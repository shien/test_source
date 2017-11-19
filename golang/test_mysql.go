package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func query1(db *sql.DB, id int) {
	result, err := db.Exec("INSERT INTO users (id, user_name, PASSWORD) VALUES (?, 'tekitou', sha2('PASSWORD', 224))", id)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func query2(db *sql.DB) {

	fmt.Println("query2")
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(columns)
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
