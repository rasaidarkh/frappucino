package handlers

import (
	"fmt"
	"net/http"
)

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.NotFoundHandler())
	mux.HandleFunc("/{x}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Context())
	})

	return Middleware(mux)
}
