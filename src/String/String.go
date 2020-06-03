package String

import (
	"fmt"
)

type String string //我自己的 String
func (this String) Len() int{
	return len(this)
}

func From(str string) String {
	return String(str)
}

func FromInt(n int) String {
	//return String(strconv.Itoa(n)) //相同
	return String(fmt.Sprintf("%d",n))
}