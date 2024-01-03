package noundo

import (
	"html/template"
	"math"
	"math/rand"
	"net/http"
	"time"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

func NewRandId() Id {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return Id(math.MaxUint64 * r.Float64())
}

func RenderStory(w http.ResponseWriter, p *Story) {
	err := tmpl.Execute(w, []*Story{p, p})
	utils.Pife(err)
}

func BaseGetFactory(baseValues BaseValues) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, _ *http.Request) {
		utils.ExecTemplLogErr(tmpl, w, "base", baseValues)
	}
}

func ComponentGetFactory(template string, v any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, template, v)
		utils.Loge(err)
	}
}

func ExecTemplHtmxSensitiveExplicitBase(tpl *template.Template, w http.ResponseWriter, r *http.Request, data any, pageName string, pageNameBase string) {
	if r.Header.Get("hx-request") == "true" {
		utils.ExecTemplLogErr(tpl, w, pageName, data)
		return
	}

	utils.ExecTemplLogErr(tpl, w, pageNameBase, data)
}

func ExecTemplHtmxSensitive(tpl *template.Template, w http.ResponseWriter, r *http.Request, data any, pageName string) {
	ExecTemplHtmxSensitiveExplicitBase(tpl, w, r, data, pageName, "page_"+pageName+".go.html")
}
