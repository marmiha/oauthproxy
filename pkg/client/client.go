package client

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
	"log"
	"net/http"
	"net/url"
)

type (
	client struct {
		config      *oauth2.Config
		callback    func(*oauth2.Token, error)
		transport   *http.Transport
		redirectURL *url.URL
	}

	Client interface {
		AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) (string, chan string, chan error)
		AuthCodeURLCtx(ctx context.Context, state string, opts ...oauth2.AuthCodeOption) (string, chan string, chan error)
	}
)

func (c *client) AuthCodeURLCtx(ctx context.Context, state string, opts ...oauth2.AuthCodeOption) (string, chan string, chan error) {
	if c.transport != nil {
		ctx = context.WithValue(
			ctx,
			oauth2.HTTPClient,
			&http.Client{Transport: c.transport},
		)
	}

	shutdownCallbackServ := make(chan struct{})
	errChan := make(chan error)
	tokenChan := make(chan string)

	mux := http.NewServeMux()
	mux.HandleFunc(c.redirectURL.Path, func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		tok, err := c.config.Exchange(
			ctx,
			code,
		)

		go func() {
			if err != nil {
				errChan <- err
			} else {
				tokenChan <- tok.AccessToken
			}
		}()

		_, _ = fmt.Fprintf(w, "Success! You can close this window now.")
		shutdownCallbackServ <- struct{}{}
	})

	srv := &http.Server{
		Addr:    c.redirectURL.Host,
		Handler: mux,
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if err != http.ErrServerClosed {
				panic(err)
			}
		}
	}()

	go func() {
		<-shutdownCallbackServ
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(shutdownCallbackServ)
	}()

	return c.config.AuthCodeURL(state, opts...), tokenChan, errChan
}

func (c *client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) (string, chan string, chan error) {
	return c.AuthCodeURLCtx(context.Background(), state, opts...)
}
