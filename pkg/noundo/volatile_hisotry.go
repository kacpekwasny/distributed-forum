package noundo

import (
	"errors"
	"sort"

	"github.com/kacpekwasny/noundo/pkg/utils"
	"golang.org/x/exp/maps"
)

// TODO:
// h *HistoryVolatile HistoryFullIface

type HistoryVolatile struct {
	name    string
	url     string
	ages    map[string]AgeIface
	stories map[Id]StoryIface
	answers map[Id]AnswerIface
	auth    AuthenticatorIface

	// email: UserIface
	users map[string]UserFullIface
}

func NewHistoryVolatile(domain string) HistoryFullIface {
	usersUsername := make(map[string]UserFullIface)
	usersEmail := make(map[string]UserFullIface)
	return &HistoryVolatile{
		name:    domain,
		url:     "http://" + domain,
		ages:    make(map[string]AgeIface),
		stories: make(map[Id]StoryIface),
		answers: make(map[Id]AnswerIface),
		auth:    NewAuthenticator(NewVolatileAuthStorage(&usersEmail, &usersUsername), DEFAULT_PASS_HASH_COST),
		users:   usersUsername,
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
	h.ages[name] = age
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
func (h *HistoryVolatile) GetStories(ageNames []string, start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) ([]StoryIface, error) {
	stories := []StoryIface{}
	for _, story := range h.stories {
		if filter == nil || filter.Keep(story) {
			stories = append(stories, story)
		}
	}
	if order != nil {
		sort.SliceStable(stories, func(i, j int) bool {
			return order.Less(stories[i], stories[j])
		})
	}
	return stories, nil
}

func (h *HistoryVolatile) GetUser(username string) (UserPublicIface, error) {
	return utils.MapGetErr[string, UserFullIface](h.users, username)
}

func (h *HistoryVolatile) CreateUser(email string, username string, password string) (UserPublicIface, error) {
	r := h.auth.SignUpUser(NewSignUpRequest(email, username, password))
	if r.Ok {
		u := (h.auth.GetUserByEmail(username)).(*User)
		u.parentServerName = h.GetName()
		return u, nil
	}
	return nil, errors.New(string(r.MsgCode))
}

// Name to be displayed. Ex.: as the value o <a> tag.
func (h *HistoryVolatile) GetName() string {
	return h.name
}

// Get the URL of the History. Ex.: value of href atribute in an <a> tag.
func (h *HistoryVolatile) GetURL() string {
	return h.url
}

func (h *HistoryVolatile) GetAges(start int, end int, order OrderIface[AgeIface], filter FilterIface[AgeIface]) ([]AgeIface, error) {
	return maps.Values(h.ages), nil
}

func (h *HistoryVolatile) GetStoriesUserJoined(user UserPublicIface, start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) ([]StoryIface, error) {
	panic("not implemented") // TODO: Implement
}

func (h *HistoryVolatile) Authenticator() AuthenticatorIface {
	return h.auth
}

// Single age
func (h *HistoryVolatile) GetAge(name string) (AgeIface, error) {
	panic("not implemented") // TODO: Implement
}

func (h *HistoryVolatile) CreateStory(ageName string, author UserPublicIface, story StoryContent) (StoryIface, error) {
	_, exists := h.ages[ageName]
	if !exists {
		return nil, errors.New("age with ageName: '" + ageName + "' doesnt exist")
	}

	id := NewRandId()
	storyInternal := NewStoryVolatile(author.FullUsername(), id, story.Content)
	h.stories[id] = storyInternal
	return storyInternal, nil
}
