package main

import (
	"github.com/kacpekwasny/distributed-forum/pkg/forum"
)

func main() {
	// cm := forum.NewContentManagerVolatile()
	// cm.AddPost(&u1, &p1)
	forum.ListenAndServe()
}
