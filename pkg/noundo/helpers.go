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
		utils.ExecTemplLogErr(tmpl, w, "base.go.html", baseValues)
	}
}

func ComponentGetFactory(template string, v any) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.ExecuteTemplate(w, template, v)
		utils.Loge(err)
	}
}

func PageHandlerFactory(pageName string, pushUrl string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ExecTemplHtmxSensitive(tmpl, w, r, pageName, pushUrl, nil)
	}
}

func ExecTemplHtmxSensitiveExplicitBase(tpl *template.Template, w http.ResponseWriter, r *http.Request, pageName string, pageNameBase string, pushUrl string, data any) {
	if r.Header.Get("hx-request") == "true" {
		w.Header().Set("HX-Push-Url", pushUrl)
		utils.ExecTemplLogErr(tpl, w, pageName, data)
		return
	}

	utils.ExecTemplLogErr(tpl, w, pageNameBase, data)
}

func ExecTemplHtmxSensitive(tpl *template.Template, w http.ResponseWriter, r *http.Request, pageName string, pushUrl string, data any) {
	ExecTemplHtmxSensitiveExplicitBase(tpl, w, r, pageName, "page_"+pageName+".go.html", pushUrl, data)
}
