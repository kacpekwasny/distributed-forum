package main

import (
	"github.com/kacpekwasny/distributed-forum/pkg/noundo"
)

func main() {
	noundo.ListenAndServe()
}
