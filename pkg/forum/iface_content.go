package forum

import "github.com/kacpekwasny/distributed-forum/pkg/enums"

type Id uint64

// History is the name for the whole server, that contains all Ages and All Stories
type HistoryIface interface {
	// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
	CreateAge(owner UserIface, name string) (AgeIface, error)

	// Create a Story written by an Author in a certain Age
	AddStory(author UserIface, age AgeIface, story StoryIface) (StoryIface, error)

	// Add a Comment to a Story
	AddComment(author UserIface, story StoryIface, comment CommentIface) (CommentIface, error)

	// Add a reaction to a comment
	ReactStory(react ReactionIface)

	// Add reaction to a comment (My assumption is Stories and Comments will be in different
	// tables for optimization purposes)
	ReactComment(react ReactionIface)

	// TODO:
	// GetFirst n stories ordered by different atributes, from []ages,
	// GetFirst n comments ordered by different atributes, and depth X, From Story, or from other comment
	// GetAges that user joined,
	// GetStories first n stories of user ordered by (maybe merge with the first method???)
	// GetComments first n comments of user ordered by
}

// Are all those other interfaces needed??? TODO
type AgeIface interface {
	GetId() (Id, error)
	GetName() (string, error)
	SetName() (string, error)
	GetOwner() (UserIface, error)
	GetAdmins() ([]UserIface, error)
}

type PostableIface interface {
	Id() Id
	Body() string
	AuthorId() Id
	ReactionableIface
}

type StoryIface interface {
	PostableIface
	Comments() ([]CommentIface, error)
}

type CommentIface interface {
	PostableIface
	Answers() ([]CommentIface, error)
}

type UserIface interface {
	Username() string
	Login() string
}

type ReactionableIface interface {
	ReactionStats() (map[enums.ReactionType]uint32, error)
	Reactions() ([]ReactionIface, error)
	React(UserIface, ReactionIface)
}

type ReactionIface interface {
	Id() Id
	AuthorId() Id
	Enum() enums.ReactionType
	Timestamp() uint64
}
