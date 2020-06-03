package Map

import "fmt"

type User map[string]interface{}

func NewUser() User {
	return make(map[string]interface{})
}

//func (this User) With(k string,v string) {
//	this[k]=v
//}
func (this User) With(k string, v interface{}) User {
	this[k] = v
	return this
}

func (this User) String() string {
	str := ""
	for k, v := range this {
		fmt.Printf("%v->%v\n", k, v)
	}
	return str
}
