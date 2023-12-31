package noundo

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {

	r := mux.NewRouter()
	r.Use(HttpAuthenticator)

	r.HandleFunc("/", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentURL: "welcome"})).Methods("GET")
	r.HandleFunc("/welcome", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentURL: "welcome"})).Methods("GET")
	r.HandleFunc("/component_welcome", HandleWelcome).Methods("GET")

	r.HandleFunc("/signin", BaseGetFactory(BaseValues{"Sign In", "signin"})).Methods("GET")
	r.HandleFunc("/signin", HandlePostSignIn).Methods("POST")

	r.HandleFunc("/signout", HandleSignOut).Methods("GET", "POST")

	r.HandleFunc("/signup", BaseGetFactory(BaseValues{"Sign Up", "signup"})).Methods("GET")
	r.HandleFunc("/signup", HandlePostSignUp).Methods("POST")

	r.HandleFunc("/component_{filename}", HandleGetPageTemplateAsComponent).Methods("GET")

	r.HandleFunc("/{filename}", HandleDefault)

	http.Handle("/", r)
	return r
}

func ListenAndServe() {
	r := setupRouter()

	const addr = "127.0.0.1:8083"

	log.Println("listening on: ", addr)
	err := http.ListenAndServe(addr, r)

	if err != nil {
		log.Fatal(err)
	}
}
