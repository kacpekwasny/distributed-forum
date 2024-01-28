package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/peer"
	"google.golang.org/grpc"
)

type PeerManagerGrpc struct {
	alive   error
	conn    *grpc.ClientConn
	history peer.HistoryReadServiceClient
}

func NewPeerManagerGrpc(serverAddr string) PeerManagerIface {
	// todo Dial in a goroutine, that tries connection every X minutes,
	conn, err := grpc.Dial(serverAddr)
	return &PeerManagerGrpc{
		alive:   err,
		conn:    conn,
		history: peer.NewHistoryReadServiceClient(conn),
	}
}

func (pm *PeerManagerGrpc) PeerAlive() error {
	return pm.alive
}

func (pm *PeerManagerGrpc) History() (HistoryPublicIface, error) {
	return pm.history, pm.alive
}

func (pm *PeerManagerGrpc) HistoryURL() string {
	panic("not implemented") // TODO: Implement
}

func (pm *PeerManagerGrpc) HistoryName() string {
	panic("not implemented") // TODO: Implement
}
