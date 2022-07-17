package helloSql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func CreateConnect() {
	if db == nil {
		create()
	}
}

func create() {
	fmt.Println("connect start ....")

	database, err := sqlx.Open("mysql", "root:helloworld123@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}

	fmt.Println(database)
	fmt.Println("connect success ....")

	db = database
}
