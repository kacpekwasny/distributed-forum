package noundo

import "github.com/kacpekwasny/noundo/pkg/enums"

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
	Body() string
	Author() UserFullIface
	AuthorId() Id
	ReactionableIface
}

type AnswerableIface interface {
	AddAnswer(author UserFullIface, answerable AnswerableIface, answer AnswerIface) (AnswerIface, error)
	Answers(start int, end int, depth int, order OrderIface, filter FilterIface, ages []AgeIface) ([]AnswerIface, error)
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
	AuthorId() Id
	Author() UserFullIface
	Timestamp() uint64
}

type OrderIface func(idx1, idx2 int) bool
type FilterIface func(v any) bool
