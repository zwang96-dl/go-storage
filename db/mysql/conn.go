package mysql

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	db, _ := sql.Open("mysql", "root:123456@tcp(157.230.169.141:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect mysql server")
		os.Exit(1)
	}
}

func DBConn() *sql.DB {
	if db != nil {
		return db
	}
	db, _ := sql.Open("mysql", "root:123456@tcp(157.230.169.141:3306)/fileserver?charset=utf8")
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect mysql server")
		os.Exit(1)
	}
	return db
}
