package noundo

import (
	"context"
	"log/slog"
	"net/http"
)

type privCtxType string

const ctxJWTkey privCtxType = "username"

// If request not authenticated - return error
func GetJWT(ctx context.Context) *JWTFields {
	v := ctx.Value(ctxJWTkey)
	if v == nil {
		return nil
	}
	jwtfields, ok := v.(JWTFields)
	if !ok {
		return nil
	}
	return &jwtfields
}

func HttpAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwt, err := JWTCheckAndParse(r)
		if err == nil {
			slog.Debug("request from an authenticated user", "username", jwt.Username)
			*r = *AddJWTtoCtx(r, jwt)
		}
		next.ServeHTTP(w, r)
	})
}

func AddJWTtoCtx(r *http.Request, jwt JWTFields) *http.Request {
	return r.WithContext(context.WithValue(r.Context(), ctxJWTkey, jwt))
}
