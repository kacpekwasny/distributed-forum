package noundo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type persistentUserAuth struct {
	id               int
	email            string
	username         string
	passwdHash       []byte // TODO remove, and do something with this object, make a UserIdentityObject, or make it the JWT
	parentServerName string
	aboutMe          string
	accountBirthDate int64
}

type persistentAuthStorage struct {
	serverName string
	dbPool     *pgxpool.Pool
}

// ~~~ Authenticator Storage ~~~
func NewPersistentAuthStorage(
	serverName string,
	dbPool *pgxpool.Pool,

) AuthenticatorStorageIface {
	return &persistentAuthStorage{
		serverName: serverName,
		dbPool:     dbPool,
	}
}

// TODO multithread handling
func (va *persistentAuthStorage) CreateUserOrErr(email, username string, password []byte) MsgEnum {
	//check if there's already a user with this email or username
	err := va.dbPool.QueryRow(context.Background(), "SELECT id FROM users WHERE email = $1 ", email).Scan()
	if err == nil {
		return EmailInUse
	}
	err = va.dbPool.QueryRow(context.Background(), "SELECT id FROM users WHERE username = $1 ", username).Scan()
	if err == nil {
		return UsernameInUse
	}
	var id int
	err = va.dbPool.QueryRow(context.Background(), "INSERT INTO users ( email, username, password_hash, parent_server, account_birth_date, about_me) VALUES ($1, $2, $3, $4,current_timestamp, $5) RETURNING id", email, username, password, va.serverName, "Hi I am: "+username).Scan(&id)
	if err != nil {
		return Err
	}
	return Ok
}

func (va *persistentAuthStorage) GetUserByEmail(email string) (UserAuthIface, error) {
	//get the user from the db
	var id int
	var username, parentServerName, aboutMe string
	var passwordHash []byte

	var accountBirthDate int64
	err := va.dbPool.QueryRow(context.Background(), "SELECT id, username, parent_server, account_birth_date, about_me, password_hash FROM users WHERE email = $1", email).Scan(&id, &username, &parentServerName, &accountBirthDate, &aboutMe, &passwordHash)
	if err != nil {
		return nil, err
	}
	return persistentUserAuth{
		id:               id,
		email:            email,
		username:         username,
		passwdHash:       nil,
		parentServerName: parentServerName,
		aboutMe:          aboutMe,
		accountBirthDate: accountBirthDate,
	}, nil
}

func (va *persistentAuthStorage) GetUserByUsername(username string) (UserAuthIface, error) {
	//get the user from the db
	var id int
	var email, parentServerName, aboutMe string
	var accountBirthDate int64
	var passwordHash []byte
	err := va.dbPool.QueryRow(context.Background(), "SELECT id, email, parent_server, account_birth_date, about_me, password_hash FROM users WHERE username = $1", username).Scan(&id, &email, &parentServerName, &accountBirthDate, &aboutMe, &passwordHash)
	if err != nil {
		return nil, err
	}
	return persistentUserAuth{
		id:               id,
		email:            email,
		username:         username,
		passwdHash:       passwordHash,
		parentServerName: parentServerName,
		aboutMe:          aboutMe,
		accountBirthDate: accountBirthDate,
	}, nil

}

func (u persistentUserAuth) Id() int {
	return u.id
}
func (u persistentUserAuth) Email() string {
	return u.email
}

func (u persistentUserAuth) Username() string {
	return u.username
}

func (u persistentUserAuth) PasswdHash() []byte {
	return u.passwdHash
}

// Domain of the server that is the parent for this account
func (u persistentUserAuth) ParentServerName() string {
	return u.parentServerName
}

// Username() + "@" + ParentServerName()`
func (u persistentUserAuth) FullUsername() string {
	return u.Username() + "@" + u.ParentServerName()
}

func (u persistentUserAuth) AboutMe() string {
	return u.aboutMe
}

func (u persistentUserAuth) AccountBirthDate() int64 {
	return u.accountBirthDate
}
