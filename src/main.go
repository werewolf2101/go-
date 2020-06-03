package main

import (
	"base/src/String"
	"fmt"
)

func main() {
	s := String.From("abc")
	fmt.Println(s, s.Len())

	//for i:=0; i < len(s); i++{
	//	fmt.Println(i, " ", fmt.Sprintf("%c", s[i])) //不转换打印出的是阿斯克码
	//}

	s.Each(func(item string) {
		fmt.Println(item)
	})

	ss := String.FromInt(123)
	fmt.Println(ss)
}
