package Object

type User struct {
	Id   int
	Name string
	Sex  int
}

//有选择性的对ID进行赋值 //代码量增加但是调用方便
func NewUser(fs ...func(u *User)) *User { //...可变参数
	u := new(User)
	for _, f := range fs {
		f(u)
	}
	return u
}
