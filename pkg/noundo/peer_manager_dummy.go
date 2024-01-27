package noundo

func NewPeerManagerDummy(his HistoryPublicIface) PeerManagerIface {
	return &PeerManagerDummy{
		his: his,
	}
}

type PeerManagerDummy struct {
	his HistoryPublicIface
}

func (pmd *PeerManagerDummy) PeerAlive() error {
	return nil
}

func (pmd *PeerManagerDummy) History() (HistoryPublicIface, error) {
	return pmd.his, nil
}

func (pmd *PeerManagerDummy) HistoryURL() string {
	return pmd.his.GetURL()
}

func (pmd *PeerManagerDummy) HistoryName() string {
	return pmd.his.GetName()
}
