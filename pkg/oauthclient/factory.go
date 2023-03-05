package oauthclient

import (
	"crypto/tls"
	"fmt"
	"github.com/gume1a/oauthproxy/pkg/identity"
	"golang.org/x/oauth2"
	"net/http"
	"net/url"
)

type (
	factory struct {
		providerID identity.ProviderId
		clientID   string
		scopes     []string

		authURL     *url.URL
		proxyURL    *url.URL
		redirectURL *url.URL

		callback  func(*oauth2.Token, error)
		transport *http.Transport
	}

	Factory interface {
		WithScopes(scopes []string) Factory
		WithProxyURL(url *url.URL) Factory
		WithRedirectURL(url *url.URL) Factory
		WithAuthURL(url *url.URL) Factory
		WithTransport(transport *http.Transport) Factory
		WithInsecureSkipVerifyTransport() Factory
		Build() Client
	}
)

func NewFactory(providerID identity.ProviderId, clientID string) Factory {
	return &factory{
		providerID: providerID,
		clientID:   clientID,
	}
}

func (f *factory) Build() Client {

	if f.redirectURL == nil {
		f.redirectURL = &url.URL{
			Scheme: "http",
			Host:   "localhost:1420",
			Path:   "/oauth2/callback",
		}
	}

	f.proxyURL.Path = fmt.Sprintf("/oauth/%s", f.providerID)

	return &client{
		transport:   f.transport,
		redirectURL: f.redirectURL,

		config: &oauth2.Config{
			ClientID: f.clientID,
			Scopes:   f.scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  f.authURL.String(),
				TokenURL: f.proxyURL.String(),
			},
			RedirectURL: f.redirectURL.String(),
		},
	}
}

func (f *factory) WithScopes(scopes []string) Factory {
	f.scopes = scopes
	return f
}

func (f *factory) WithProxyURL(url *url.URL) Factory {
	f.proxyURL = url
	return f
}

func (f *factory) WithAuthURL(url *url.URL) Factory {
	f.authURL = url
	return f
}

func (f *factory) WithRedirectURL(url *url.URL) Factory {
	f.redirectURL = url
	return f
}

func (f *factory) WithTransport(transport *http.Transport) Factory {
	f.transport = transport
	return f
}

func (f *factory) WithInsecureSkipVerifyTransport() Factory {
	f.transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	return f
}
