package noundo

type AuthenticatorStorageIface interface {
	CreateUserOrErr(email string, username string, password []byte) MsgEnum
	GetUserByEmail(email string) (UserAuthIface, error)
	GetUserByUsername(username string) (UserAuthIface, error)
}
