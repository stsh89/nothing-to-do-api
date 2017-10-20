package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Conn *sql.DB

func GetConn() {
	Conn, _ = sql.Open("postgres", "user=postgres dbname=nothing_to_do_api host=127.0.0.1 sslmode=disable")
}
