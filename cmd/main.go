package main

import (
	"github.com/kacpekwasny/noundo/pkg/noundo"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

func main() {
	utils.Loge(noundo.NewNoUndo(volUniverse).ListenAndServe(":8080"))
}
