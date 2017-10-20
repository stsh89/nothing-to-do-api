package routings

import (
	"../queries"
	"encoding/json"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	users := queries.GetAllUsers()
	b, _ := json.Marshal(users)
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
