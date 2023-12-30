package noundo

type SimpleUser struct {
	email    string
	username string
}

func NewSimpleUser(email string, username string) *SimpleUser {
	return &SimpleUser{
		email:    email,
		username: username,
	}
}

func (u *SimpleUser) Email() string {
	return u.email
}

func (u *SimpleUser) Username() string {
	return u.username
}

func (u *SimpleUser) Id() Id {
	return 1234567
}

func (u *SimpleUser) ParentServer() string {
	return "http://parent"
}
