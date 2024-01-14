package noundo

import "net/http"

func Handle404(w http.ResponseWriter, r *http.Request) {
	ExecTemplHtmxSensitive(tmpl, w, r, "404", r.URL.Path, nil)
}
