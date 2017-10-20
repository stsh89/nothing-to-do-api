package main

import (
	"./db"
	"./routings"
	"net/http"
)

func main() {
	db.GetConn()
	routings.Route()
	http.ListenAndServe(":3000", nil)
}
