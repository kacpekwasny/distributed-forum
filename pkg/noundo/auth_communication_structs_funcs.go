package noundo

import (
	"errors"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	jwtUsernameKey         = "username"
	jwtIssuedTimestampKey  = "issuedTimestamp"
	jwtParentServerNameKey = "parentServerName"
)

func newJWTMapClaims(jf JWTFields) jwt.MapClaims {
	return jwt.MapClaims{
		jwtUsernameKey:         jf.username,
		jwtParentServerNameKey: jf.parentServerName,
		jwtIssuedTimestampKey:  jf.jwtIssuedTimestamp,
	}
}

func valueFromMapClaims(c jwt.MapClaims, k string) (string, error) {
	v, ok := c[k]
	if !ok {
		return "", errors.New("no key '" + k + "'")
	}
	s, oks := v.(string)
	if !oks {
		return "", errors.New("conversion to string failed")
	}
	return s, nil
}

func jwtFieldsFromMapClaims(mp jwt.MapClaims) (JWTFields, error) {
	var jf JWTFields
	username, errUn := valueFromMapClaims(mp, jwtUsernameKey)
	parentServerName, errSn := valueFromMapClaims(mp, jwtParentServerNameKey)
	issuedTimestampIface, ok := mp[jwtIssuedTimestampKey]
	issuedTimestamp, ok2 := issuedTimestampIface.(float64)

	if (errUn != nil) || (errSn != nil) || !ok || !ok2 {
		return jf, errors.New("read map claims failed")
	}
	return JWTFields{
		username:           username,
		parentServerName:   parentServerName,
		jwtIssuedTimestamp: int64(issuedTimestamp),
	}, nil
}

func NewSignUpRequest(email, username, password string) *SignUpRequest {
	return &SignUpRequest{
		Email:    email,
		Username: username,
		Password: password,
	}
}
