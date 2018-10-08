package main

import (
	"account"
	"fmt"
)

func main() {
	var u account.UserPasswordChange
	u.ID = "4"
	u.Password = "4567asd"
	result := account.ChangeUserPassword("4", u)
	fmt.Println(result)
}