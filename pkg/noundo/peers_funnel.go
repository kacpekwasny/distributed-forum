package noundo

import (
	"log"
	"slices"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

type PeersFunnelIface interface {
	AlivePeers() []HistoryIface
	GetHistory(name string) (HistoryIface, error)
	RegisterPeerManager(pm PeerManagerIface)
	UnregisterPeerManager(historyName string)
}

type PeersFunnel struct {
	peerManagers    []PeerManagerIface
	historyNamePeer map[string]PeerManagerIface
}

func NewPeersFunnel() *PeersFunnel {
	return &PeersFunnel{
		peerManagers:    []PeerManagerIface{},
		historyNamePeer: make(map[string]PeerManagerIface),
	}
}

func (p *PeersFunnel) AlivePeers() []HistoryIface {
	peers := utils.Filter(p.peerManagers, func(t PeerManagerIface) bool {
		return t.PeerAlive() == nil
	})
	return utils.Map(peers, func(p PeerManagerIface) HistoryIface {
		return utils.LeftCallbackIfErr(p.History())(func(err error) {
			log.Println(
				"PeerManager says the connection is alive,",
				"but cannot get History.",
				"The error while trying to get history:", err)
		})
	})
}

func (p *PeersFunnel) GetHistory(name string) (HistoryIface, error) {
	histPeer, err := utils.MapGetErr(p.historyNamePeer, name)
	if err != nil {
		return nil, err
	}
	return histPeer.History()
}

func (p *PeersFunnel) RegisterPeerManager(pm PeerManagerIface) {
	p.peerManagers = append(p.peerManagers, pm)
	p.historyNamePeer[pm.HistoryURL()] = pm
}

func (p *PeersFunnel) UnregisterPeerManager(historyName string) {
	for i, pm := range p.peerManagers {
		if pm.HistoryName() == historyName {
			p.peerManagers = slices.Delete(p.peerManagers, i, i+1)
		}
	}
	delete(p.historyNamePeer, historyName)
}
