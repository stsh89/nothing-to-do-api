package main

import (
	"./config"
	"./db"
	"net/http"
)

func main() {
	db.GetConn()
	config.Routes()
	http.ListenAndServe(":3000", nil)
}
