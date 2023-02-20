package provider

import (
	"github.com/gume1a/oauthproxy/pkg/identity"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"net/url"
)

func NewGithub(secret string) Provider {
	tokenURL, _ := url.Parse(github.Endpoint.TokenURL)
	return NewProvider(identity.GITHUB, tokenURL, secret)
}

func NewGoogle(secret string) Provider {
	tokenURL, _ := url.Parse(google.Endpoint.TokenURL)
	return NewProvider(identity.GOOGLE, tokenURL, secret)
}
