package noundo

import (
	ctx "context"

	"github.com/kacpekwasny/noundo/pkg/peer"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func NewHistoryPublicIfaceFromGrpcService(g peer.HistoryReadServiceClient) HistoryPublicIface {
	return &historyPublicGrpcClient{}
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
	return age, err
}

// Get ages ordered and filtered and sliced by the start & end integers
func (h *historyPublicGrpcClient) GetAges(start int, end int, order OrderIface, filter FilterIface) ([]AgeIface, error) {
	panic("not implemented") // TODO: Implement
}

// Get a single story
func (h *historyPublicGrpcClient) GetStory(id string) (Story, error) {
	panic("not implemented") // TODO: Implement
}

// Get `n` stories ordered by different atributes, from []ages,
func (h *historyPublicGrpcClient) GetStories(ageNames []string, start int, end int, order OrderIface, filter FilterIface) ([]*Story, error) {
	panic("not implemented") // TODO: Implement
}

// Get answer from anywhere in
func (h *historyPublicGrpcClient) GetAnswer(id string) (Answer, error) {
	panic("not implemented") // TODO: Implement
}

// Get tree of answers, to the specified postable with the specified depth
func (h *historyPublicGrpcClient) GetAnswers(postableId string, start int, end int, depth int, order OrderIface, filter FilterIface) ([]*Answer, error) {
	panic("not implemented") // TODO: Implement
}

// Create a 'subreddit', but for the sake of naming, it will be called an `Age`
func (h *historyPublicGrpcClient) CreateAge(owner UserIdentityIface, ageName string) (AgeIface, error) {
	panic("not implemented") // TODO: Implement
}

// Create a new story within
func (h *historyPublicGrpcClient) CreateStory(author UserIdentityIface, ageName string, story StoryContentIface) (Story, error) {
	panic("not implemented") // TODO: Implement
}

// Create an Answer under a post or other Answer
func (h *historyPublicGrpcClient) CreateAnswer(author UserIdentityIface, parentId string, answerContent string) (Answer, error) {
	panic("not implemented") // TODO: Implement
}

type peerAgeWrapper struct {
}
