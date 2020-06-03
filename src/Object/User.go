package Object

type User struct {
	Id   int
	Name string
	Sex  int
}

func NewUser() User {
	return User{}
}
