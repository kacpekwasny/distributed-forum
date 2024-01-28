package main

import (
	"crypto/rand"
	"math/big"

	n "github.com/kacpekwasny/noundo/pkg/noundo"
	"github.com/kacpekwasny/noundo/pkg/utils"
	lor "gopkg.in/loremipsum.v1"
)

var uni0 *n.Universe
var uni1 *n.Universe
var uni2 *n.Universe
var lorGen = lor.New()

func init() {
}

func createStories(h n.HistoryFullIface, m int, age string, author n.UserPublicIface, text string) {
	for i := 0; i < m; i++ {
		s, _ := h.CreateStory(author, age, &n.StoryContent{
			Title:   lorGen.Words(5),
			Content: lorGen.Sentences(rint(30)),
		})
		a1, _ := h.CreateAnswer(author, s.Id(), lorGen.Words(5+rint(30)))
		a12, _ := h.CreateAnswer(author, a1.PostableId, lorGen.Words(5+rint(30)))
		h.CreateAnswer(author, a12.PostableId, lorGen.Words(5+rint(30)))
		h.CreateAnswer(author, s.Id(), lorGen.Words(5+rint(30)))
	}
}

func rint(n int64) int {
	return int(utils.LeftLogRight(rand.Int(rand.Reader, big.NewInt(n))).Int64())
}
