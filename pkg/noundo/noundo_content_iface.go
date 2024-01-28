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

type StoryContentIface interface {
	GetTitle() string
	GetContent() string
}

type AnswerableIface interface {
	Answers() []AnswerIface
}

func (u *UserInfo) GetUsername() string {
	return u.username
}

func (u *UserInfo) GetParentServerName() string {
	return u.parentServer
}

func (u *UserInfo) GetFUsername() string {
	return u.username + "@" + u.parentServer
}

func (s *StoryContent) GetTitle() string {
	return s.Title
}

func (s *StoryContent) GetContent() string {
	return s.Content
}
