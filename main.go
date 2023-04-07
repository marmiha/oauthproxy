package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gume1a/oauthproxy/internal/config"
	"github.com/gume1a/oauthproxy/internal/registry"
	"github.com/gume1a/oauthproxy/pkg/identity"
	"github.com/joho/godotenv"
)

//go:embed assets/terminal/logo_banner.txt
var terminalAsciiArt string

//go:embed all:web/build/*
var embedWebsiteFS embed.FS
var websiteFS, _ = fs.Sub(embedWebsiteFS, "web/build")

func main() {
	fmt.Print(terminalAsciiArt)
	fmt.Print("\n\n")

	// Router configuration.
	r := chi.NewRouter()

	// Load the .env file.
	err := godotenv.Load()
	if err != nil {
		if !os.IsNotExist(err) {
			log.Fatalf("%s failed reading .env: %v",
				color.RedString("INIT"),
				err,
			)
		}
	} else {
		log.Printf("%s .env loaded",
			color.BlueString("INIT"),
		)
	}

	// Get the configuration.
	cfg := config.Get()

	// HTTP server configuration.
	addr := fmt.Sprintf("%s:%v", cfg.Host, cfg.Port)
	providers, err := config.GetProviders()
	if err != nil {
		log.Fatalf("%s failed getting providers: %v",
			color.RedString("INIT"),
			err,
		)
	}

	// Get the registry.
	registrar := registry.NewRegistry(providers)

	// Proxy Functionality.
	r.Group(func(r chi.Router) {
		r.Use(
			middleware.RequestID,
			middleware.RealIP,
			middleware.Logger,
			middleware.Recoverer,
		)

		r.Route("/oauth", func(r chi.Router) {
			r.Post("/{provider_id}", func(rw http.ResponseWriter, req *http.Request) {
				// Get the provider id from the path.
				providerId := identity.ProviderId(chi.URLParam(req, "provider_id"))
				// Forward the request to the right provider.
				registrar.ProxyServeHTTP(providerId, rw, req)
			})
		})

		// Get the supported providers.
		r.Get("/supported", func(rw http.ResponseWriter, req *http.Request) {
			providers := registrar.Providers()

			enc := json.NewEncoder(rw)
			if err := enc.Encode(providers); err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
			}
		})
	})

	// Static Svelte website files hosting.
	websiteHandler := http.FileServer(http.FS(websiteFS))
	r.Group(func(r chi.Router) {
		r.Use(
			middleware.Compress(5),
		)
		r.Handle("/*", websiteHandler)
	})

	serverChan := make(chan error)

	// Start the non-blocking server.
	go func() {
		log.Printf("%s %v",
			color.BlueString("PROVIDERS"),
			registrar.Providers(),
		)

		log.Printf("%s starting listening on %v",
			color.BlueString("SERVER"),
			color.CyanString("http://"+addr),
		)

		serverChan <- http.ListenAndServe(addr, r)

		log.Printf("%s %s",
			color.RedString("SERVER"),
			"exited",
		)
	}()

	// Wait for the server error.
	select {
	case err := <-serverChan:
		log.Printf("%s %v\n",
			color.RedString("SERVER"),
			err,
		)
	}
}
