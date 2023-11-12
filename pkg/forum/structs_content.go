package forum

import "github.com/kacpekwasny/distributed-forum/pkg/enums"

// ~~ Structs ~~
type TimeStampable struct {
	Timestamp uint64
}

type Postable struct {
	Id       Id
	UserId   Id
	Contents string

	TimeStampable
}

type Reaction struct {
	UserId    Id
	ReactType enums.ReactionType

	TimeStampable
}

type Reactionable struct {
	Reactions []Reaction
}

type Story struct {
	Postable
	Reactionable
}

type Comment struct {
	PostId Id

	Postable
	Reactionable
}

type User struct {
	Id       Id
	Username string
}
