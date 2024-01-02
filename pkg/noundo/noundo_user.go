package noundo

type UserIdentityIface interface {
	// Email is the string that the user will use to authenticated themselves, Email is unique in context of History
	Email() string

	// Username is the string that the user will go by, Username is unique in context of History
	Username() string
}

type ParentServerNameIface interface {
	// Domain of the server that is the parent for this account
	ParentServerName() string
}

type PasswdhashIface interface {
	PasswdHash() []byte
}

type FullUsernameIface interface {

	// Username() + "@" + ParentServerName()`
	FullUsername() string
}

type UserPublicIface interface {
	UserIdentityIface
	FullUsernameIface
}

type UserFullIface interface {
	UserIdentityIface
	PasswdhashIface
	FullUsernameIface
}

type UserAuthIface interface {
	UserIdentityIface
	PasswdhashIface
}

type UserStruct struct {
	email            string
	username         string
	passwdHash       []byte
	parentServerName string
}

func (u *UserStruct) Email() string {
	return u.email
}

func (u *UserStruct) Username() string {
	return u.username
}

func (u *UserStruct) PasswdHash() []byte {
	return u.passwdHash
}

// Domain of the server that is the parent for this account
func (u *UserStruct) ParentServerName() string {
	return u.parentServerName
}

// Username() + "@" + ParentServerName()`
func (u *UserStruct) FullUsername() string {
	return u.Username() + "@" + u.ParentServerName()
}
