package routings

import (
	"../models"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	user := models.NewUser("Stan")
	b, _ := user.ToJSON()
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
