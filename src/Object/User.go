package Object

type User struct {
	Id   int
	Name string
	Sex  int
}

func NewUser() User {
	return User{}
}

func NewUsers() *User {
	//return new(User)
	//效果一样
	return &User{}
}
