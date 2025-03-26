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

	mux.Handle("GET /songs", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.GetSongs(w, r)
	}))

	mux.Handle("GET /song", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.GetSong(w, r)
	}))

	mux.Handle("GET /text", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.GetText(w, r)
	}))

	mux.Handle("DELETE /song", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.DeleteSong(w, r)
	}))

	mux.Handle("POST /song", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.CreateSong(w, r)
	}))

	mux.Handle("PUT /song", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v.View.EditSong(w, r)
	}))

	return http.StripPrefix("/api/v1", mux)
}
