package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
)

type User struct {
	id                  int
	user_name           string
	mail_address        string
	PASSWORD            string
	create_account_date mysql.NullTime
	level               int
}

func query1(db *sql.DB, id int) {
	result, err := db.Exec("INSERT INTO users (id, user_name, PASSWORD) VALUES (?, 'tekitou', sha2(?, 224))", id, "PASSWORD")
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
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(columns)

	column_types, err := rows.ColumnTypes()
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(column_types); i++ {
		fmt.Println(column_types[i].Name())
		fmt.Println(column_types[i].Length())
		fmt.Println(column_types[i].ScanType().Name())
	}

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
