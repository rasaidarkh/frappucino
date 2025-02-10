package middleware

import (
	"context"
	"fmt"
	"frappuccino/internal/helpers"
	"frappuccino/pkg/config"
	"frappuccino/pkg/jtoken"
	"net/http"
)

type key string

const ctxKey key = "payload"

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PathValue("token")
		if len(token) == 0 {
			helpers.WriteError(w, http.StatusForbidden, fmt.Errorf("token was not provided"))
			return
		}

		payload, err := jtoken.VerifyJWT(token, config.GetJWTSecret())
		if err != nil {
			helpers.WriteError(w, http.StatusForbidden, err)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxKey, payload)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
