package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/peer"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func CreatePeerUserIdentity(u UserIdentityIface) *peer.UserIdentity {
	return &peer.UserIdentity{
		Username:     u.Username(),
		ParentServer: u.ParentServerName(),
	}
}

func CreatePeerUserPublicInfo(u UserPublicIface) *peer.UserPublicInfo {
	return &peer.UserPublicInfo{
		User:      CreatePeerUserIdentity(u),
		BirthDate: u.AccountBirthDate(),
		AboutMe:   u.AboutMe(),
	}
}

func CreatePeerStory(s *Story) *peer.Story {
	return &peer.Story{
		Title:       s.Title,
		AgeName:     s.AgeName,
		HistoryName: s.AgeName,
		Postable: &peer.Postable{
			Id:        s.PostableId,
			Author:    CreatePeerUserIdentity(&s.Postable.Author),
			Content:   s.Content(),
			Timestamp: s.Timestamp,
		},
		Answerable: &peer.Answerable{},
	}
}

func CreatePeerAnswer(a *Answer) *peer.Answer {
	return &peer.Answer{
		ParentId: a.ParentId,
		Postable: &peer.Postable{
			Id:        a.PostableId,
			Author:    CreatePeerUserIdentity(&a.Author),
			Content:   a.Contents,
			Timestamp: a.Timestamp,
		},
		Answerable: &peer.Answerable{
			Answers: utils.Map(a.Answerable.Answers, func(a Answer) *peer.Answer {
				return CreatePeerAnswer(&a)
			}),
		},
	}
}
