package forum

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

func newJWTMapClaims(jf JWTFields) jwt.MapClaims {
	return jwt.MapClaims{
		"login": jf.Login,
	}
}

func jwtFieldsFromMapClaims(mp jwt.MapClaims) (JWTFields, error) {
	var jf JWTFields
	login, ok := mp["login"]
	if !ok {
		return jf, errors.New("missing field login")
	}
	return JWTFields{
		Login: login,
	}
}
