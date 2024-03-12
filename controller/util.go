package controller

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

const (
	username = "root"

	hostname = "127.0.0.1:3306"
	dbname   = "perpusdigital"
)

func ConnectDb() {
	password := "Tasikmalaya123.."
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DB = db

}
