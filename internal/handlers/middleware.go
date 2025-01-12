package handlers

import (
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// TODO: User Authentication
			ctx := r.Context()
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		})
}
