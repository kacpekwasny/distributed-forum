package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const PasswordHashCost = 14

type VolatileAuthenticator struct {
	loginPasswdHash map[string][]byte
	loginUsers      map[string]UserIface
}

func NewVolatileAuthenticator() *VolatileAuthenticator {
	va := &VolatileAuthenticator{
		loginPasswdHash: make(map[string][]byte),
		loginUsers:      make(map[string]UserIface),
	}
	va.RegisterUser(&RegisterMe{
		Login:    "awd",
		Username: "awd",
		Password: "awd",
	})
	return va
}

func (va *VolatileAuthenticator) ValidateAuthMe(am *LoginMe) error {
	hash, ok := va.loginPasswdHash[am.Login]
	if !ok {
		return errors.New("login not found to be registered")
	}
	return bcrypt.CompareHashAndPassword(hash, []byte(am.Password))
}

func (va *VolatileAuthenticator) RegisterUser(rm *RegisterMe) *RegisterMeResponse {
	_, ok := va.loginPasswdHash[rm.Login]
	if ok {
		return &RegisterMeResponse{RestResp{false, LoginInUse}}
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(rm.Password), PasswordHashCost)
	if err != nil {
		return &RegisterMeResponse{RestResp{false, Err}}
	}
	va.loginPasswdHash[rm.Login] = bytes
	va.loginUsers[rm.Login] = NewSimpleUser(rm.Login, rm.Username)
	return &RegisterMeResponse{RestResp{true, Ok}}
}

func (va *VolatileAuthenticator) GetUserByLogin(login string) UserIface {
	return va.loginUsers[login]
}
