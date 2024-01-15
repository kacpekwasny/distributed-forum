package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

const DEFAULT_PASS_HASH_COST = 14

type AuthenticatorIface interface {
	// Validate if the passed in credentials are valid
	ValidateAuthMe(*SignInRequest) error

	// Add User to the database of users
	SignUpUser(*SignUpRequest) *SignUpResponse

	GetUserByEmail(email string) UserAuthIface
	GetUserByUsername(username string) UserAuthIface
}

// ~~~ Authenticator ~~~

type Authenticator struct {
	authStorage      AuthenticatorStorageIface
	PasswordHashCost int
}

func NewAuthenticator(as AuthenticatorStorageIface, PasswordHashCost int) AuthenticatorIface {
	return &Authenticator{as, PasswordHashCost}
}

func (a *Authenticator) ValidateAuthMe(am *SignInRequest) error {
	user, err := a.authStorage.GetUserByEmail(am.Email)
	if err != nil {
		return err
	}
	return bcrypt.CompareHashAndPassword(user.PasswdHash(), []byte(am.Password))
}

func (a *Authenticator) SignUpUser(rm *SignUpRequest) *SignUpResponse {
	_, err := a.authStorage.GetUserByEmail(rm.Email)
	if err == nil {
		return &SignUpResponse{RestResp{false, EmailInUse}}
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(rm.Password), a.PasswordHashCost)
	if err != nil {
		return &SignUpResponse{RestResp{false, Err}}
	}

	msg := a.authStorage.CreateUserOrErr(rm.Email, rm.Username, hashBytes)
	return &SignUpResponse{RestResp{msg == Ok, msg}}
}

func (a *Authenticator) GetUserByEmail(email string) UserAuthIface {
	return utils.Left(a.authStorage.GetUserByEmail(email))
}

func (a *Authenticator) GetUserByUsername(username string) UserAuthIface {
	return utils.Left(a.authStorage.GetUserByUsername(username))
}
