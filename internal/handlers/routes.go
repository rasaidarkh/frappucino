package handlers

import (
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.HandleFunc("/{x}", func(w http.ResponseWriter, r *http.Request) {})

	return Middleware(mux)
}
