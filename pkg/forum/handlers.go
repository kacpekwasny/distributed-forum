package forum

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/distributed-forum/pkg/auth"
	"github.com/kacpekwasny/distributed-forum/pkg/utils"
)

var authenticator auth.Authenticator

func init() {
	authenticator = auth.NewVolatileAuthenticator()
}

// Return html from template when the request was made by HTMX, for the returned HTML,
// to be passed into the existing DOM
func HandleGetPageTemplateAsComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	w.Header().Set("HX-Push-Url", filename)
	err := tplPages.ExecuteTemplate(w, filename+".go.html", auth.GetJWTFieldsFromContext(r.Context()))
	utils.Loge(err)
}

// This function as oposing to the HandleGetPageTemplateAsComponent returns the component, but
func HandleGetPageTemplateStandalone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]

	err := tplPages.ExecuteTemplate(w, filename+".go.html", nil)
	utils.Loge(err)
}

// When HTMX calls to retrieve a template from backend, it wants only the component
// but when the page is reloaded, the request has to return the component wrapped in the actucal
// proper HTML DOC body, with HEAD containing bootstrap, htmx, etc...
// This function is supposed to return the component wrapped in the base template.
func HandleDefault(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	err := tplPages.ExecuteTemplate(w, "base.go.html",
		BaseValues{
			Title:          "",
			MainContentUrl: filename},
	)
	utils.Loge(err)
}

func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	jwtf := auth.GetJWTFieldsFromContext(r.Context())
	if jwtf == nil {
		utils.Loge(tplPages.ExecuteTemplate(w, "welcome.go.html", WelcomeValues{Msg: "It's a shame you didn't Log In :(("}))
		return
	}
	utils.Loge(tplPages.ExecuteTemplate(w, "welcome.go.html", WelcomeValues{Username: jwtf.Username}))
}

func AddStory(w http.ResponseWriter, r *http.Request) {
	u1 := User{
		Id:       NewRandId(),
		Username: "kacper",
	}
	p1 := Story{
		Postable: Postable{
			Id:            NewRandId(),
			UserId:        u1.Id,
			Contents:      "wubba lubba dab dab",
			TimeStampable: TimeStampable{Timestamp: 0}},
		Reactionable: Reactionable{Reactions: []Reaction{}},
	}
	RenderStory(w, &p1)
}
