package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticatorIface interface {
	// Validate if the passed in credentials are valid
	ValidateAuthMe(*LoginMe) error

	// Add User to the database of users
	RegisterUser(*RegisterMe) *RegisterMeResponse

	//
	GetUserByEmail(email string) UserAuthIface
	GetUserByUsername(username string) UserAuthIface
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

// ~~~ Authenticator ~~~

type Authenticator struct {
	authStorage      AuthenticatorStorageIface
	PasswordHashCost int
}

func NewAuthenticator(as AuthenticatorStorageIface, PasswordHashCost int) AuthenticatorIface {
	return &Authenticator{as, PasswordHashCost}
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

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(rm.Password), a.PasswordHashCost)
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
