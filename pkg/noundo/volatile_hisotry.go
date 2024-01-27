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
	users map[string]UserAllIface
}

func NewHistoryVolatile(historyName string) HistoryFullIface {
	usersUsername := make(map[string]UserAllIface)
	usersEmail := make(map[string]UserAllIface)
	return &HistoryVolatile{
		name:    historyName,
		url:     "http://" + historyName,
		ages:    make(map[string]*AgeVolatile),
		stories: make(map[string]*Story),
		answers: make(map[string]*Answer),
		auth:    NewAuthenticator(NewVolatileAuthStorage(historyName, &usersEmail, &usersUsername), DEFAULT_PASS_HASH_COST, []byte(historyName)),
		users:   usersUsername,
	}
}

// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
func (h *HistoryVolatile) CreateAge(owner UserIdentityIface, name string) (AgeIface, error) {
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
	return utils.MapGetErr(h.users, username)
}

func (h *HistoryVolatile) CreateUser(email string, username string, password string) (UserPublicIface, error) {
	r := h.auth.SignUpUser(NewSignUpRequest(email, username, password))
	if r.Ok {
		return h.auth.GetUserByUsername(username).(UserPublicIface), nil
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

func (h *HistoryVolatile) GetStoriesUserJoined(user UserIdentityIface, start int, end int, order OrderIface[StoryIface], filter FilterIface[StoryIface]) ([]StoryIface, error) {
	panic("not implemented") // TODO: Implement
}

func (h *HistoryVolatile) Authenticator() AuthenticatorIface {
	return h.auth
}

// Single age
func (h *HistoryVolatile) GetAge(name string) (AgeIface, error) {
	return utils.MapGetErr[string, *AgeVolatile](h.ages, name)
}

func (h *HistoryVolatile) CreateStory(author UserIdentityIface, ageName string, story StoryContent) (Story, error) {
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
			PostableId: id,
			Author: UserInfo{
				Username:     author.Username(),
				ParentServer: author.ParentServerName(),
				FUsername:    author.FullUsername(),
			},
			Contents: story.Content,
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

// Create an Answer under a post or other Answer
func (h *HistoryVolatile) CreateAnswer(author UserIdentityIface, parentId string, answerContent string) (Answer, error) {
	a, okA := h.answers[parentId]
	s, okS := h.stories[parentId]
	if !(okA || okS) {
		return Answer{}, errors.New("no such parent id")
	}
	id := NewRandId()
	answer := Answer{
		ParentId: parentId,
		Postable: &Postable{
			PostableId:    id,
			Author:        CreateUserInfo(author, h.GetName()),
			Contents:      answerContent,
			TimeStampable: CreateTimeStamp(),
		},
		Reactionable: &Reactionable{
			Reactions: []Reaction{},
		},
		Answerable: &Answerable{
			Answers: []Answer{},
		},
	}
	var answers *[]Answer
	if okA {
		answers = &a.Answerable.Answers
	}
	if okS {
		answers = &s.Answerable.Answers
	}
	*answers = append(*answers, answer)

	h.answers[id] = &answer
	return answer, nil
}

// Get tree of answers, with the specified depth
func (h *HistoryVolatile) GetAnswers(postableId string, start int, end int, depth int, order OrderIface[*Story], filter FilterIface[*Story]) ([]*Story, error) {
	panic("not implemented") // TODO: Implement
}
