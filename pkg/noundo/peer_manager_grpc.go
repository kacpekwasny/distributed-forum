package noundo

import (
	"log/slog"

	"github.com/kacpekwasny/noundo/pkg/peer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type PeerManagerGrpc struct {
	alive      error
	conn       *grpc.ClientConn
	history    HistoryPublicIface
	serverAddr string
}

func NewPeerManagerGrpc(serverAddr string) PeerManagerIface {
	// todo Dial in a goroutine, that tries connection every X minutes,
	conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		slog.Error("grpc.Dial to", "serverAddr", serverAddr, "err", err)
	}
	return &PeerManagerGrpc{
		alive:      err,
		conn:       conn,
		history:    NewHistoryPublicIfaceFromGrpcService(peer.NewHistoryReadServiceClient(conn)),
		serverAddr: serverAddr,
	}
}

func (pm *PeerManagerGrpc) PeerAlive() error {
	return pm.alive
}

func (pm *PeerManagerGrpc) History() (HistoryPublicIface, error) {
	return pm.history, pm.alive
}

func (pm *PeerManagerGrpc) HistoryURL() string {
	return pm.history.GetURL()
}

func (pm *PeerManagerGrpc) HistoryName() string {
	return pm.history.GetName()
}
