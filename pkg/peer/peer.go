package peer

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GrpcServer struct {
	listener   *net.Listener
	grpcServer *grpc.Server
}

type PeerConnection struct {
	UnimplementedHelloServiceServer
}

// SayHello implements helloworld.GreeterServer
func (s *PeerConnection) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	log.Printf("Received: %v", in.GetGreeting())
	return &HelloResponse{Reply: "Hello " + in.GetGreeting()}, nil
}

// Serve grpc and return error channel, close Channel.
// func (pc *PeerConnection) Serve() <-chan error {
// 	errCh := make(chan error)
// 	go func() {
// 		errCh <- pc.grpcServer.Serve(*pc.listener)
// 		close(errCh)
// 	}()
// 	return errCh
// }
