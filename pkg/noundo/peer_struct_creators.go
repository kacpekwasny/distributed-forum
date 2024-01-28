package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/peer"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func CreatePeerUserIdentity(u UserIdentityIface) *peer.UserIdentity {
	return &peer.UserIdentity{
		Username:         u.GetUsername(),
		ParentServerName: u.GetParentServerName(),
	}
}

func CreatePeerUserPublicInfo(u UserPublicIface) *peer.UserPublicInfo {
	return &peer.UserPublicInfo{
		User:             CreatePeerUserIdentity(u),
		AccountBirthDate: u.GetAccountBirthDate(),
		AboutMe:          u.GetAboutMe(),
	}
}

func CreatePeerStory(historyName string) func(s *Story) *peer.Story {
	return func(s *Story) *peer.Story {
		return &peer.Story{
			Title:       s.Title,
			AgeName:     s.AgeName,
			HistoryName: historyName,
			Postable: &peer.Postable{
				Id:        s.PostableId,
				Author:    CreatePeerUserIdentity(&s.Postable.Author),
				Content:   s.Content(),
				Timestamp: s.Timestamp,
			},
			Answerable: &peer.Answerable{},
		}
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

func CreatePeerAge(a AgeIface) *peer.Age {
	return &peer.Age{
		Name:        a.GetName(),
		Description: a.GetDescription(),
		Owner:       CreatePeerUserIdentity(a.GetOwner()),
	}
}

func CreateNoundoStory(s *peer.Story) *Story {
	return &Story{
		Title:       s.GetTitle(),
		AgeName:     s.GetAgeName(),
		HistoryName: s.GetHistoryName(),
		Postable: Postable{
			PostableId:    s.Postable.Id,
			Author:        CreateNoundoUser(s.Postable.Author),
			Contents:      s.Postable.Content,
			TimeStampable: TimeStampable{s.Postable.Timestamp},
		},
	}
}

func CreateNoundoUser(u *peer.UserIdentity) UserInfo {
	return UserInfo{
		username:       u.GetUsername(),
		parentServer:   u.GetParentServerName(),
		FUsername:      u.GetFUsername(),
		UserProfileURL: "todo, or remove field",
	}
}

func CreateNoundoAnswer(a *peer.Answer) *Answer {
	return &Answer{
		ParentId: a.ParentId,
		Postable: &Postable{
			PostableId:    a.Postable.Id,
			Author:        CreateNoundoUser(a.Postable.Author),
			Contents:      a.Postable.Content,
			TimeStampable: TimeStampable{a.Postable.Timestamp},
		},
		Answerable: &Answerable{
			Answers: utils.Map(a.Answerable.Answers, func(a *peer.Answer) Answer { return *CreateNoundoAnswer(a) }),
		},
	}
}
