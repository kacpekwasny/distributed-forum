package noundo

func NewUserStruct(email, username, parentServerName string, passwdHash []byte) *User {
	return &User{
		email:            email,
		username:         username,
		passwdHash:       passwdHash,
		parentServerName: parentServerName,
	}
}

type User struct {
	email            string
	username         string
	passwdHash       []byte
	parentServerName string
}

type EmailIface interface {
	// Email is the string that the user will use to authenticated themselves, Email is unique in context of History
	Email() string
}

type UsernameIface interface {
	// Username is the string that the user will go by, Username is unique in context of History
	Username() string
}

type PasswdhashIface interface {
	PasswdHash() []byte
}

type FullUsernameIface interface {
	// Domain of the server that is the parent for this account
	ParentServerName() string

	// Username() + "@" + ParentServerName()`
	FullUsername() string
}

type UserPublicIface interface {
	UsernameIface
	FullUsernameIface
}

type UserFullIface interface {
	EmailIface
	UsernameIface
	PasswdhashIface
	FullUsernameIface
}

type UserAuthIface interface {
	EmailIface
	UserPublicIface
	PasswdhashIface
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Username() string {
	return u.username
}

func (u *User) PasswdHash() []byte {
	return u.passwdHash
}

// Domain of the server that is the parent for this account
func (u *User) ParentServerName() string {
	return u.parentServerName
}

// Username() + "@" + ParentServerName()`
func (u *User) FullUsername() string {
	return u.Username() + "@" + u.ParentServerName()
}
