package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gume1a/oauth-proxy/internal/registry"
	"log"
	"net/http"
)

func main() {
	registrar := registry.GetRegistry()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/oauth", func(r chi.Router) {
		r.Post("/", func(rw http.ResponseWriter, req *http.Request) {
			registrar.ProxyServeHTTP(rw, req)
		})
	})

	r.Get("/supported", func(rw http.ResponseWriter, req *http.Request) {
		providers := registrar.Providers()

		enc := json.NewEncoder(rw)
		if err := enc.Encode(providers); err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
		}
	})

	fs := http.FileServer(http.Dir("./assets/static"))
	r.Handle("/*", fs)

	log.Fatal(http.ListenAndServe("localhost:8081", r))
}
