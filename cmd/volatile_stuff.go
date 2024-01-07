package main

import n "github.com/kacpekwasny/noundo/pkg/noundo"

var h0 n.HistoryFullIface
var volUniverse n.UniverseIface
var peersNexus n.PeersNexusIface

func init() {
	h0 = n.NewHistoryVolatile("localhost:8080")
	h1 := n.NewHistoryVolatile("localhost:8081")
	h2 := n.NewHistoryVolatile("localhost:8082")

	u0k, _ := h0.CreateUser("k", "k", "k")
	u1k, _ := h1.CreateUser("k", "k", "k")
	u2k, _ := h2.CreateUser("k", "k", "k")

	a0, _ := h0.CreateAge(u0k, "age0")
	a1, _ := h1.CreateAge(u1k, "age1")
	a2, _ := h2.CreateAge(u2k, "age2")

	h0.CreateStory(a0.GetName(), n.NewCreateStory(u0k.FullUsername(), "# My first post\n prev was the header. this is the content."))
	h1.CreateStory(a1.GetName(), n.NewCreateStory(u1k.FullUsername(), "# My first post\n prev was the header. this is the content."))
	h2.CreateStory(a2.GetName(), n.NewCreateStory(u2k.FullUsername(), "# My first post\n prev was the header. this is the content."))

	peersNexus = n.NewPeersNexus()
	peersNexus.RegisterPeerManager(n.NewPeerManagerDummy(h1))
	peersNexus.RegisterPeerManager(n.NewPeerManagerDummy(h2))

	volUniverse = n.NewUniverse(h0, peersNexus)
}
