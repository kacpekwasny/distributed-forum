package noundo

import (
	"log/slog"

	"github.com/go-chi/httplog/v2"
	"github.com/gorilla/mux"
)

func (n *NoUndo) setupRouter() {
	r := mux.NewRouter()

	r.Use(httpLogger())

	r.HandleFunc("/", n.HandleIndex).Methods("GET")

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

func httpLogger() mux.MiddlewareFunc {
	// Logger
	logger := httplog.NewLogger("httplog-example", httplog.Options{
		LogLevel:         slog.LevelDebug,
		Concise:          true,
		RequestHeaders:   true,
		MessageFieldName: "message",
		Tags: map[string]string{
			"version": "v1.0-81aa4244d9fc8076a",
			"env":     "dev",
		},
	})

	return httplog.RequestLogger(logger)
}
