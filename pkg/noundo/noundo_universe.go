package noundo

// UniverseIface is an interface for <object? structure? methods?> giving the ability to retrive
// any HistoryIface, either to the peered ones, or a read-only iface.
// Also should have knowledge of the peers.
type UniverseIface interface {
	Self() HistoryIface
	Peers() []HistoryIface
}

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
