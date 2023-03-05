package oauthclient

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

	TokenResponse struct {
		Token *oauth2.Token
		Err   error
	}

	Client interface {
		AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) (string, chan TokenResponse)
		AuthCodeURLCtx(ctx context.Context, state string, opts ...oauth2.AuthCodeOption) (string, chan TokenResponse)
		GetClient(ctx context.Context, token *oauth2.Token) *http.Client
	}
)

func (c *client) GetClient(ctx context.Context, token *oauth2.Token) *http.Client {
	return c.config.Client(ctx, token)
}

func (c *client) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) (string, chan TokenResponse) {
	return c.AuthCodeURLCtx(context.Background(), state, opts...)
}

func (c *client) AuthCodeURLCtx(ctx context.Context, state string, opts ...oauth2.AuthCodeOption) (string, chan TokenResponse) {
	if c.transport != nil {
		ctx = context.WithValue(
			ctx,
			oauth2.HTTPClient,
			&http.Client{Transport: c.transport},
		)
	}

	shutdownCallbackServ := make(chan struct{})
	tokenChan := make(chan TokenResponse)

	mux := http.NewServeMux()
	mux.HandleFunc(c.redirectURL.Path, func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		token, err := c.config.Exchange(
			ctx,
			code,
		)

		go func() {
			if err != nil {
				tokenChan <- TokenResponse{Err: err}
			} else {
				tokenChan <- TokenResponse{Token: token}
			}
		}()

		_, _ = fmt.Fprintf(w, "Successfully authenticated! You can close this window now.")
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

	return c.config.AuthCodeURL(state, opts...), tokenChan
}
