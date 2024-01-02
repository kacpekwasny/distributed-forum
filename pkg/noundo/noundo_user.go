package noundo

type UserIdentityIface interface {
	// Email is the string that the user will use to authenticated themselves, Email is unique in context of History
	Email() string

	// Username is the string that the user will go by, Username is unique in context of History
	Username() string
}

type ParentServerNameIface interface {
	// Domain of the server that is the parent for this account
	ParentServerName() string
}

type PasswdhashIface interface {
	PasswdHash() []byte
}

type FullUsernameIface interface {

	// Username() + "@" + ParentServerName()`
	FullUsername() string
}

type UserFullIface interface {
	UserIdentityIface
	PasswdhashIface
	FullUsernameIface
}
