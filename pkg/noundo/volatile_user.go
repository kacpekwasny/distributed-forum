package noundo

func NewVolatileUser(id Id, login string, username string, parentServerURL string) UserIface {
	return &volatileUser{
		id:              id,
		username:        username,
		login:           login,
		parentServerURL: parentServerURL,
	}
}

type volatileUser struct {
	id              Id
	username        string
	login           string
	parentServerURL string
}

// Id is unchangable, is unique, and is used by server for relations
func (u *volatileUser) Id() Id {
	return u.id
}

// Login is the string that the user will use to authenticated themselves, Login is unique in context of History
func (u *volatileUser) Login() string {
	return u.login
}

// Username is the string that the user will go by, Username is unique in context of History
func (u *volatileUser) Username() string {
	return u.username
}

// The server that is the parent for this account
func (u *volatileUser) ParentServer() string {
	return u.parentServerURL
}
