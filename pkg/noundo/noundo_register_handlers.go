package noundo

import "github.com/gorilla/mux"

func (n *NoUndo) setupRouterAndHandlers() {
	r := mux.NewRouter()

	r.HandleFunc("/", n.handleIndex).Methods("GET")

	// TODO - substitute these function for methods of NoUndo
	r.HandleFunc("/index", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentURL: "welcome"})).Methods("GET")
	r.HandleFunc("/welcome", BaseGetFactory(BaseValues{Title: "Welcome, to the internet!", MainContentURL: "welcome"})).Methods("GET")
	r.HandleFunc("/component_welcome", HandleWelcome).Methods("GET")

	r.HandleFunc("/login", BaseGetFactory(BaseValues{"Login", "login"})).Methods("GET")
	r.HandleFunc("/login", HandlePostLogin).Methods("POST")

	r.HandleFunc("/logout", HandleLogout).Methods("GET", "POST")

	r.HandleFunc("/register", BaseGetFactory(BaseValues{"Register", "register"})).Methods("GET")
	r.HandleFunc("/register", HandlePostRegister).Methods("POST")

	r.HandleFunc("/component_{filename}", HandleGetPageTemplateAsComponent).Methods("GET")

	r.HandleFunc("/{filename}", HandleDefault)

	n.r = r
}
