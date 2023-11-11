package forum

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/distributed-forum/pkg/auth"
	"github.com/kacpekwasny/distributed-forum/pkg/utils"
)

var authenticator auth.Authenticator

func init() {
	authenticator = auth.NewVolatileAuthenticator()
}

func HandleGetComponent(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	w.Header().Set("HX-Push-Url", filename)
	err := tpl.ExecuteTemplate(w, filename+".go.html", nil)
	utils.Pife(err)
}

func HandleWelcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("waaaaa")
	user, ok := auth.GetUsernameFromContext(r.Context())
	if !ok {
		utils.Pife(tpl.ExecuteTemplate(w, "welcome.go.html", WelcomeValues{Msg: "It's a shame you didn't Log In :(("}))
		return
	}
	fmt.Println("welcome", user)
	utils.Pife(tpl.ExecuteTemplate(w, "welcome.go.html", WelcomeValues{Username: user}))
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	u1 := User{
		Id:       NewRandId(),
		Username: "kacper",
	}
	p1 := Post{
		Postable: Postable{
			Id:            NewRandId(),
			UserId:        u1.Id,
			Contents:      "wubba lubba dab dab",
			TimeStampable: TimeStampable{Timestamp: 0}},
		Reactionable: Reactionable{Reactions: []Reaction{}},
	}
	RenderPost(w, &p1)
}
