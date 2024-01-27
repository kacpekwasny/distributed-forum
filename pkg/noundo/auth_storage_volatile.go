package noundo

import "github.com/kacpekwasny/noundo/pkg/utils"

type volatileUserAuth struct {
	email            string
	username         string
	passwdHash       []byte // TODO remove, and do something with this object, make a UserIdentityObject, or make it the JWT
	parentServerName string
	aboutMe          string
	accountBirthDate int64
}

type volatileAuthStorage struct {
	serverName    string
	emailUsers    *map[string]UserAllIface
	usernameUsers *map[string]UserAllIface
}

// ~~~ Authenticator Storage ~~~
func NewVolatileAuthStorage(
	serverName string,
	emailUsers *map[string]UserAllIface,
	usernameUsers *map[string]UserAllIface,
) AuthenticatorStorageIface {
	return &volatileAuthStorage{
		serverName:    serverName,
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
	u := &volatileUserAuth{
		email:            email,
		username:         username,
		passwdHash:       password,
		parentServerName: va.serverName,
		aboutMe:          "Hi I am: " + username,
		accountBirthDate: UnixTimeNow(),
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

func (u *volatileUserAuth) Email() string {
	return u.email
}

func (u *volatileUserAuth) Username() string {
	return u.username
}

func (u *volatileUserAuth) PasswdHash() []byte {
	return u.passwdHash
}

// Domain of the server that is the parent for this account
func (u *volatileUserAuth) ParentServerName() string {
	return u.parentServerName
}

// Username() + "@" + ParentServerName()`
func (u *volatileUserAuth) FullUsername() string {
	return u.Username() + "@" + u.ParentServerName()
}

func (u *volatileUserAuth) AboutMe() string {
	return u.aboutMe
}

func (u *volatileUserAuth) AccountBirthDate() int64 {
	return u.accountBirthDate
}
