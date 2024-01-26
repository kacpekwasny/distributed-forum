package main

import (
	n "github.com/kacpekwasny/noundo/pkg/noundo"
	lor "gopkg.in/loremipsum.v1"
)

var uni0 *n.Universe
var uni1 *n.Universe
var uni2 *n.Universe
var lorGen = lor.New()

func init() {
	h0 := n.NewHistoryVolatile("localhost:8080")
	h1 := n.NewHistoryVolatile("localhost:8081")
	h2 := n.NewHistoryVolatile("localhost:8082")

	u0k, _ := h0.CreateUser("k", "k", "k")
	u1k, _ := h1.CreateUser("k", "k", "k")
	u2k, _ := h2.CreateUser("k", "k", "k")

	a0, _ := h0.CreateAge(u0k, "age0")
	a1, _ := h1.CreateAge(u1k, "age1")
	a2, _ := h2.CreateAge(u2k, "age2")

	createStories(h0, 5, a0.GetName(), u0k, "# My first post\n prev was the header. this is the content.")
	createStories(h1, 5, a1.GetName(), u1k, "# My first post\n prev was the header. this is the content.")
	createStories(h2, 5, a2.GetName(), u2k, "# My first post\n prev was the header. this is the content.")

	peersNexus0 := n.NewPeersNexus()
	peersNexus0.RegisterPeerManager(n.NewPeerManagerDummy(h1))
	peersNexus0.RegisterPeerManager(n.NewPeerManagerDummy(h2))

	peersNexus1 := n.NewPeersNexus()
	peersNexus1.RegisterPeerManager(n.NewPeerManagerDummy(h0))
	peersNexus1.RegisterPeerManager(n.NewPeerManagerDummy(h2))

	peersNexus2 := n.NewPeersNexus()
	peersNexus2.RegisterPeerManager(n.NewPeerManagerDummy(h1))
	peersNexus2.RegisterPeerManager(n.NewPeerManagerDummy(h0))

	uni0 = n.NewUniverse(h0, peersNexus0)
	uni1 = n.NewUniverse(h1, peersNexus1)
	uni2 = n.NewUniverse(h2, peersNexus2)
}

func createStories(h n.HistoryFullIface, m int, age string, author n.UserPublicIface, text string) {
	for i := 0; i < m; i++ {
		h.CreateStory(age, author, n.StoryContent{
			Title:   lorGen.Words(5),
			Content: lorGen.Sentences(3),
		})
	}
}
