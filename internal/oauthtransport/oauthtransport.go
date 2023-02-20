package oauthtransport

import (
	"log"
	"net/http"
)

type (
	oauthTransport struct{}
)

func NewOauthTransport() http.RoundTripper {
	return &oauthTransport{}
}

func (t *oauthTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Custom functions can be added here for logging or modifying the request.
	res, err := http.DefaultTransport.RoundTrip(req)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}
