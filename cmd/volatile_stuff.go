package main

import "github.com/kacpekwasny/noundo/pkg/noundo"

var volHistory noundo.HistoryIface
var volUniverse noundo.UniverseIface
var peersFunnel *noundo.PeersFunnel

func init() {
	volHistory = &noundo.HistoryVolatile{}
	peersFunnel = noundo.NewPeersFunnel()
	volUniverse = noundo.NewUniverse(volHistory, peersFunnel)
}
