package enums

type ReactionType uint64

const (
	LIKE = ReactionType(iota)
	BULLSHIT
	DISLIKE
)
