package main

import "github.com/kacpekwasny/noundo/pkg/noundo"

var volHistory noundo.HistoryIface
var volUniverse noundo.UniverseIface
var peersNexus *noundo.PeersNexus

func init() {
	volHistory = &noundo.HistoryVolatile{}
	peersNexus = noundo.NewPeersNexus()
	volUniverse = noundo.NewUniverse(volHistory, peersNexus)
}
