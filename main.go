package main

import (
	"./models"
	"fmt"
)

func main() {
	user := models.NewUser("Stan")
	fmt.Println(user.GetName())
	user.SetName("Roland")
	fmt.Println(user.GetName())
}
