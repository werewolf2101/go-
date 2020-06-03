package Object

//只返回func本身不返回结果
func WithUserID(id int) func(u *User) {
	return func(u *User) {
		u.Id = id
	}
}
func WithUserName(name string) func(u *User) {
	return func(u *User) {
		u.Name = name
	}
}
