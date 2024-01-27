package noundo

import "net/http"

func (n *NoUndo) Handle404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	ExecTemplHtmxSensitive(tmpl, w, r, "404", r.URL.Path, CreatePageBaseValues("404", n.Self(), n.Self(), r))
}

func (n *NoUndo) Handle401(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusUnauthorized)
	ExecTemplHtmxSensitive(tmpl, w, r, "401", r.URL.Path, Page401Values{
		RequestedPath:  r.URL.Path,
		PageBaseValues: CreatePageBaseValues("401", n.Self(), n.Self(), r),
	})
}
