package noundo

type PeerManagerDummy struct {
	his HistoryIface
}

func (pmd *PeerManagerDummy) PeerAlive() error {
	return nil
}

func (pmd *PeerManagerDummy) History() (HistoryIface, error) {
	return pmd.his, nil
}

func (pmd *PeerManagerDummy) HistoryURL() string {
	return pmd.his.GetURL()
}

func (pmd *PeerManagerDummy) HistoryName() string {
	return pmd.his.GetName()
}
