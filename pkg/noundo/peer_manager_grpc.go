package noundo

import (
	"github.com/kacpekwasny/noundo/pkg/peer"
	"google.golang.org/grpc"
)

type PeerManagerGrpc struct {
	alive   bool
	conn    *grpc.ClientConn
	history peer.HistoryReadServiceClient
}

func NewGrpcPeerManager(serverAddr string) PeerManagerIface {
	// todo Dial in a goroutine, that tries connection every X minutes,
	conn, err := grpc.Dial(serverAddr)
	return &PeerManagerGrpc{
		alive:   err == nil,
		conn:    conn,
		history: peer.NewHistoryReadServiceClient(conn),
	}
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
