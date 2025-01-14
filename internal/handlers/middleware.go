package handlers

import (
	"context"
	"net/http"
	"time"
)

func Middleware(handler http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			// TODO: User Authentication
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			r = r.WithContext(ctx)
			handler.ServeHTTP(w, r)
		})
}

func IsAuthenticated() bool {
	// TODO: Redis session with user authentication
	return false
}
