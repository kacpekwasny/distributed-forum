package noundo

import (
	"context"
	"log"
	"net/http"
)

type privCtxType string

const ctxUsernameKey privCtxType = "username"

// If request not authenticated - return error
func GetJWTFieldsFromContext(ctx context.Context) *JWTFields {
	v := ctx.Value(ctxUsernameKey)
	if v == nil {
		return nil
	}
	jwtfields, ok := v.(JWTFields)
	if !ok {
		return nil
	}
	return &jwtfields
}

func HttpAuthenticator(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, err := JWTCheckAndParse(r)
		if err == nil {
			log.Println("request from an authenticated user:", jwt.Username)
			ctx := r.Context()
			ctx = context.WithValue(ctx, ctxUsernameKey, jwt)
			*r = *r.WithContext(ctx)
		}
		h.ServeHTTP(w, r)
	})
}
