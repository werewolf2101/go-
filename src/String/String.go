package String

import (
	"fmt"
)

type String string //我自己的 String
func (this String) Len() int {
	return len(this)
}

func From(str string) String {
	return String(str)
}

func (this String) Each(f func(item string)) {
	//英文
	//for i := 0; i < len(this); i++ {
	//	f(fmt.Sprintf("%c", this[i]))
	//}
	//中文
	for _, c := range this {
		f(fmt.Sprintf("%c", c))
	}
}

func FromInt(n int) String {
	//return String(strconv.Itoa(n)) //相同
	return String(fmt.Sprintf("%d", n))
}
