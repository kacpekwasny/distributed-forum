package auth

import (
	"context"
	"log"
	"net/http"
)

type privCtxType string

const ctxUsernameKey privCtxType = "username"

// If request not authenticated - return error
func GetUsernameFromContext(ctx context.Context) (string, bool) {
	v := ctx.Value(ctxUsernameKey)
	username, ok := v.(string)
	return username, ok
}

func HttpAuthenticator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, err := JWTCheckAndParse(r)
		if err == nil {
			log.Println("request from an authenticated user:", jwt.Username)
			ctx := r.Context()
			ctx = context.WithValue(ctx, ctxUsernameKey, jwt.Username)
			*r = *r.WithContext(ctx)
		}
		h.ServeHTTP(w, r)
	})
}
