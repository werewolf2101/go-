package Object

type User struct {
	Id   int
	Name string
	Sex  int
}

//有选择性的对ID进行赋值
func NewUser(f func(u *User)) *User {
	u := new(User)
	f(u)
	return u
}

//只返回func本身不返回结果
func WithUserID(id int) func(u *User) {
	return func(u *User) {
		u.Id = id
	}
}
