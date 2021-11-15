package util

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// GetConnection get a connection to the mysql database
func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", "myuser:password@tcp(mysql:3306)/QR_CODE_SYSTEM")
	if err != nil {
		panic(err.Error())
	}

	return db
}
