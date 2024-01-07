package noundo

import "github.com/kacpekwasny/noundo/pkg/enums"

// TODO iface read + iface write

type StoryIface interface {
	PostableIface
	AnswerableIface
}

type AnswerIface interface {
	PostableIface
	AnswerableIface
}

type PostableIface interface {
	Id() Id
	Content() string
	AuthorFUsername() string
	ReactionableIface
}

type AnswerableIface interface {
	AddAnswer(author UserFullIface, answerable AnswerableIface, answer AnswerIface) (AnswerIface, error)
	Answers(start int, end int, depth int, order OrderIface[AnswerableIface], filter FilterIface[AnswerableIface], ages []AgeIface) ([]AnswerIface, error)
}

type ReactionableIface interface {
	ReactionStats() (map[enums.ReactionType]int, error)
	Reactions() ([]ReactionIface, error)
	React(user UserFullIface, reaction ReactionIface) error
}

type ReactionIface interface {
	Id() Id
	Enum() enums.ReactionType
	ParentId() Id
	AuthorFUsername() string
	Timestamp() uint64
}
