package noundo

import "net/http"

func (n *NoUndo) AuthOr401(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if GetJWT(r.Context()) == nil {
			n.Handle401(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
