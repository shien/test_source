package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

type User struct {
	id                  int
	user_name           string
	mail_address        string
	password            string
	create_account_date mysql.NullTime
	level               int
}

func Query1(db *sql.DB, id int) {
	result, err := db.Exec("INSERT INTO users (id, user_name, password) VALUES (?, 'tekitou', sha2(?, 224))", id, "password")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func Query2(db *sql.DB) {

	fmt.Println("Query2")
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

func TestGorm() {
	db, err := gorm.Open("mysql", "root:password@/test_db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}

func main() {
	db, dberr := sql.Open("mysql", "root:password@/test_db")

	if dberr != nil {
		log.Fatal(dberr)
	}
	defer db.Close()

	id := 105

	Query1(db, id)
	Query2(db)

	TestGorm()
}
