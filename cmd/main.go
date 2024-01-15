package main

import (
	"sync"

	"github.com/kacpekwasny/noundo/pkg/noundo"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func runUniverse(wg *sync.WaitGroup, uni *noundo.Universe, addr string) {
	utils.Loge(noundo.NewNoUndo(uni).ListenAndServe(addr))
	wg.Done()
}

func main() {
	var wg = sync.WaitGroup{}
	wg.Add(3)
	go runUniverse(&wg, uni0, ":8080")
	go runUniverse(&wg, uni1, ":8081")
	go runUniverse(&wg, uni2, ":8082")
	wg.Wait()
}
