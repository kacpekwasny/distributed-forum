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
	passwdHash       []byte // TODO remove, and do something with this object, make a UserIdentityObject, or make it the JWT
	parentServerName string
	aboutMe          string
}

type UserIdentityIface interface {
	// GetUsername is the string that the user will go by, GetUsername is unique in context of History
	GetUsername() string

	// Domain of the server that is the parent for this account
	GetParentServerName() string

	// Username() + "@" + ParentServerName()`
	GetFUsername() string
}

type UserMoreInfoIface interface {
	GetAboutMe() string
	GetAccountBirthDate() int64
}

type UserPublicIface interface {
	UserIdentityIface
	UserMoreInfoIface
}

type UserAuthIface interface {
	UserIdentityIface

	// Email is the string that the user will use to authenticated themselves, Email is unique in context of History
	Email() string

	// Password Hash - retrieve bytes from database and compare
	PasswdHash() []byte
}

type UserAllIface interface {
	UserPublicIface
	UserAuthIface
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

func (u *User) AboutMe() string {
	return u.aboutMe
}
