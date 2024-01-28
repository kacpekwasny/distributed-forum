package noundo

import "github.com/kacpekwasny/noundo/pkg/peer"

type peerAgeWrapper struct {
	*peer.Age
}

func (a *peerAgeWrapper) GetOwner() UserIdentityIface {
	return a.GetOwner()
}
