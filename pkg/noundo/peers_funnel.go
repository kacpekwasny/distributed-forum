package noundo

import (
	"log"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

type PeersFunnelIface interface {
	AlivePeers() []HistoryIface
	GetHistory(name string) (HistoryIface, error)
}

type peersFunnel struct {
	peerManagers    []PeerManagerIface
	historyNamePeer map[string]PeerManagerIface
}

func (p *peersFunnel) AlivePeers() []HistoryIface {
	peers := utils.Filter(p.peerManagers, func(t PeerManagerIface) bool {
		return t.PeerAlive() == nil
	})
	return utils.Map(peers, func(p PeerManagerIface) HistoryIface {
		return utils.LeftCallbackIfErr(p.History())(func(err error) {
			log.Println(
				"PeerManager says the connection is alive,",
				"but cannot get History.",
				"Error while trying to get history:", err)
		})
	})
}

func (p *peersFunnel) GetHistory(name string) (HistoryIface, error) {
	histPeer, err := utils.MapGetErr(p.historyNamePeer, name)
	if err != nil {
		return nil, err
	}
	return histPeer.History()
}
