package noundo

type CreateStory struct {
	authorFUsername string
	content         string
}

func NewCreateStory(authorFUsername string, content string) CreateStoryIface {
	return &CreateStory{
		authorFUsername: authorFUsername,
		content:         content,
	}
}

type CreateStoryIface interface {
	AuthorFUsername() string
	Content() string
}

func (s *CreateStory) AuthorFUsername() string {
	return s.authorFUsername
}

func (s *CreateStory) Content() string {
	return s.content
}

// func NewStoryVolatile(authorFullUsername string, id Id, content string) StoryIface {
// 	return &Story{
// 		Postable: Postable{
// 			Id:            id,
// 			UserFUsername: authorFullUsername,
// 			Contents:      content,
// 			TimeStampable: TimeStampable{
// 				Timestamp: uint64(time.Now().Unix()),
// 			},
// 		},
// 	}
// }

func (s *Story) Id() int {
	return s.Postable.PostableId
}

func (s *Story) Content() string {
	return s.Contents
}

func (s *Story) AuthorFUsername() string {
	return s.Author.FUsername
}

// func (s *Story) ReactionStats() (map[enums.ReactionType]int, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (s *Story) Reactions() ([]ReactionIface, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (s *Story) React(user UserIdentityIface, reaction ReactionIface) error {
// 	panic("not implemented") // TODO: Implement
// }

// func (s *Story) AddAnswer(author UserIdentityIface, answerable AnswerableIface, answer AnswerIface) (AnswerIface, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (s *Story) Answers(start int, end int, depth int, order OrderIface[AnswerableIface], filter FilterIface[AnswerableIface], ages []AgeIface) ([]AnswerIface, error) {
// 	panic("not implemented") // TODO: Implement
// }
