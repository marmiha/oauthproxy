package main

import (
	"fmt"
	"github.com/gume1a/oauthproxy/pkg/client"
	"github.com/gume1a/oauthproxy/pkg/identity"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"net/url"
)

func main() {
	redirectURL, _ := url.Parse("http://localhost:1420/oauth2/callback")
	proxyURL, _ := url.Parse("http://localhost:8081")
	authURL, _ := url.Parse(github.Endpoint.AuthURL)

	clt := client.
		NewFactory(identity.GITHUB, "242f79440a257b6370b8").
		WithScopes([]string{"read:org"}).
		WithRedirectURL(redirectURL).
		WithProxyURL(proxyURL).
		WithAuthURL(authURL).
		Build()

	authUrl, tokenChan, errChan := clt.AuthCodeURL(
		"5ca75bd30",
		oauth2.AccessTypeOffline,
	)

	fmt.Printf("Please authenticate on the following url:\n%s\n", authUrl)
	_ = open.Start(authUrl)

	log.Printf("waiting for user authentication")
	var token string
	select {
	case token = <-tokenChan:
		log.Printf("Successfully authenticated")
	case err := <-errChan:
		log.Fatal(err)
	}

	log.Printf(token)
}
