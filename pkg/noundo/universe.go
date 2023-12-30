package noundo

func NewUniverse(self HistoryIface) UniverseIface {
	return &universe{
		self: self,
	}
}

type universe struct {
	self HistoryIface
}

func (u *universe) Self() HistoryIface {
	return u.self
}

func (u *universe) Peers() []HistoryIface {
	panic("not implemented") // TODO: Implement
}
