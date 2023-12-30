package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const PasswordHashCost = 14

type AuthenticatorIface interface {
	// Validate if the passed in credentials are valid
	ValidateAuthMe(*LoginMe) error

	// Add User to the database of users
	RegisterUser(*RegisterMe) *RegisterMeResponse

	//
	GetUserByEmail(email string) UserAuthIface
	GetUserByUsername(username string) UserAuthIface
}

type AuthenticatorStorageIface interface {
	CreateUserOrErr(email string, username string, password []byte) MsgEnum
	GetUserByEmail(email string) (UserAuthIface, error)
	GetUserByUsername(username string) (UserAuthIface, error)
}

type UserIdentity interface {
	Email() string
	Username() string
}

type UserPasswd interface {
	PasswdHash() []byte
}

type UserAuthIface interface {
	UserIdentity
	UserPasswd
}

type inramUser struct {
	email      string
	username   string
	passwdHash []byte
}

func (u *inramUser) Email() string {
	return u.email
}

func (u *inramUser) Username() string {
	return u.username
}

func (u *inramUser) PasswdHash() []byte {
	return u.passwdHash
}

type Authenticator struct {
	authStorage AuthenticatorStorageIface
}

type volatileAuthStorage struct {
	emailUsers    map[string]UserAuthIface
	usernameUsers map[string]UserAuthIface
}

// ~~~ Authenticator Storage ~~~
func NewVolatileAuthStorage() AuthenticatorStorageIface {
	return &volatileAuthStorage{
		emailUsers:    make(map[string]UserAuthIface),
		usernameUsers: make(map[string]UserAuthIface),
	}
}

func (va *volatileAuthStorage) CreateUserOrErr(email, username string, password []byte) MsgEnum {
	if _, ok := va.emailUsers[email]; ok {
		return EmailInUse
	}
	if _, ok := va.usernameUsers[username]; ok {
		return UsernameInUse
	}
	u := &inramUser{
		email:      email,
		username:   username,
		passwdHash: password,
	}
	va.emailUsers[email] = u
	va.usernameUsers[username] = u
	return Ok
}

func (va *volatileAuthStorage) GetUserByEmail(email string) (UserAuthIface, error) {
	user, ok := va.emailUsers[email]
	return utils.ResultOkToErr(user, ok)("email_not_found")
}

func (va *volatileAuthStorage) GetUserByUsername(username string) (UserAuthIface, error) {
	user, ok := va.usernameUsers[username]
	return utils.ResultOkToErr(user, ok)("username_not_found")
}

// ~~~ Authenticator ~~~
func NewAuthenticator(as AuthenticatorStorageIface) AuthenticatorIface {
	a := &Authenticator{as}
	return a
}

func (a *Authenticator) ValidateAuthMe(am *LoginMe) error {
	user, err := a.authStorage.GetUserByEmail(am.Email)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(user.PasswdHash(), []byte(am.Password))
}

func (a *Authenticator) RegisterUser(rm *RegisterMe) *RegisterMeResponse {
	_, err := a.authStorage.GetUserByEmail(rm.Email)
	if err == nil {
		return &RegisterMeResponse{RestResp{false, EmailInUse}}
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(rm.Password), PasswordHashCost)
	if err != nil {
		return &RegisterMeResponse{RestResp{false, Err}}
	}

	msg := a.authStorage.CreateUserOrErr(rm.Email, rm.Username, hashBytes)
	return &RegisterMeResponse{RestResp{msg == Ok, msg}}
}

func (a *Authenticator) GetUserByEmail(email string) UserAuthIface {
	return utils.Left(a.authStorage.GetUserByEmail(email))
}

func (a *Authenticator) GetUserByUsername(username string) UserAuthIface {
	return utils.Left(a.authStorage.GetUserByUsername(username))
}
