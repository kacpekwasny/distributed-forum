package noundo

import "github.com/gorilla/mux"

func (n *NoUndo) setupRouterAndHandlers() {
	r := mux.NewRouter()

	r.HandleFunc("/", n.handleIndex).Methods("GET")

	// TODO - substitute these function for methods of NoUndo
	r.HandleFunc("/index", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentURL: "welcome"})).Methods("GET")
	r.HandleFunc("/welcome", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentURL: "welcome"})).Methods("GET")
	r.HandleFunc("/component_welcome", HandleWelcome).Methods("GET")

	r.HandleFunc("/signin", BaseGetFactory(BaseValues{"SignIn", "signin"})).Methods("GET")
	r.HandleFunc("/signin", HandlePostSignIn).Methods("POST")

	r.HandleFunc("/signout", HandleSignOut).Methods("GET", "POST")

	r.HandleFunc("/signup", BaseGetFactory(BaseValues{"Sign Up", "signup"})).Methods("GET")
	r.HandleFunc("/signup", HandlePostSignUp).Methods("POST")

	r.HandleFunc("/component_{filename}", HandleGetPageTemplateAsComponent).Methods("GET")

	r.HandleFunc("/{filename}", HandleDefault)

	n.r = r
}
