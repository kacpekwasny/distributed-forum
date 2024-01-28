package noundo

import "github.com/kacpekwasny/noundo/pkg/utils"

type persistentUserAuth struct {
	email            string
	username         string
	passwdHash       []byte // TODO remove, and do something with this object, make a UserIdentityObject, or make it the JWT
	parentServerName string
	aboutMe          string
	accountBirthDate int64
}

type persistentAuthStorage struct {
	serverName    string
	emailUsers    *map[string]UserAllIface
	usernameUsers *map[string]UserAllIface
}

// ~~~ Authenticator Storage ~~~
func NewPersistentAuthStorage(
	serverName string,
	emailUsers *map[string]UserAllIface,
	usernameUsers *map[string]UserAllIface,
) AuthenticatorStorageIface {
	return &persistentAuthStorage{
		serverName:    serverName,
		emailUsers:    emailUsers,
		usernameUsers: usernameUsers,
	}
}

// TODO multithread handling
func (va *persistentAuthStorage) CreateUserOrErr(email, username string, password []byte) MsgEnum {
	if _, ok := (*va.emailUsers)[email]; ok {
		return EmailInUse
	}
	if _, ok := (*va.usernameUsers)[username]; ok {
		return UsernameInUse
	}
	u := &persistentUserAuth{
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

func (va *persistentAuthStorage) GetUserByEmail(email string) (UserAuthIface, error) {
	user, ok := (*va.emailUsers)[email]
	return utils.ResultOkToErr(user, ok)("email_not_found")
}

func (va *persistentAuthStorage) GetUserByUsername(username string) (UserAuthIface, error) {
	user, ok := (*va.usernameUsers)[username]
	return utils.ResultOkToErr(user, ok)("username_not_found")
}

func (u *persistentUserAuth) Email() string {
	return u.email
}

func (u *persistentUserAuth) Username() string {
	return u.username
}

func (u *persistentUserAuth) PasswdHash() []byte {
	return u.passwdHash
}

// Domain of the server that is the parent for this account
func (u *persistentUserAuth) ParentServerName() string {
	return u.parentServerName
}

// Username() + "@" + ParentServerName()`
func (u *persistentUserAuth) FullUsername() string {
	return u.Username() + "@" + u.ParentServerName()
}

func (u *persistentUserAuth) AboutMe() string {
	return u.aboutMe
}

func (u *persistentUserAuth) AccountBirthDate() int64 {
	return u.accountBirthDate
}
