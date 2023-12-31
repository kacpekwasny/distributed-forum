package noundo

// Peer Manager is responsible for keeping and managing a connection to a single Peer (History)
type PeerManagerIface interface {
	History() (HistoryIface, error)
	PeerAlive() error
}
