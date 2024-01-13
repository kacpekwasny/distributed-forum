package noundo

import "github.com/kacpekwasny/noundo/pkg/utils"

type volatileAuthStorage struct {
	emailUsers    *map[string]UserFullIface
	usernameUsers *map[string]UserFullIface
}

// ~~~ Authenticator Storage ~~~
func NewVolatileAuthStorage(
	emailUsers *map[string]UserFullIface,
	usernameUsers *map[string]UserFullIface,
) AuthenticatorStorageIface {
	return &volatileAuthStorage{
		emailUsers:    emailUsers,
		usernameUsers: usernameUsers,
	}
}

// TODO multithread handling
func (va *volatileAuthStorage) CreateUserOrErr(email, username string, password []byte) MsgEnum {
	if _, ok := (*va.emailUsers)[email]; ok {
		return EmailInUse
	}
	if _, ok := (*va.usernameUsers)[username]; ok {
		return UsernameInUse
	}
	u := &User{
		email:      email,
		username:   username,
		passwdHash: password,
	}
	(*va.emailUsers)[email] = u
	(*va.usernameUsers)[username] = u
	return Ok
}

func (va *volatileAuthStorage) GetUserByEmail(email string) (UserAuthIface, error) {
	user, ok := (*va.emailUsers)[email]
	return utils.ResultOkToErr(user, ok)("email_not_found")
}

func (va *volatileAuthStorage) GetUserByUsername(username string) (UserAuthIface, error) {
	user, ok := (*va.usernameUsers)[username]
	return utils.ResultOkToErr(user, ok)("username_not_found")
}
