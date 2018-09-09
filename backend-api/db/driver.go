package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var Conn *sql.DB

func init() {
	user := os.Getenv("MYSQL_USER") // mysql
	pass := os.Getenv("MYSQL_PASSWORD") // mysql
	name := os.Getenv("MYSQL_DATABASE") // itemsDB

	dbconf := user + ":" + pass "@/" + name
	conn, err := sql.Open("mysql", dbconf)
	if err != nil {
		panic(err.Error)
	}
	Conn = conn
}
