package noundo

// TODO iface read + iface write

// type StoryIface interface {
// 	PostableIface
// 	AnswerableIface
// 	Title() string
// }

// type AnswerIface interface {
// 	PostableIface
// 	AnswerableIface
// }

// type PostableIface interface {
// 	Id() string
// 	Content() string
// 	AuthorFUsername() string
// 	ReactionableIface
// }

// type AnswerableIface interface {
// 	AddAnswer(author UserIdentityIface, answerable AnswerableIface, answer AnswerIface) (AnswerIface, error)
// 	Answers(start int, end int, depth int, order OrderIface[AnswerableIface], filter FilterIface[AnswerableIface], ages []AgeIface) ([]AnswerIface, error)
// }

// type ReactionableIface interface {
// 	ReactionStats() (map[enums.ReactionType]int, error)
// 	Reactions() ([]ReactionIface, error)
// 	React(user UserIdentityIface, reaction ReactionIface) error
// }

// type ReactionIface interface {
// 	Id() Id
// 	Enum() enums.ReactionType
// 	ParentId() Id
// 	AuthorFUsername() string
// 	Timestamp() uint64
// }

func (u *UserInfo) Username() string {
	return u.username
}

func (u *UserInfo) ParentServerName() string {
	return u.parentServer
}

func (u *UserInfo) FullUsername() string {
	return u.username + "@" + u.parentServer
}
