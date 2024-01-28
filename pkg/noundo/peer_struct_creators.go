package noundo

import "github.com/kacpekwasny/noundo/pkg/peer"

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
