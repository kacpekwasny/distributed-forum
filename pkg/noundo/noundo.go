package noundo

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type NoUndo struct {
	uni UniverseIface
	r   *mux.Router
}

func NewNoUndo(uni UniverseIface) *NoUndo {
	n := &NoUndo{
		uni: uni,
	}
	n.setupRouter()
	return n
}

func (n *NoUndo) ListenAndServe(addr string) error {
	log.Println("Listening on", addr)
	return http.ListenAndServe(addr, n.r)
}

// Alias for NoUndo.uni.Self()
func (n *NoUndo) Self() HistoryFullIface {
	return n.uni.Self()
}

// Alias for NoUndo.uni.Self()
func (n *NoUndo) Peers() []HistoryPublicIface {
	return n.uni.Peers()
}
