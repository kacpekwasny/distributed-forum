package forum

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kacpekwasny/distributed-forum/pkg/auth"
)

func setupRouter() *mux.Router {

	r := mux.NewRouter()
	r.Use(auth.HttpAuthenticator)

	r.HandleFunc("/", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentUrl: "welcome"})).Methods("GET")
	r.HandleFunc("/welcome", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentUrl: "welcome"})).Methods("GET")
	r.HandleFunc("/component_welcome", HandleWelcome).Methods("GET")

	r.HandleFunc("/login", BaseGetFactory(BaseValues{"Login", "login"})).Methods("GET")
	r.HandleFunc("/login", HandlePostLogin).Methods("POST")

	r.HandleFunc("/register", BaseGetFactory(BaseValues{"Register", "register"})).Methods("GET")
	r.HandleFunc("/register", HandlePostRegister).Methods("POST")

	r.HandleFunc("/component_{filename}", HandleGetComponent).Methods("GET")

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
