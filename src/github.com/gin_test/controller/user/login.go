package user

import (
	"fmt"
)

type User struct {
	name     string
	password string
}

func (u *User) UserLogin(data string) string {
	fmt.Println("data==", data)
	return "hello " + data
}
