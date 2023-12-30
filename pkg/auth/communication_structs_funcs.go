package auth

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

func newJWTMapClaims(jf JWTFields) jwt.MapClaims {
	return jwt.MapClaims{
		"login": jf.Username,
	}
}

func jwtFieldsFromMapClaims(mp jwt.MapClaims) (JWTFields, error) {
	var jf JWTFields
	loginIface, ok1 := mp["login"]
	login, ok2 := loginIface.(string)
	if !(ok1 && ok2) {
		return jf, errors.New("missing field login")
	}
	return JWTFields{
		Username: login,
	}, nil
}

func NewRegisterMe(login, username, password string) *RegisterMe {
	return &RegisterMe{
		Login:    login,
		Username: username,
		Password: password,
	}
}
