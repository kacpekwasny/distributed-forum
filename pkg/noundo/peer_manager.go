package noundo

// Peer Manager is responsible for keeping and managing a connection to a single Peer (History)
type PeerManagerIface interface {
	PeerAlive() error
	History() (HistoryIface, error)
	HistoryURL() string
	HistoryName() string
}

type PeerManager struct {
}
