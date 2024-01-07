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
	userFUsername string
	reactType     enums.ReactionType

	TimeStampable
}

type Reactionable struct {
	reactions []Reaction
}

type Story struct {
	Postable
	Reactionable
}

type Answer struct {
	PostId Id

	Postable
	Reactionable
}

// TODO remove
type User struct {
	Id       Id
	Username string
}
