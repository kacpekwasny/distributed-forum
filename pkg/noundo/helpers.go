package noundo

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
	"unsafe"

	"github.com/kacpekwasny/noundo/pkg/utils"
)

type Id string

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	idLen         = 4
)

var src = rand.NewSource(time.Now().UnixNano())

// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
func NewRandId() Id {
	b := make([]byte, idLen)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := idLen-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return Id(*(*string)(unsafe.Pointer(&b)))
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

func ExecTemplHtmxSensitiveExplicitBase(tmpl *template.Template, w http.ResponseWriter, r *http.Request, pageName string, pageNameBase string, pushUrl string, data any) {
	if r.Header.Get("hx-request") == "true" {
		w.Header().Set("HX-Push-Url", pushUrl)
		utils.ExecTemplLogErr(tmpl, w, pageName, data)
		return
	}

	utils.ExecTemplLogErr(tmpl, w, pageNameBase, data)
}

func ExecTemplHtmxSensitive(tmpl *template.Template, w http.ResponseWriter, r *http.Request, pageName string, pushUrl string, data any) {
	ExecTemplHtmxSensitiveExplicitBase(tmpl, w, r, pageName, "page_"+pageName+".go.html", pushUrl, data)
}
