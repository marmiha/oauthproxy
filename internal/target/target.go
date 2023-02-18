package target

import (
	"net/url"
	"strings"

	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

var (
	GITHUB Provider = NewGithub()
	GOOGLE Provider = NewGoogle()
)

type (
	Provider struct {
		TokenEndpoint *url.URL
		RequestHost   string
	}
)

func NewGithub() Provider {
	strTarget := github.Endpoint.TokenURL
	target, _ := url.Parse(strTarget)

	return Provider{
		TokenEndpoint: target,
		RequestHost:   strings.TrimPrefix(target.Host, "www."),
	}
}

func NewGoogle() Provider {
	strTarget := google.Endpoint.TokenURL
	target, _ := url.Parse(strTarget)

	return Provider{
		TokenEndpoint: target,
		RequestHost:   strings.TrimPrefix(target.Host, "www."),
	}
}
