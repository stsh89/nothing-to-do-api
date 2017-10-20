package models

import (
	"encoding/json"
)

type User struct {
	name string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) ToJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"name": u.name,
	})
}

func NewUser(name string) User {
	return User{name}
}
