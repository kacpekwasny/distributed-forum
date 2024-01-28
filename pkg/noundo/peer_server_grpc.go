package noundo

import (
	"context"

	"github.com/kacpekwasny/noundo/pkg/peer"
	"github.com/kacpekwasny/noundo/pkg/utils"
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
	ages, err := gs.hr.GetAges(int(ar.Start), int(ar.End), ar.Order, ar.Filter)
	if err != nil {
		return nil, err
	}
	return &peer.AgeList{
		Ages: utils.Map(ages, func(a AgeIface) *peer.Age {
			return &peer.Age{
				Name:        a.GetName(),
				Description: a.GetDescription(),
				Owner:       CreatePeerUserIdentity(utils.LeftOr(a.GetOwner())(&volatileUserAuth{})),
				History:     gs.hr.GetName(),
			}
		}),
	}, nil
}

func (gs *GrpcServer) GetStory(_ context.Context, sr *peer.StoryRequest) (*peer.Story, error) {
	s, err := gs.hr.GetStory(sr.GetId())
	if err != nil {
		return nil, err
	}
	return CreatePeerStory(&s), nil
}

func (gs *GrpcServer) GetStories(_ context.Context, sr *peer.StoriesRequest) (*peer.StoryList, error) {
	stories, err := gs.hr.GetStories(sr.AgeNames, int(sr.Start), int(sr.End), sr.Order, sr.Filter)
	if err != nil {
		return nil, err
	}
	return &peer.StoryList{
		Stories: utils.Map(stories, CreatePeerStory),
	}, nil
}

func (gs *GrpcServer) GetAnswer(_ context.Context, ar *peer.AnswerRequest) (*peer.Answer, error) {
	ans, err := gs.hr.GetAnswer(ar.Id)
	if err != nil {
		return nil, err
	}
	return CreatePeerAnswer(&ans), nil
}

func (gs *GrpcServer) GetAnswers(_ context.Context, as *peer.AnswersRequest) (*peer.AnswerList, error) {
	answers, err := gs.hr.GetAnswers(as.PostableId, int(as.Start), int(as.End), int(as.Depth), as.Order, as.Filter)
	if err != nil {
		return nil, err
	}
	return &peer.AnswerList{
		Answers: utils.Map(answers, CreatePeerAnswer),
	}, nil
}

func (gs *GrpcServer) mustEmbedUnimplementedHistoryReadServiceServer() {}
