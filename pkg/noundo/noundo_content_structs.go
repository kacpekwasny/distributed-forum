package noundo

import "github.com/kacpekwasny/noundo/pkg/enums"

// ~~ Structs ~~
type TimeStampable struct {
	timestamp uint64
}

type Postable struct {
	id            Id
	userFUsername string
	contents      string

	TimeStampable
}

type Reaction struct {
	UserFUsername string
	ReactType     enums.ReactionType

	TimeStampable
}

type Reactionable struct {
	reactions []Reaction
}

type Story struct {
	Title string
	Postable
	Reactionable
}

type StoryContent struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Answer struct {
	PostId Id

	Postable
	Reactionable
}
