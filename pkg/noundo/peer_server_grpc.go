package noundo

import (
	"context"

	"github.com/kacpekwasny/noundo/pkg/peer"
)

type GrpcServer struct {
	hr HistoryReadIface
	peer.UnimplementedHistoryReadServiceServer
}

// gs *GrpcServer HistoryReadServiceServer
func (gs *GrpcServer) GetUser(_ context.Context, ur *peer.UserRequest) (*peer.UserPublicInfo, error) {
	user, err := gs.hr.GetUser(ur.Username)
	if err != nil {
		return nil, err
	}
	return CreatePeerUserPublicInfo(user), nil
}

func (gs *GrpcServer) GetAge(_ context.Context, ar *peer.AgeRequest) (*peer.Age, error) {
	age, err := gs.hr.GetAge(ar.Name)
	if err != nil {
		return nil, err
	}
	owner, err := age.GetOwner()
	if err != nil {
		return nil, err
	}
	return &peer.Age{
		Name:        age.GetName(),
		Description: age.GetDescription(),
		Owner:       CreatePeerUserIdentity(owner),
	}, nil
}

func (gs *GrpcServer) GetAges(_ context.Context, ar *peer.AgesRequest) (*peer.AgeList, error) {
	gs.hr.GetAges(int(ar.Start), int(ar.End))
}

func (gs *GrpcServer) GetStory(_ context.Context, _ *peer.StoryRequest) (*peer.Story, error) {
	panic("not implemented") // TODO: Implement
}

func (gs *GrpcServer) GetStories(_ context.Context, _ *peer.StoriesRequest) (*peer.StoryList, error) {
	panic("not implemented") // TODO: Implement
}

func (gs *GrpcServer) GetAnswer(_ context.Context, _ *peer.AnswerRequest) (*peer.Answer, error) {
	panic("not implemented") // TODO: Implement
}

func (gs *GrpcServer) GetAnswers(_ context.Context, _ *peer.AnswersRequest) (*peer.StoryList, error) {
	panic("not implemented") // TODO: Implement
}

func (gs *GrpcServer) mustEmbedUnimplementedHistoryReadServiceServer() {}
