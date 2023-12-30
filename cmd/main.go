package main

import (
	"github.com/kacpekwasny/noundo/pkg/noundo"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func main() {
	his := &noundo.HistoryVolatile{}
	uni := noundo.NewUniverse(his)
	utils.Loge(noundo.NewNoUndo(uni).ListenAndServe(":8080"))
}
