package noundo

import (
	"log"
	"slices"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

type PeersNexusIface interface {
	AlivePeers() []HistoryPublicIface
	GetHistory(name string) (HistoryPublicIface, error)
	RegisterPeerManager(pm PeerManagerIface)
	UnregisterPeerManager(historyName string)
}

type PeersNexus struct {
	peerManagers    []PeerManagerIface
	historyNamePeer map[string]PeerManagerIface
}

func NewPeersNexus() *PeersNexus {
	return &PeersNexus{
		peerManagers:    []PeerManagerIface{},
		historyNamePeer: make(map[string]PeerManagerIface),
	}
}

func (p *PeersNexus) AlivePeers() []HistoryPublicIface {
	peers := utils.Filter(p.peerManagers, func(t PeerManagerIface) bool {
		return t.PeerAlive() == nil
	})
	return utils.Map(peers, func(p PeerManagerIface) HistoryPublicIface {
		return utils.LeftCallbackIfErr(p.History())(func(err error) {
			log.Println(
				"PeerManager says the connection is alive,",
				"but cannot get History.",
				"The error while trying to get history:", err)
		})
	})
}

func (p *PeersNexus) GetHistory(name string) (HistoryPublicIface, error) {
	histPeer, err := utils.MapGetErr(p.historyNamePeer, name)
	if err != nil {
		return nil, err
	}
	return histPeer.History()
}

func (p *PeersNexus) RegisterPeerManager(pm PeerManagerIface) {
	p.peerManagers = append(p.peerManagers, pm)
	p.historyNamePeer[pm.HistoryURL()] = pm
}

func (p *PeersNexus) UnregisterPeerManager(historyName string) {
	for i, pm := range p.peerManagers {
		if pm.HistoryName() == historyName {
			p.peerManagers = slices.Delete(p.peerManagers, i, i+1)
		}
	}
	delete(p.historyNamePeer, historyName)
}
