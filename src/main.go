package main

import (
	"base/src/Object"
	"fmt"
)

//func change(u Object.User) Object.User {
//	u.Id=200
//	return u
//}

func main() {
	u := Object.NewUser(
		Object.WithUserName("aaa"),
		Object.WithUserID(110),
	)
	fmt.Println(u)
}
