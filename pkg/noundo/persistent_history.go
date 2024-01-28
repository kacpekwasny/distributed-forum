package noundo

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kacpekwasny/noundo/pkg/utils"
	"golang.org/x/exp/maps"
)

// TODO:

type HistoryPersistent struct {
	name   string
	url    string
	auth   AuthenticatorIface
	dbPool *pgxpool.Pool

	// email: UserIface
	users map[string]UserAllIface
}

func NewHistoryPersistent(historyName string) HistoryFullIface {
	// TODO: use .env file for database connection
	dbpool, err := pgxpool.New(context.Background(), "postgresql://localhost:5432/distributed_forum")
	usersUsername := make(map[string]UserAllIface)
	usersEmail := make(map[string]UserAllIface)
	return &HistoryPersistent{
		name:   historyName,
		url:    "http://" + historyName,
		dbPool: dbpool,
		auth:   NewAuthenticator(NewPersistentAuthStorage(historyName, &usersEmail, &usersUsername), DEFAULT_PASS_HASH_COST, []byte(historyName)),
		users:  usersUsername,
	}
}

// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
func (h *HistoryPersistent) CreateAge(owner UserIdentityIface, name string) (AgeIface, error) {

	//var id, err = h.dbPool.Exec(context.Background(), "INSERT INTO ages ( name, owner_id) VALUES ($1, $2) RETURNING id", name, owner.Id())
	//age := &AgePersistent{id.String(), name, owner.Id(), ""}
	//return age, err

	//the bellow code INSERTS into the db and returns the created age
	var id string
	err := h.dbPool.QueryRow(context.Background(), "INSERT INTO ages ( name, owner_id) VALUES ($1, $2) RETURNING id", name, owner.Id()).Scan(&id)
	age := &AgePersistent{id, name, owner.Id(), ""}
	return age, err
}

func (h *HistoryPersistent) GetStory(id int) (Story, error) {
	// queries for story by id
	var age_id, owner_id int
	var title, content string
	var timestamp int64
	err := h.dbPool.QueryRow(context.Background(), "SELECT  title, asset_url, age_id, owner_id, content, timestamp  FROM stories WHERE id = $1", id).Scan(&title, &age_id, &owner_id, &content, &timestamp)

	//get the answers
	var answers []Answer
	rows, err := h.dbPool.Query(context.Background(), "SELECT  id, parent_id, owner_id, content, score, timestamp  FROM answers WHERE story_id = $1", id)
	for rows.Next() {
		var id, parent_id, owner_id, score int
		var content string
		var timestamp int64
		err = rows.Scan(&id, &parent_id, &owner_id, &content, &score, &timestamp)
		if err != nil {
			return Story{}, err
		}
		answers = append(answers, Answer{
			ParentId: -1,
			Postable: Postable{PostableId: id, AuthorId: owner_id, Contents: content, TimeStampable: TimeStampable{Timestamp: timestamp}},
		})
	}
	//if there is no story with this id
	if err != nil {
		return Story{}, err
	}

	// creates story object
	story := Story{Title: title, AgeId: age_id, Postable: Postable{
		PostableId:    id,
		AuthorId:      owner_id,
		Contents:      content,
		TimeStampable: TimeStampable{Timestamp: timestamp},
	}}
	return story, err
}

// Get answer from anywhere in
func (h *HistoryPersistent) GetAnswer(id int) (Answer, error) {
	// queries for answer by id
	var parent_id, owner_id int
	var content string
	var timestamp int64
	err := h.dbPool.QueryRow(context.Background(), "SELECT  parent_id, owner_id, content, timestamp  FROM answers WHERE id = $1", id).Scan(&parent_id, &owner_id, &content, &timestamp)
	//create answer object
	answer := Answer{ParentId: -1, Postable: Postable{PostableId: id, AuthorId: owner_id, Contents: content, TimeStampable: TimeStampable{Timestamp: timestamp}}}

	return answer, err
}

// GetFirst n stories ordered by different atributes, from []ages,
func (h *HistoryPersistent) GetStories(ageNames []string, start int, end int, order OrderIface, filter FilterIface) ([]Story, error) {
	// queries for stories by age name
	var stories []Story
	rows, err := h.dbPool.Query(context.Background(), "SELECT  id, title, asset_url, age_id, owner_id, content, timestamp  FROM stories WHERE age_id = $1", ageNames[0])
	for rows.Next() {
		var id, age_id, owner_id int
		var title, content string
		var timestamp int64
		err = rows.Scan(&id, &title, &age_id, &owner_id, &content, &timestamp)
		if err != nil {
			return nil, err
		}
		stories = append(stories, Story{Title: title, AgeId: age_id, Postable: Postable{
			PostableId:    id,
			AuthorId:      owner_id,
			Contents:      content,
			TimeStampable: TimeStampable{Timestamp: timestamp},
		}})
	}
	// if order != nil {
	// 	sort.SliceStable(stories, func(i, j int) bool {
	// 			return order.Less(stories[i], stories[j])
	// 	})
	// }
	return stories, nil
}

func (h *HistoryPersistent) GetUser(username string) (UserPublicIface, error) {
	// queries for user by username
	var id int
	var email, parent_server string
	err := h.dbPool.QueryRow(context.Background(), "SELECT  id, email, parentServer FROM users WHERE username = $1", username).Scan(&id, &email, &parent_server)
	//create user object
	user := User{username: username, email: email, parentServerName: parent_server, id: id}

	return utils.MapGetErr(h.users, username)
}

func (h *HistoryPersistent) CreateUser(email string, username string, password string) (UserPublicIface, error) {
	r := h.auth.SignUpUser(NewSignUpRequest(email, username, password))
	if r.Ok {
		return h.auth.GetUserByUsername(username).(UserPublicIface), nil
	}
	return nil, errors.New(string(r.MsgCode))
}

// Name to be displayed. Ex.: as the value o <a> tag.
func (h *HistoryPersistent) GetName() string {
	return h.name
}

// Get the URL of the History. Ex.: value of href atribute in an <a> tag.
func (h *HistoryPersistent) GetURL() string {
	return h.url
}

func (h *HistoryPersistent) GetAges(start int, end int, order OrderIface, filter FilterIface) ([]AgeIface, error) {
	ages := []AgeIface{}
	for _, a := range maps.Values(h.ages) {
		ages = append(ages, a)
	}
	return ages, nil
}

// func (h *HistoryVolatile) GetStoriesUserJoined(user UserIdentityIface, start int, end int, order OrderIface, filter FilterIface) ([]StoryIface, error) {
// 	panic("not implemented") // TODO: Implement
// }

func (h *HistoryPersistent) Authenticator() AuthenticatorIface {
	return h.auth
}

// Single age
func (h *HistoryPersistent) GetAge(name string) (AgeIface, error) {
	return utils.MapGetErr[string, *AgePersistent](h.ages, name)
}

func (h *HistoryPersistent) CreateStory(author UserIdentityIface, ageName string, story StoryContent) (Story, error) {
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
				username:     author.Username(),
				parentServer: author.ParentServerName(),
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
func (h *HistoryPersistent) CreateAnswer(author UserIdentityIface, parentId string, answerContent string) (Answer, error) {
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
func (h *HistoryPersistent) GetAnswers(postableId string, start int, end int, depth int, order OrderIface, filter FilterIface) ([]*Answer, error) {
	panic("not implemented") // TODO: Implement
}
