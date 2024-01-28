package main

import (
	"sync"
	"time"

	n "github.com/kacpekwasny/noundo/pkg/noundo"

	"github.com/kacpekwasny/noundo/pkg/noundo"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func runUniverse(wg *sync.WaitGroup, uni *noundo.Universe, httpAddr string, grpcAddr string) {
	n := noundo.NewNoUndo(uni)
	n.SetupListen(httpAddr, grpcAddr)
	utils.Loge(n.Serve())
	wg.Done()
}

func main() {
	h0 := n.NewHistoryVolatile("127.0.0.1:8080")
	h1 := n.NewHistoryVolatile("127.0.0.1:8081")
	h2 := n.NewHistoryVolatile("127.0.0.1:8082")

	u0k, _ := h0.CreateUser("k", "k0", "k")
	u1k, _ := h1.CreateUser("k", "k1", "k")
	u2k, _ := h2.CreateUser("k", "k2", "k")

	a0, _ := h0.CreateAge(u0k, "age0")
	a1, _ := h1.CreateAge(u1k, "age1")
	a2, _ := h2.CreateAge(u2k, "age2")

	createStories(h0, 5, a0.GetName(), u0k, "# My first post\n prev was the header. this is the content.")
	createStories(h1, 5, a1.GetName(), u1k, "# My first post\n prev was the header. this is the content.")
	createStories(h2, 5, a2.GetName(), u2k, "# My first post\n prev was the header. this is the content.")

	var wg = sync.WaitGroup{}
	wg.Add(3)
	peersNexus0 := n.NewPeersNexus()
	peersNexus1 := n.NewPeersNexus()
	peersNexus2 := n.NewPeersNexus()

	uni0 = n.NewUniverse(h0, peersNexus0)
	uni1 = n.NewUniverse(h1, peersNexus1)
	uni2 = n.NewUniverse(h2, peersNexus2)

	go runUniverse(&wg, uni0, ":8080", ":8090")
	go runUniverse(&wg, uni1, ":8081", ":8091")
	go runUniverse(&wg, uni2, ":8082", ":8092")

	time.Sleep(time.Second * 1)

	peersNexus0.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8091"))
	peersNexus0.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8092"))
	// peersNexus0.RegisterPeerManager(n.NewPeerManagerDummy(h1))
	// peersNexus0.RegisterPeerManager(n.NewPeerManagerDummy(h2))

	peersNexus1.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8090"))
	peersNexus1.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8092"))
	// peersNexus1.RegisterPeerManager(n.NewPeerManagerDummy(h0))
	// peersNexus1.RegisterPeerManager(n.NewPeerManagerDummy(h2))

	peersNexus2.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8090"))
	peersNexus2.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8091"))
	// peersNexus2.RegisterPeerManager(n.NewPeerManagerDummy(h1))
	// peersNexus2.RegisterPeerManager(n.NewPeerManagerDummy(h0))

	wg.Wait()
}
