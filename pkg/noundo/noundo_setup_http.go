package noundo

import (
	"log/slog"

	"github.com/go-chi/httplog/v2"
	"github.com/gorilla/mux"
)

func (n *NoUndo) setupRouter() {
	r := mux.NewRouter()

	r.Use(HttpAuthenticator(n.Self().Authenticator()))
	r.Use(httpLogger())

	// p - endpoints using Subrouter `p` require request to be authenticated
	// otherwise, they will return 404
	// if more personalized unauth handling is required
	// user router r, and then handle it yourself with 'GetJWT'
	// p := r.NewRoute().Subrouter()
	// p.Use(n.AuthOr401)

	r.HandleFunc("/", n.HandleHome).Methods("GET")
	r.HandleFunc("/a/{age}", n.HandleAgeShortcut).Methods("GET")
	r.HandleFunc("/a/{history}/{age}", n.HandleAge).Methods("GET")

	r.HandleFunc("/a/{history}/{age}/create-story", n.HandleCreateStoryPost).Methods("POST")
	r.HandleFunc("/a/{history}/story/{story-id}", n.HandleStoryGet).Methods("GET")

	r.HandleFunc("/write-answer-box/{postable-id}", n.HandleCreateAnswerGet).Methods("GET")
	r.HandleFunc("/write-answer/{history}/{postable-id}", n.HandleCreateAnswerPost).Methods("POST")

	r.HandleFunc("/profile", n.HandleSelfProfile).Methods("GET")
	r.HandleFunc("/profile/{username}", n.HandleProfile).Methods("GET")

	r.HandleFunc("/signin", n.HandleSignInGet).Methods("GET")
	r.HandleFunc("/signin", n.HandleSignInPost).Methods("POST")

	r.HandleFunc("/signout", n.HandleSignOut).Methods("GET", "POST")

	r.HandleFunc("/signup", n.HandleSignUpGet).Methods("GET")
	r.HandleFunc("/signup", n.HandleSignUpPost).Methods("POST")

	r.NotFoundHandler = r.NewRoute().HandlerFunc(n.Handle404).GetHandler()

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
