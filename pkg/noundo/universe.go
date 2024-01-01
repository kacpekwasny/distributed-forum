package noundo

func NewUniverse(self HistoryIface, peersNexus_ PeersNexusIface) UniverseIface {
	return &universe{
		self:       self,
		peersNexus: peersNexus_,
	}
}

type universe struct {
	self       HistoryIface
	peersNexus PeersNexusIface
}

func (u *universe) Self() HistoryIface {
	return u.self
}

func (u *universe) Peers() []HistoryIface {
	return u.peersNexus.AlivePeers()
}
