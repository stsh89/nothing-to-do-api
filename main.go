package main

import (
	"./db"
	"./routings"
	"net/http"
)

func main() {
	db.GetConn()
	http.HandleFunc("/", routings.Root)
	http.ListenAndServe(":3000", nil)
}
