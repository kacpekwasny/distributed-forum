package noundo

type PeerManagerGrpc struct {
}

func (pm *PeerManagerGrpc) PeerAlive() error {
	panic("not implemented") // TODO: Implement
}

func (pm *PeerManagerGrpc) History() (HistoryPublicIface, error) {
	panic("not implemented") // TODO: Implement
}

func (pm *PeerManagerGrpc) HistoryURL() string {
	panic("not implemented") // TODO: Implement
}

func (pm *PeerManagerGrpc) HistoryName() string {
	panic("not implemented") // TODO: Implement
}
