package noundo

import (
	"net/http"
)

type NoUndo struct {
	uni UniverseIface
}

func (n *NoUndo) InjectHistory(a AgeIface) {

}

func (n *NoUndo) Inject() {

}

func (n *NoUndo) handleLogin(w http.ResponseWriter, r *http.Request) {

}
