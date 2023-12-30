package noundo

func NewVolatileUser(id Id, email string, username string, parentServerURL string) UserIface {
	return &volatileUser{
		id:              id,
		username:        username,
		email:           email,
		parentServerURL: parentServerURL,
	}
}

type volatileUser struct {
	id              Id
	username        string
	email           string
	parentServerURL string
}

// Id is unchangable, is unique, and is used by server for relations
func (u *volatileUser) Id() Id {
	return u.id
}

// Email is the string that the user will use to authenticated themselves, Email is unique in context of History
func (u *volatileUser) Email() string {
	return u.email
}

// Username is the string that the user will go by, Username is unique in context of History
func (u *volatileUser) Username() string {
	return u.username
}

// The server that is the parent for this account
func (u *volatileUser) ParentServer() string {
	return u.parentServerURL
}
