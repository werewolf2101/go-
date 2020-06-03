package Map

type User map[string]string

func NewUser() User {
	return make(map[string]string)
}
