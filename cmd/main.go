package main

import (
	"sync"

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
	var wg = sync.WaitGroup{}
	wg.Add(3)
	go runUniverse(&wg, uni0, ":8080", ":8090")
	go runUniverse(&wg, uni1, ":8081", ":8091")
	go runUniverse(&wg, uni2, ":8082", ":8092")
	wg.Wait()
}
