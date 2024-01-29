package noundo

import (
	ctx "context"

	"github.com/kacpekwasny/noundo/pkg/peer"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func NewHistoryPublicIfaceFromGrpcService(g peer.HistoryReadServiceClient) HistoryPublicIface {
	return &historyPublicGrpcClient{
		g: g,
	}
}

type historyPublicGrpcClient struct {
	g peer.HistoryReadServiceClient
}

// domain name
func (h *historyPublicGrpcClient) GetName() string {
	// todo will fail on error
	return utils.LeftLogRight(h.g.GetName(ctx.Background(), &peer.Empty{})).GetName()
}

// Get the URL of the History. schema + domain
func (h *historyPublicGrpcClient) GetURL() string {
	// todo will fail on error
	return utils.LeftLogRight(h.g.GetURL(ctx.Background(), &peer.Empty{})).GetURL()
}

// Retrive all user info by supplying a username
func (h *historyPublicGrpcClient) GetUser(username string) (UserPublicIface, error) {
	user, err := h.g.GetUser(ctx.Background(), &peer.UserRequest{Username: username})
	return user, err
}

// get a single AgeIface
func (h *historyPublicGrpcClient) GetAge(name string) (AgeIface, error) {
	age, err := h.g.GetAge(ctx.Background(), &peer.AgeRequest{Name: name})
	return &peerAgeWrapper{age}, err
}

// Get ages ordered and filtered and sliced by the start & end integers
func (h *historyPublicGrpcClient) GetAges(start int, end int, order OrderIface, filter FilterIface) ([]AgeIface, error) {
	// todo order & filter
	ages, err := h.g.GetAges(ctx.Background(), &peer.AgesRequest{Start: int32(start), End: int32(end), Order: &peer.Order{}, Filter: &peer.Filter{}})
	return utils.Map(ages.Ages, func(age *peer.Age) AgeIface {
		return &peerAgeWrapper{age}
	}), err
}

// Get a single story
func (h *historyPublicGrpcClient) GetStory(id string) (Story, error) {
	s, err := h.g.GetStory(ctx.Background(), &peer.StoryRequest{Id: id})
	if err != nil {
		return Story{}, nil
	}
	return *CreateNoundoStory(s), nil
}

// Get `n` stories ordered by different atributes, from []ages,
func (h *historyPublicGrpcClient) GetStories(ageNames []string, start int, end int, order OrderIface, filter FilterIface) ([]*Story, error) { // todo change to []StoryIface
	stories, err := h.g.GetStories(ctx.Background(), &peer.StoriesRequest{AgeNames: ageNames, Start: int32(start), End: int32(end), Order: &peer.Order{}, Filter: &peer.Filter{}})
	if err != nil {
		return nil, err
	}
	return utils.Map(stories.Stories, CreateNoundoStory), nil
}

// Get answer from anywhere in
func (h *historyPublicGrpcClient) GetAnswer(id string) (Answer, error) {
	answer, err := h.g.GetAnswer(ctx.Background(), &peer.AnswerRequest{Id: id})
	if err != nil {
		return Answer{}, err
	}
	return *CreateNoundoAnswer(answer), nil
}

// Get tree of answers, to the specified postable with the specified depth
func (h *historyPublicGrpcClient) GetAnswers(postableId string, start int, end int, depth int, order OrderIface, filter FilterIface) ([]*Answer, error) {
	answers, err := h.g.GetAnswers(ctx.Background(), &peer.AnswersRequest{PostableId: postableId, Start: int32(start), End: int32(end), Depth: int32(depth), Order: &peer.Order{}, Filter: &peer.Filter{}})
	if err != nil {
		return []*Answer{}, err
	}
	return utils.Map(answers.Answers, CreateNoundoAnswer), nil
}

// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
func (h *historyPublicGrpcClient) CreateAge(owner UserIdentityIface, ageName string) (AgeIface, error) {
	panic("not implemented") // TODO: Implement
}

// Create a new story within
func (h *historyPublicGrpcClient) CreateStory(author UserIdentityIface, ageName string, story StoryContentIface) (Story, error) {
	newStory, err := h.g.CreateStory(ctx.Background(), &peer.CreateStoryRequest{
		Author: &peer.UserIdentity{
			Username:         author.GetUsername(),
			ParentServerName: author.GetParentServerName(),
		},
		AgeName: ageName,
		StoryContent: &peer.StoryContent{
			Title:   story.GetTitle(),
			Content: story.GetContent(),
		},
	})
	if err != nil {
		return Story{}, err
	}
	return *CreateNoundoStory(newStory), err
}

// Create an Answer under a post or other Answer
func (h *historyPublicGrpcClient) CreateAnswer(author UserIdentityIface, parentId string, answerContent string) (Answer, error) {
	newAnswer, err := h.g.CreateAnswer(ctx.Background(), &peer.CreateAnswerRequest{
		Author: &peer.UserIdentity{
			Username:         author.GetUsername(),
			ParentServerName: author.GetParentServerName(),
		},
		ParentId: parentId,
		AnswerContent: &peer.AnswerContent{
			Content: answerContent,
		},
	})
	if err != nil {
		return Answer{}, err
	}
	return *CreateNoundoAnswer(newAnswer), err
}
