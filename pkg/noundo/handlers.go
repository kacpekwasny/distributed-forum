package noundo

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/noundo/pkg/utils"
)

// Return html from template when the request was made by HTMX, for the returned HTML,
// to be passed into the existing DOM
func HandleGetPageTemplateAsComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	w.Header().Set("HX-Push-Url", filename)
	err := tmpl.ExecuteTemplate(w, filename+".go.html", GetJWTFieldsFromContext(r.Context()))
	utils.Loge(err)
}

// This function as oposing to the HandleGetPageTemplateAsComponent returns the component, but
func HandleGetPageTemplateStandalone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	err := tmpl.ExecuteTemplate(w, filename+".go.html", nil)
	utils.Loge(err)
}

// When HTMX calls to retrieve a template from backend, it wants only the component
// but when the page is reloaded, the request has to return the component wrapped in the actucal
// proper HTML DOC body, with HEAD containing bootstrap, htmx, etc...
// This function is supposed to return the component wrapped in the base template.
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	err := tmpl.ExecuteTemplate(w, "base.go.html",
		BaseValues{
			Title:            "",
			MainComponentURL: filename},
	)
	utils.Loge(err)
}

func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	jwtf := GetJWTFieldsFromContext(r.Context())
	if jwtf == nil {
		utils.Loge(tmpl.ExecuteTemplate(w, "welcome.go.html", WelcomeValues{Msg: "It's a shame you didn't Sign In :(("}))
		return
	}
	utils.Loge(tmpl.ExecuteTemplate(w, "welcome.go.html", WelcomeValues{Username: jwtf.Username}))
}

func AddStory(w http.ResponseWriter, r *http.Request) {
	u1 := User{
		Id:       NewRandId(),
		Username: "kacper",
	}
	p1 := Story{
		Postable: Postable{
			id:            NewRandId(),
			userFUsername: u1.Username,
			contents:      "wubba lubba dab dab",
			TimeStampable: TimeStampable{timestamp: 0}},
		Reactionable: Reactionable{reactions: []Reaction{}},
	}
	RenderStory(w, &p1)
}
