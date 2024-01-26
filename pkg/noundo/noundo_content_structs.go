package noundo

import "github.com/kacpekwasny/noundo/pkg/enums"

// ~~ Structs ~~

type UserInfo struct {
	Username       string
	FUsername      string
	ParentServer   string
	UserProfileURL string
}

type Postable struct {
	PostableId string
	Author     UserInfo
	Contents   string

	TimeStampable
}
type Story struct {
	Title       string
	AgeName     string
	HistoryName string

	Postable // TODO
	Reactionable
	Answerable
}

type Answer struct {
	ParentId string
	*Postable
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
