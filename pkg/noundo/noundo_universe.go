package noundo

// UniverseIface is an interface for <object? structure? methods?> giving the ability to retrive
// any HistoryIface, either to the peered ones, or a read-only iface.
// Also should have knowledge of the peers.
type UniverseIface interface {
	Self() HistoryFullIface
	Authenticator() AuthenticatorIface
	Peers() []HistoryPublicIface
}

func NewUniverse(self HistoryFullIface, peersNexus_ PeersNexusIface) UniverseIface {
	return &universe{
		self:       self,
		peersNexus: peersNexus_,
	}
}

type universe struct {
	self       HistoryFullIface
	peersNexus PeersNexusIface
}

func (u *universe) Self() HistoryFullIface {
	return u.self
}

func (u *universe) Peers() []HistoryPublicIface {
	return u.peersNexus.AlivePeers()
}

func (u *universe) Authenticator() AuthenticatorIface {
	return u.self.Authenticator()
}
