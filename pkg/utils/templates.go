package utils

import (
	"html/template"
	"net/http"
)

func ExecTemplLogErr(t *template.Template, w http.ResponseWriter, name string, data any) {
	Loge(t.ExecuteTemplate(w, "welcome.go.html", data))
}
