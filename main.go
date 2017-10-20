package main

import (
	"./routings"
	"net/http"
)

func main() {
	http.HandleFunc("/", routings.Root)
	http.ListenAndServe(":3000", nil)
}
