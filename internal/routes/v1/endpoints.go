package v1

import (
	"fmt"
	"net/http"
)

func (v *Router) Endpoints() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /ping", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))

	mux.Handle("GET /health", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "ok")
	}))

	return http.StripPrefix("/api/v1", mux)
}
