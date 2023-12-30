package noundo

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

func newJWTMapClaims(jf JWTFields) jwt.MapClaims {
	return jwt.MapClaims{
		"email": jf.Username,
	}
}

func jwtFieldsFromMapClaims(mp jwt.MapClaims) (JWTFields, error) {
	var jf JWTFields
	emailIface, ok1 := mp["email"]
	email, ok2 := emailIface.(string)
	if !(ok1 && ok2) {
		return jf, errors.New("missing field email")
	}
	return JWTFields{
		Username: email,
	}, nil
}

func NewRegisterMe(email, username, password string) *RegisterMe {
	return &RegisterMe{
		Email:    email,
		Username: username,
		Password: password,
	}
}
