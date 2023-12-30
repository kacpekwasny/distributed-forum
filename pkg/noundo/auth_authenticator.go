package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const PasswordHashCost = 14

type Authenticator interface {
	// Validate if the passed in credentials are valid
	ValidateAuthMe(*LoginMe) error

	// Add User to the database of users
	RegisterUser(*RegisterMe) *RegisterMeResponse

	//
	GetUserByEmail(email string) UserIface
	GetUserByUsername(username string) UserIface
}

type AuthenticatorUserStorage interface {
	CreateUserOrErr(email string, username string, password []byte) MsgEnum
	GetUserByEmail(email string) (UserIface, error)
	GetUserByUsername(username string) (UserIface, error)
}

type authenticator_ struct {
	createUserOrErr   func(email string, username string, password []byte) MsgEnum
	getUserByEmail    func(email string) (UserIface, error)
	getUserByUsername func(username string) (UserIface, error)
}

type volatileAuthStorage struct {
	emailPasswdHash map[string][]byte
	emailUsers      map[string]UserIface
	usernameUsers   map[string]UserIface
}

func NewVolatileAuthenticator() Authenticator {

	va := &volatileAuthStorage{
		emailPasswdHash: make(map[string][]byte),
		emailUsers:      make(map[string]UserIface),
		usernameUsers:   make(map[string]UserIface),
	}
	a := &authenticator_{
		createUserOrErr: func(email, username string, password []byte) MsgEnum {
			if _, ok := va.emailUsers[email]; ok {
				return EmailInUse
			}
			if _, ok := va.usernameUsers[username]; ok {
				return UsernameInUse
			}
			va.emailPasswdHash[email] = password
			return Ok
		},
		getUserByEmail: func(email string) (UserIface, error) {
			user, ok := va.emailUsers[email]
			return utils.ResultOkToErr(user, ok)("email_not_found")
		},
		getUserByUsername: func(username string) (UserIface, error) {
			user, ok := va.usernameUsers[username]
			return utils.ResultOkToErr(user, ok)("username_not_found")
		},
	}
	a.RegisterUser(&RegisterMe{
		Email:    "a",
		Username: "a",
		Password: "a",
	})
	return a
}

func (va *authenticator_) ValidateAuthMe(am *EmailMe) error {
	user, err := va.getUserByEmail(am.Email)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(hash, []byte(am.Password))
}

func (va *authenticator_) RegisterUser(rm *RegisterMe) *RegisterMeResponse {
	_, ok := va.emailPasswdHash[rm.Email]
	if ok {
		return &RegisterMeResponse{RestResp{false, EmailInUse}}
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(rm.Password), PasswordHashCost)
	if err != nil {
		return &RegisterMeResponse{RestResp{false, Err}}
	}
	va.emailPasswdHash[rm.Email] = bytes

	user := NewSimpleUser(rm.Email, rm.Username)
	va.emailUsers[rm.Email] = user
	va.usernameUsers[rm.Username] = user

	return &RegisterMeResponse{RestResp{true, Ok}}
}

func (va *authenticator_) GetUserByEmail(email string) UserIface {
	return va.emailUsers[email]
}

func (va *authenticator_) GetUserByUsername(username string) UserIface {
	return va.usernameUsers[username]
}
