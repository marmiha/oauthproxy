package provider

import (
	"github.com/gume1a/oauth-proxy/pkg/identity"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"net/url"
)

var (
	GITHUB = NewGithub("")
	GOOGLE = NewGoogle("")
)

func NewGithub(secret string) Provider {
	tokenURL, _ := url.Parse(github.Endpoint.TokenURL)
	return NewProvider(identity.GITHUB, tokenURL, secret)
}

func NewGoogle(secret string) Provider {
	tokenURL, _ := url.Parse(google.Endpoint.TokenURL)
	return NewProvider(identity.GOOGLE, tokenURL, secret)
}
