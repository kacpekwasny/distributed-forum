package noundo

import "github.com/kacpekwasny/noundo/pkg/enums"

// ~~ Structs ~~

type UserInfo struct {
	username       string
	FUsername      string
	parentServer   string
	UserProfileURL string
}

type Postable struct {
	PostableId int
	AuthorId   int
	Contents   string

	TimeStampable
}
type Story struct {
	Title string
	AgeId int

	Postable // TODO
	Reactionable
	Answerable
}

type Answer struct {
	ParentId int
	Postable
	*Reactionable
	*Answerable
}

type Answerable struct {
	Answers []Answer
}

type Reaction struct {
	UserFUsername string
	ReactType     enums.ReactionType

	TimeStampable
}

type Reactionable struct {
	Reactions []Reaction
}

type TimeStampable struct {
	Timestamp int64
}

// ~~ Structs ~~
type StoryContent struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type AnswerContent struct {
	Content string `json:"content"`
}
