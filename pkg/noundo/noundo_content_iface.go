package noundo

// TODO iface read + iface write

// ############################ //
// ~~~~~~ Postable Ifaces~~~~~ //
type PostableIface interface {
	Id() string
	Content() string
	Author() UserIdentityIface
	Timestamp() int64
}

type StoryIface interface {
	PostableIface
	AnswerableIface

	Title() string
}

type AnswerIface interface {
	PostableIface
	AnswerableIface

	ParentId() string
}

type AnswerableIface interface {
	Answers() []AnswerIface
}

func (u *UserInfo) Username() string {
	return u.username
}

func (u *UserInfo) ParentServerName() string {
	return u.parentServer
}

func (u *UserInfo) FullUsername() string {
	return u.username + "@" + u.parentServer
}
