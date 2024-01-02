package noundo

import (
	"sort"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

// TODO:
// h *HistoryVolatile HistoryIface

type HistoryVolatile struct {
	name    string
	url     string
	ages    []AgeIface
	stories map[Id]StoryIface
	answers map[Id]AnswerIface
	auth    AuthenticatorIface

	// email: UserIface
	users map[string]UserAuthIface
}

func NewHistoryVolatile(domain string) HistoryIface {
	usersUsername := make(map[string]UserFullIface)
	usersEmail := make(map[string]UserFullIface)
	return &HistoryVolatile{
		name:    domain,
		url:     "http://" + domain,
		ages:    []AgeIface{},
		stories: make(map[Id]StoryIface),
		answers: make(map[Id]AnswerIface),
		auth:    NewAuthenticator(NewVolatileAuthStorage(&usersEmail, &usersUsername), DEFAULT_PASS_HASH_COST),
	}
}

// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
func (h *HistoryVolatile) CreateAge(owner UserPublicIface, name string) (AgeIface, error) {
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
func (h *HistoryVolatile) GetStories(start int, end int, order OrderIface, filter FilterIface, ages []AgeIface) ([]StoryIface, error) {
	stories := []StoryIface{}
	for _, story := range h.stories {
		if filter(story) {
			stories = append(stories, story)
		}
	}
	sort.SliceStable(stories, order)
	return stories, nil
}

func (h *HistoryVolatile) GetUser(username string) (UserPublicIface, error) {
	return nil, nil
}

func (h *HistoryVolatile) AddUser(email string, username string, password string) (UserPublicIface, error) {
	r := h.auth.SignUpUser(NewSignUpRequest(email, username, password))
	return NewVolatileUser(NewRandId(), email, username, []byte{}, h.GetURL()), utils.ErrIfNotOk(r.Ok, string(r.MsgCode))
}

// Name to be displayed. Ex.: as the value o <a> tag.
func (h *HistoryVolatile) GetName() string {
	return h.name
}

// Get the URL of the History. Ex.: value of href atribute in an <a> tag.
func (h *HistoryVolatile) GetURL() string {
	return h.url
}

func (h *HistoryVolatile) GetAges(start int, end int, order OrderIface, filter FilterIface) ([]AgeIface, error) {
	return h.ages, nil
}
