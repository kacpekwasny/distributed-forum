package main

import n "github.com/kacpekwasny/noundo/pkg/noundo"

var volHistory n.HistoryFullIface
var volUniverse n.UniverseIface
var peersNexus n.PeersNexusIface

func init() {
	volHistory = n.NewHistoryVolatile("localhost:8080")
	uk, _ := volHistory.AddUser("k", "k", "k")
	volHistory.CreateAge(uk, "age1")
	peersNexus = n.NewPeersNexus()
	peersNexus.RegisterPeerManager(n.NewPeerManagerDummy(n.NewHistoryVolatile("localhost:8081")))
	peersNexus.RegisterPeerManager(n.NewPeerManagerDummy(n.NewHistoryVolatile("localhost:8082")))
	volUniverse = n.NewUniverse(volHistory, peersNexus)
}
