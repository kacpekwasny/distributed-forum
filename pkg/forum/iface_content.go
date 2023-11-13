package forum

import "github.com/kacpekwasny/distributed-forum/pkg/enums"

type Id uint64

// When I think about this, history is an interface for a database in ints goal,
// I might make the volatile database, but in the end HistoryIface and all other interfaces must made in a way,
// That is efficient with the database.
// So my idea is, that everything that cannot be associated strictly with any Iface will be History method.

// History is the name for the whole server, that contains all Ages and All Stories
type HistoryIface interface {

	// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
	CreateAge(owner UserIface, name string) (AgeIface, error)

	//
	GetStory(id Id) (StoryIface, error)

	// Get answer from anywhere in
	GetAnswer(id Id) (AnswerIface, error)

	// GetFirst n stories ordered by different atributes, from []ages,
	GetStories(start uint32, end uint32, order OrderIface, filter FilterIface, ages []AgeIface) []StoryIface

	GetUser(id Id) (UserIface, error)

	// TODO:
	// GetAges that user joined,
	// GetStories first n stories of user ordered by (maybe merge with the first method???)
	// GetComments first n comments of user ordered by
}

type UserIface interface {
	// Id is unchangable, is unique, and is used by server for relations
	Id() Id

	// Login is the string that the user will use to authenticated themselves, Login is unique in context of History
	Login() string

	// Username is the string that the user will go by, Username is unique in context of History
	Username() string

	// The server that is the parent for this account
	ParentServer() string
}

// Are all those other interfaces needed??? TODO
type AgeIface interface {
	GetId() (Id, error)
	GetName() (string, error)
	SetName(name string) error

	GetOwner() (UserIface, error)
	ChangeOwner(user UserIface) error

	GetAdmins() ([]UserIface, error)
	AddAdmin(user UserIface) error
	RemoveAdmin(user UserIface) error

	GetMembers(start uint32, end uint32) ([]UserIface, error)
	GetMembersNumber() (uint32, error)

	// Create a Story written by an Author in a certain Age
	AddStory(author UserIface, age AgeIface, story StoryIface) (StoryIface, error)
}

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
	Author() UserIface
	AuthorId() Id
	ReactionableIface
}

type AnswerableIface interface {
	AddAnswer(author UserIface, answerable AnswerableIface, answer AnswerIface) (AnswerIface, error)
	Answers(start uint32, end uint32, depth uint32, order OrderIface, filter FilterIface, ages []AgeIface) ([]AnswerIface, error)
}

type ReactionableIface interface {
	ReactionStats() (map[enums.ReactionType]uint32, error)
	Reactions() ([]ReactionIface, error)
	React(user UserIface, reaction ReactionIface)
}

type ReactionIface interface {
	Id() Id
	Enum() enums.ReactionType
	ParentId() Id
	AuthorId() Id
	Timestamp() uint64
}

type OrderIface interface{}
type FilterIface interface{}
