package noundo

func NewUniverse(self HistoryIface, peersFunnel_ PeersFunnelIface) UniverseIface {
	return &universe{
		self:        self,
		peersFunnel: peersFunnel_,
	}
}

type universe struct {
	self        HistoryIface
	peersFunnel PeersFunnelIface
}

func (u *universe) Self() HistoryIface {
	return u.self
}

func (u *universe) Peers() []HistoryIface {
	return u.peersFunnel.AlivePeers()
}
