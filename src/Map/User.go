package Map

import (
	"fmt"
	"sort"
)

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

func (this User) Fields() []string {
	keys := []string{}
	for k, _ := range this {
		keys = append(keys, k)
	}
	//sort.Strings(keys)
	sort.Sort(sort.Reverse(sort.StringSlice(keys))) //倒序
	return keys
}

func (this User) String() string {
	//str := ""
	//for k, v := range this {
	//	fmt.Printf("%v->%v\n", k, v)
	//}
	//return str

	str := ""
	for index, k := range this.Fields() {
		str += fmt.Sprintf("%d.%v->%v\n", index+1, k, this[k])
	}
	return str
}
