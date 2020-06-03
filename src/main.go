package main

import (
	"base/src/Map"
	"fmt"
)

//func change(u Object.User) Object.User {
//	u.Id=200
//	return u
//}

//map chan slice
//func change(u Map.User) {
//	u["id"] = "503"
//}

func main() {
	u := Map.NewUser()
	//u["id"] = "100"
	//u["name"] = "MIAbon"
	//change(u)
	//fmt.Println(u)

	//u.With("id", "100")
	//u.With("name", "MIAbon")

	u.With("id", 100).
		With("name", "MIAbon")
	fmt.Println(u)
}
