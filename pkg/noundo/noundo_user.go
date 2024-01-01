package noundo

type UserIface interface {
	// Id is unchangable, is unique, and is used by server for relations
	Id() Id

	// Email is the string that the user will use to authenticated themselves, Email is unique in context of History
	Email() string

	// Username is the string that the user will go by, Username is unique in context of History
	Username() string

	// The server that is the parent for this account
	ParentServer() string
}
