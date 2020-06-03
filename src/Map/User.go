package Map

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
