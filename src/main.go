package main

import (
	"base/src/String"
	"fmt"
)

func main()  {
	s:=String.From("abc")
	fmt.Println(s,s.Len())

	ss:=String.FromInt(123)
	fmt.Println(ss)
}
