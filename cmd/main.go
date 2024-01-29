package main

import (
	"crypto/rand"
	"math/big"
	"sync"
	"time"

	n "github.com/kacpekwasny/noundo/pkg/noundo"
	lor "gopkg.in/loremipsum.v1"

	"github.com/kacpekwasny/noundo/pkg/noundo"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func runUniverse(wg *sync.WaitGroup, uni *noundo.Universe, httpAddr string, grpcAddr string) {
	n := noundo.NewNoUndo(uni)
	n.SetupListen(httpAddr, grpcAddr)
	utils.Loge(n.Serve())
	wg.Done()
}

var uni0 *n.Universe
var uni1 *n.Universe
var uni2 *n.Universe
var lorGen = lor.New()

func createStories(h n.HistoryFullIface, m int, age string, author n.UserPublicIface) {
	for i := 0; i < m; i++ {
		s, _ := h.CreateStory(author, age, &n.StoryContent{
			Title:   lorGen.Words(5),
			Content: lorGen.Sentences(rint(30)),
		})
		a1, _ := h.CreateAnswer(author, s.Id(), lorGen.Words(5+rint(30)))
		a12, _ := h.CreateAnswer(author, a1.PostableId, lorGen.Words(5+rint(30)))
		h.CreateAnswer(author, a12.PostableId, lorGen.Words(5+rint(30)))
		h.CreateAnswer(author, s.Id(), lorGen.Words(5+rint(30)))
	}
}

func rint(n int64) int {
	return int(utils.LeftLogRight(rand.Int(rand.Reader, big.NewInt(n))).Int64())
}

func main() {
	// An interface for in RAM history
	h0 := n.NewHistoryVolatile("127.0.0.1:8080")
	h1 := n.NewHistoryVolatile("127.0.0.1:8081")
	h2 := n.NewHistoryVolatile("127.0.0.1:8082")

	// Create users one in every History
	u0k, _ := h0.CreateUser("k", "k0", "k")
	u1k, _ := h1.CreateUser("k", "k1", "k")
	u2k, _ := h2.CreateUser("k", "k2", "k")

	// Create Ages in every history
	a00, _ := h0.CreateAge(u0k, "a00")
	a01, _ := h0.CreateAge(u0k, "a01")

	a10, _ := h1.CreateAge(u1k, "a10")
	a11, _ := h1.CreateAge(u1k, "a11")

	a20, _ := h2.CreateAge(u2k, "a20")
	a21, _ := h2.CreateAge(u2k, "a21")

	createStories(h0, 5, a00.GetName(), u0k)
	createStories(h0, 5, a01.GetName(), u0k)
	createStories(h1, 5, a10.GetName(), u1k)
	createStories(h1, 5, a11.GetName(), u1k)
	createStories(h2, 5, a20.GetName(), u2k)
	createStories(h2, 5, a21.GetName(), u2k)

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

	// Wait for gRPC servers to start
	time.Sleep(time.Second * 1)

	peersNexus0.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8091"))
	peersNexus0.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8092"))

	peersNexus1.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8090"))
	peersNexus1.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8092"))

	peersNexus2.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8090"))
	peersNexus2.RegisterPeerManager(n.NewPeerManagerGrpc("127.0.0.1:8091"))

	wg.Wait()
}
