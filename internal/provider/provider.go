package provider

import (
	"github.com/gume1a/oauthproxy/pkg/identity"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type (
	Provider struct {
		Id            identity.ProviderId
		TokenEndpoint *url.URL
		clientSecret  string
	}
)

func NewProvider(id identity.ProviderId, tokenURL *url.URL, secret string) Provider {
	return Provider{
		Id:            id,
		TokenEndpoint: tokenURL,
		clientSecret:  secret,
	}
}

func (p *Provider) AddSecret(req *http.Request) error {
	// Add the oauthclient secret to the query parameters.
	err := req.ParseForm()
	if err != nil {
		return err
	}

	req.PostForm.Add("client_secret", p.clientSecret)
	encodedBody := req.PostForm.Encode()
	req.Body = io.NopCloser(strings.NewReader(encodedBody))
	req.ContentLength = int64(len(encodedBody))

	return nil
}

func (p *Provider) GetEndpointHost() string {
	return strings.TrimPrefix(p.TokenEndpoint.Host, "www.")
}
