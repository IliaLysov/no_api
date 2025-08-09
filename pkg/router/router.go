package router

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(LogMiddleware)

	r.Get("/live", probe)

	return r
}

func probe(w http.ResponseWriter, _ *http.Request) {
	fmt.Println("Probe called")
	w.WriteHeader(http.StatusOK)
}
