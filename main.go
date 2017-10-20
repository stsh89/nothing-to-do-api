package main

import (
	"./database"
	"./routings"
	"net/http"
)

func main() {
	database.GetConn()
	http.HandleFunc("/", routings.Root)
	http.ListenAndServe(":3000", nil)
}
