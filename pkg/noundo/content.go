package noundo

import (
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/kacpekwasny/distributed-forum/pkg/utils"
)

func NewRandId() Id {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return Id(math.MaxUint64 * r.Float64())
}

func RenderStory(w http.ResponseWriter, p *Story) {
	err := tplPages.Execute(w, []*Story{p, p})
	utils.Pife(err)
}
