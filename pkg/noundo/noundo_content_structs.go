package noundo

import "github.com/kacpekwasny/noundo/pkg/enums"

// ~~ Structs ~~
type Postable struct {
	Id            string
	UserFUsername string
	Contents      string

	TimeStampable
}
type Story struct {
	Title string

	Postable
	Reactionable
}

type Answer struct {
	PostId Id

	Postable
	Reactionable
}

type Reactionable struct {
	Reactions []Reaction
}

type TimeStampable struct {
	Timestamp int64
}

type Reaction struct {
	UserFUsername string
	ReactType     enums.ReactionType

	TimeStampable
}

// ~~ Structs ~~
type StoryContent struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
