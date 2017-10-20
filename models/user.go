package models

type User struct {
	name string
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) SetName(name string) {
	u.name = name
}

func NewUser(name string) User {
	return User{name}
}
