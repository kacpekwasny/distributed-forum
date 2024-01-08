package noundo

func NewVolatileUser(email string, username string, passwdHash []byte, parentServerURL string) UserFullIface {
	return &volatileUser{
		username:         username,
		email:            email,
		parentServerName: parentServerURL,
		passwdHash:       passwdHash,
	}
}

type volatileUser struct {
	username         string
	email            string
	parentServerName string
	passwdHash       []byte
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
func (u *volatileUser) ParentServerName() string {
	return u.parentServerName
}

// The server that is the parent for this account
func (u *volatileUser) FullUsername() string {
	return u.username + "@" + u.parentServerName
}

func (u *volatileUser) PasswdHash() []byte {
	return u.passwdHash
}
