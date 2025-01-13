package handlers

import (
	"net/http"
)

func Middleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// TODO: User Authentication
			w.Header().Set("middle-level-projcet", "true")

			ctx := r.Context()
			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		})
}

func IsAuthenticated() bool {
	// TODO: Redis session with user authentication
	return false
}
