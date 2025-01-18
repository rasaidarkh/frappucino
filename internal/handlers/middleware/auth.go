package middleware

import (
	"context"
	"fmt"
	"frappuccino/internal/helpers"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type key string

const ctxKey key = "payload"

func WithJWTAuth(rdb *redis.Client, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PathValue("token")
		if len(token) == 0 {
			helpers.WriteError(w, http.StatusForbidden, fmt.Errorf("token was not provided"))
			return
		}

		payload, err := rdb.HGetAll(context.Background(), token).Result()
		if err != nil {
			helpers.WriteError(w, http.StatusForbidden, fmt.Errorf("invalid token"))
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxKey, payload)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
