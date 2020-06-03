package Object

//参数可省略
//type UserAttrFunc func(u *User)
type UserAttrFunc func(*User) //设置User属性的函数类型
type UserAttrFuncs []UserAttrFunc

func (this UserAttrFuncs) apply(u *User) {
	for _, f := range this {
		f(u)
	}
}

//只返回func本身不返回结果
func WithUserID(id int) UserAttrFunc {
	return func(u *User) {
		u.Id = id
	}
}
func WithUserName(name string) UserAttrFunc {
	return func(u *User) {
		u.Name = name
	}
}
