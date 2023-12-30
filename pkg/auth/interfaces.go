package auth

import "github.com/kacpekwasny/noundo/pkg/noundo"

type UserIface interface {
	Id() noundo.Id
	Login() string
	Username() string
	ParentServer() string
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

func (u *SimpleUser) Id() noundo.Id {
	return 1234567
}

func (u *SimpleUser) ParentServer() string {
	return "http://parent"
}

type Authenticator interface {
	// Validate if the passed in credentials are valid
	ValidateAuthMe(*LoginMe) error

	// Add User to the database of users
	RegisterUser(*RegisterMe) *RegisterMeResponse

	//
	GetUserByLogin(login string) UserIface
	GetUserByUsername(username string) UserIface
}
