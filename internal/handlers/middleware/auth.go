package middleware

import "net/http"

func WithJWTAuth(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Perform JWT authentication logic here
		// If authenticated, call the original handler
		handlerFunc(w, r)
	}
}
