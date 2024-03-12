package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDb() {

	db, err := sql.Open("mysql", "root:Tasikmalaya123..@tcp(localhost:3306)/perpusdigital")

	if err != nil {
		panic(err.Error())
	}

	DB = db
}
