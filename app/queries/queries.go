package queries

import (
	"../../db"
	"../models"
)

func GetAllUsers() []models.User {
	var name string
	var user models.User

	users := []models.User{}
	rows, _ := db.Conn.Query("select name from users")
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&name)
		user = models.NewUser(name)
		users = append(users, user)
	}

	return users
}
