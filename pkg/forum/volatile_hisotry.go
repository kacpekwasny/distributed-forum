package forum

import (
	"sort"

	"github.com/kacpekwasny/distributed-forum/pkg/auth"
	"github.com/kacpekwasny/distributed-forum/pkg/utils"
)

// TODO:
// h *HistoryVolatile HistoryIface

type HistoryVolatile struct {
	ages    []AgeIface
	stories map[Id]StoryIface
	answers map[Id]AnswerIface
	auth    auth.Authenticator
}

// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
func (h *HistoryVolatile) CreateAge(owner UserIface, name string) (AgeIface, error) {
	age := &AgeVolatile{
		id:              NewRandId(),
		name:            name,
		ownerUsername:   owner.Username(),
		adminsUsernames: []string{},
	}
	h.ages = append(h.ages, age)
	return age, nil
}

func (h *HistoryVolatile) GetStory(id Id) (StoryIface, error) {
	s, ok := h.stories[id]
	return s, utils.ErrIfNotOk(ok, "id not found")
}

// Get answer from anywhere in
func (h *HistoryVolatile) GetAnswer(id Id) (AnswerIface, error) {
	s, ok := h.answers[id]
	return s, utils.ErrIfNotOk(ok, "id not found")
}

// GetFirst n stories ordered by different atributes, from []ages,
func (h *HistoryVolatile) GetStories(start uint32, end uint32, order OrderIface, filter FilterIface, ages []AgeIface) []StoryIface {
	stories := []StoryIface{}
	for _, story := range h.stories {
		if filter(story) {
			stories = append(stories, story)
		}
	}
	sort.SliceStable(stories, order)
	return stories
}

func (h *HistoryVolatile) GetUser(username string) (UserIface, error) {
	panic("TODO ")
}

func (h *HistoryVolatile) AddUser(login string, username string, password string) (UserIface, error) {
	panic("TODO ")
}
