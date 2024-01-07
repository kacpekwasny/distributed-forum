package noundo

import (
	"time"

	"github.com/kacpekwasny/noundo/pkg/enums"
)

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

func NewStoryVolatile(authorFullUsername string, id Id, content string) StoryIface {
	return &Story{
		Postable: Postable{
			id:            id,
			userFUsername: authorFullUsername,
			contents:      content,
			TimeStampable: TimeStampable{
				timestamp: uint64(time.Now().Unix()),
			},
		},
	}
}

func (s *Story) Id() Id {
	return s.Postable.id
}

func (s *Story) Content() string {
	return s.contents
}

func (s *Story) AuthorFUsername() string {
	return s.userFUsername
}

func (s *Story) ReactionStats() (map[enums.ReactionType]int, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Story) Reactions() ([]ReactionIface, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Story) React(user UserFullIface, reaction ReactionIface) error {
	panic("not implemented") // TODO: Implement
}

func (s *Story) AddAnswer(author UserFullIface, answerable AnswerableIface, answer AnswerIface) (AnswerIface, error) {
	panic("not implemented") // TODO: Implement
}

func (s *Story) Answers(start int, end int, depth int, order OrderIface[AnswerableIface], filter FilterIface[AnswerableIface], ages []AgeIface) ([]AnswerIface, error) {
	panic("not implemented") // TODO: Implement
}
