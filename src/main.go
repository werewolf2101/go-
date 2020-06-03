package main

import (
	"base/src/Object"
	"base/src/String"
	"fmt"
)

//func change(u Object.User) Object.User {
//	u.Id=200
//	return u
//}

func main() {
	s := String.From("我abc")
	fmt.Println(s, s.Len())

	//for i:=0; i < len(s); i++{
	//	fmt.Println(i, " ", fmt.Sprintf("%c", s[i])) //不转换打印出的是阿斯克码
	//}

	s.Each(func(item string) {
		fmt.Println(item)
	})

	ss := String.FromInt(123)
	fmt.Println(ss)

	u := Object.NewUser()
	u.Id = 100
	//u = change(u)

	//fmt.Println(u)
	//fmt.Printf("%+v", u) //显示属性名

	//打印指针地址
	fmt.Println(&u)
	fmt.Printf("%p\n", &u)

	uu := Object.NewUsers()
	fmt.Println(uu)
}
