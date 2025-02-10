package middleware

import (
	"net/http"
)

type key string

const ctxKey key = "payload"

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
