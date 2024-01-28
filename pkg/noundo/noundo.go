package noundo

import (
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/noundo/pkg/peer"
	"google.golang.org/grpc"
)

type NoUndo struct {
	// Stores interfaces to Peers, and to own history
	uni *Universe

	// net.Listener, and grpc.Server - needed for the gRPC server,
	lis  net.Listener
	grpc *grpc.Server

	// Http Server (using the struct, to imlpement graceful shutdown)
	http *http.Server
	r    *mux.Router
}

func NewNoUndo(uni *Universe) *NoUndo {
	n := &NoUndo{
		uni:  uni,
		grpc: grpc.NewServer(),
	}
	n.setupRouter()
	return n
}

func (n *NoUndo) SetupListen(httpAddr string, grpcAddr string) error {
	n.http = &http.Server{
		Addr:         httpAddr,
		Handler:      n.r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	lis, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		slog.Error("could not listen on tcp/"+grpcAddr+". error:", err)
		return err
	}
	n.lis = lis
	return nil
}

func (n *NoUndo) Serve() error {
	peer.RegisterHistoryReadServiceServer(n.grpc, NewGrpcServer(n.Self()))

	chGrpcErr := make(chan error)
	chHttpErr := make(chan error)

	go func(ch chan error) {
		ch <- n.grpc.Serve(n.lis)
	}(chGrpcErr)

	go func(ch chan error) {
		ch <- n.http.ListenAndServe()
	}(chHttpErr)

	slog.Debug("gRPC listening", "addr", n.lis.Addr().String())
	slog.Debug("http listening", "addr", n.http.Addr)

	// Gracefull stop is the waiting/sync point, as it waits to read from error channels
	// if it recieves a value (or a singal, it proceeds to Gracefuly stop all running services)
	n.GracefulStop(chGrpcErr, chHttpErr)
	return nil // todo
}

func (n *NoUndo) GracefulStop(chGrpcErr <-chan error, chHttpErr <-chan error) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	var err error
	select {
	case <-sigChan:
		break
	case err = <-chGrpcErr:
		slog.Error("chan grpc error:", err)
		break
	case err = <-chHttpErr:
		slog.Error("chan http error:", err)
		break
	}
	slog.Info("stopping servers", "noundo", n.Self().GetName())
	slog.Info("exiting select", "err", err)
	slog.Info("closing http", "err", n.http.Close())
	n.grpc.GracefulStop()
	slog.Info("grpc server gracefuly stopped.")
	slog.Info("closing net.Listener", "err", n.lis.Close())
}

// Alias for NoUndo.uni.Self()
func (n *NoUndo) Self() HistoryFullIface {
	return n.uni.Self()
}

// Alias for NoUndo.uni.Self()
func (n *NoUndo) Peers() []HistoryPublicIface {
	return n.uni.Peers()
}
