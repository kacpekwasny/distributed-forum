package noundo

// UniverseIface is an interface for <object? structure? methods?> giving the ability to retrive
// any HistoryIface, either to the peered ones, or a read-only iface.
// Also should have knowledge of the peers.
type UniverseIface interface {
	Self() HistoryFullIface
	Authenticator() AuthenticatorIface
	Peers() []HistoryPublicIface
}

func NewUniverse(self HistoryFullIface, peersNexus_ PeersNexusIface) *Universe {
	return &Universe{
		self:       self,
		peersNexus: peersNexus_,
	}
}

type Universe struct {
	self       HistoryFullIface
	peersNexus PeersNexusIface
}

func (u *Universe) Self() HistoryFullIface {
	return u.self
}

func (u *Universe) Peers() []HistoryPublicIface {
	return u.peersNexus.AlivePeers()
}

func (u *Universe) Authenticator() AuthenticatorIface {
	return u.self.Authenticator()
}

func (u *Universe) GetHistoryByName(name string) (HistoryPublicIface, error) {
	if u.self.GetName() == name {
		return u.self, nil
	}
	return u.peersNexus.GetHistory(name)
}
