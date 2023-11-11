package auth

type UserI interface {
	Login() string
	Username() string
}

type SimpleUser struct {
	login    string
	username string
}

func NewSimpleUser(login string, username string) *SimpleUser {
	return &SimpleUser{
		login:    login,
		username: username,
	}
}

func (u *SimpleUser) Login() string {
	return u.login
}

func (u *SimpleUser) Username() string {
	return u.username
}

type Authenticator interface {
	// Validate if the passed in credentials are valid
	ValidateAuthMe(*LoginMe) error

	// Add User to the database of users
	RegisterUser(*RegisterMe) *RegisterMeResponse

	//
	GetUserByLogin(login string) UserI
}
