package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	"github.com/kacpekwasny/distributed-forum/pkg/peer"
	df "github.com/kacpekwasny/distributed-forum/pkg/test"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func grpcServerAndListener(port string) (*grpc.Server, net.Listener, error) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, nil, err
	}
	return grpc.NewServer(), listener, nil
}

func serve(s *grpc.Server, lis net.Listener, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 1900)
	go func() {
		time.Sleep(time.Second * 3)
		s.GracefulStop()
	}()
	s.Serve(lis)
	wg.Done()
}

func client(addr string) peer.HelloServiceClient {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	log.Println("before dial")
	conn, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.Println("after dial")
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := peer.NewHelloServiceClient(conn)
	return c
}

func main() {
	fmt.Println(df.Constant)
	grpcServer1, listener1, err := grpcServerAndListener(":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer2, listener2, err := grpcServerAndListener(":50052")
	if err != nil {
		log.Fatal(err)
	}
	peer.RegisterHelloServiceServer(grpcServer1, &peer.PeerConnection{})
	peer.RegisterHelloServiceServer(grpcServer2, &peer.PeerConnection{})

	wg := sync.WaitGroup{}
	wg.Add(2)
	go serve(grpcServer1, listener1, &wg)
	go serve(grpcServer2, listener2, &wg)

	// time.Sleep(time.Second * 1)
	c1 := client("localhost:50051")
	c2 := client("localhost:50052")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	c1.SayHello(ctx, &peer.HelloRequest{Greeting: "Hello from c1 :)"})
	c2.SayHello(ctx, &peer.HelloRequest{Greeting: "Hello from c2 :)"})
	log.Println("wg wait")
	wg.Wait()
}
