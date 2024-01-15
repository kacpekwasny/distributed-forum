package noundo

import (
	"html/template"
	"path/filepath"
	"runtime"

	"github.com/kacpekwasny/noundo/pkg/utils"
	"github.com/leekchan/gtf"
)

var (

	// Pages with and without head built in
	// "home" 			-> page only, prepared for an htmx response.
	// "home.go.html" 	-> page with htmx, bootstrap, etc links in head built in.
	tmpl *template.Template
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("was not able to aquire the filename of current running file")
	}
	dirname := filepath.Dir(filename)

	templatesPagesGlobSelector := filepath.Join(dirname, "templates_pages", "*.go.html")
	tmpl = template.Must(
		template.
			New("forumPages").
			Funcs(gtf.GtfFuncMap).
			Funcs(utils.FuncMapCommon).
			ParseGlob(templatesPagesGlobSelector))
}
