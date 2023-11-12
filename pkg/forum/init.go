package forum

import (
	"html/template"
	"path/filepath"
	"runtime"

	"github.com/kacpekwasny/distributed-forum/pkg/utils"
	"github.com/leekchan/gtf"
)

var (
	tplPages      *template.Template
	tplComponents *template.Template
)

func init() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("was not able to aquire the filename of current running file")
	}
	dirname := filepath.Dir(filename)

	templatesPagesGlobSelector := filepath.Join(dirname, "templates_pages", "*.go.html")
	tplPages = template.Must(
		template.
			New("forumPages").
			Funcs(gtf.GtfFuncMap).
			Funcs(utils.FuncMapCommon).
			ParseGlob(templatesPagesGlobSelector))

	templatesComponentsGlobSelector := filepath.Join(dirname, "templates_components", "*.go.html")
	tplComponents = template.Must(
		template.
			New("forumComponents").
			Funcs(gtf.GtfFuncMap).
			Funcs(utils.FuncMapCommon).
			ParseGlob(templatesComponentsGlobSelector))

}
