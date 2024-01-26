package noundo

import (
	"errors"
	"sort"
	"time"

	"github.com/kacpekwasny/noundo/pkg/utils"
	"golang.org/x/exp/maps"
)

// TODO:
// h *HistoryVolatile HistoryFullIface

type HistoryVolatile struct {
	name    string
	url     string
	ages    map[string]*AgeVolatile
	stories map[string]*Story
	answers map[string]*Answer
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
		ages:    make(map[string]*AgeVolatile),
		stories: make(map[string]*Story),
		answers: make(map[string]*Answer),
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

func (h *HistoryVolatile) GetStory(id string) (Story, error) {
	s, ok := h.stories[id]
	if ok {
		return *s, nil
	}
	return Story{}, utils.ErrIfNotOk(ok, "id not found")
}

// Get answer from anywhere in
func (h *HistoryVolatile) GetAnswer(id string) (Answer, error) {
	s, ok := h.answers[id]
	return *s, utils.ErrIfNotOk(ok, "id not found")
}

// GetFirst n stories ordered by different atributes, from []ages,
func (h *HistoryVolatile) GetStories(ageNames []string, start int, end int, order OrderIface[*Story], filter FilterIface[*Story]) ([]*Story, error) {
	stories := []*Story{}
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
	ages := []AgeIface{}
	for _, a := range maps.Values(h.ages) {
		ages = append(ages, a)
	}
	return ages, nil
}

func (h *HistoryVolatile) GetStoriesUserJoined(user UserPublicIface, start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) ([]StoryIface, error) {
	panic("not implemented") // TODO: Implement
}

func (h *HistoryVolatile) Authenticator() AuthenticatorIface {
	return h.auth
}

// Single age
func (h *HistoryVolatile) GetAge(name string) (AgeIface, error) {
	return utils.MapGetErr[string, *AgeVolatile](h.ages, name)
}

func (h *HistoryVolatile) CreateStory(ageName string, author UserPublicIface, story StoryContent) (Story, error) {
	_, exists := h.ages[ageName]
	if !exists {
		return Story{}, errors.New("age with ageName: '" + ageName + "' doesnt exist")
	}

	id := NewRandId()

	storyInternal := Story{
		Title:       story.Title,
		AgeName:     ageName,
		HistoryName: h.name,
		Postable: Postable{
			Id:            id,
			UserFUsername: author.FullUsername(),
			Contents:      story.Content,
			TimeStampable: TimeStampable{
				Timestamp: time.Now().Unix(),
			},
		},
		Reactionable: Reactionable{
			Reactions: []Reaction{},
		},
	}
	h.stories[id] = &storyInternal
	return storyInternal, nil
}
