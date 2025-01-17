package middleware

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type key string

const ctxKey key = "payload"

func WithJWTAuth(rdb *redis.Client, handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.PathValue("token")
		if len(token) == 0 {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		payload, err := rdb.HGetAll(context.Background(), token).Result()
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := r.Context()
		ctx = context.WithValue(ctx, ctxKey, payload)
		r = r.WithContext(ctx)

		handlerFunc(w, r)
	}
}
