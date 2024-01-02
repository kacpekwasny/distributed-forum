package noundo

type StorageIface interface {
	CreateAge(owner UserFullIface, ageName string) (AgeIface, error)
	CreateStory(author UserFullIface, story PostableIface, ageName string) (StoryIface, error)
	CreateAnswer(author UserFullIface, answer AnswerIface, parentPostableId Id) (AnswerIface, error)
	ReactPostable(user UserFullIface, postableId Id) error
}
