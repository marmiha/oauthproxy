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

	authUrl, tokenResponseChan := clt.AuthCodeURL(
		"5ca75bd30",
		oauth2.AccessTypeOffline,
	)

	fmt.Printf("Please authenticate on the following url:\n%s\n", authUrl)
	_ = open.Start(authUrl)

	tokenResponse := <-tokenResponseChan
	if err := tokenResponse.Err; err != nil {
		log.Fatal(err)
	}

	log.Printf("Successfully authenticated")
	token := tokenResponse.Token

	log.Printf("Got the access token: %v", token.AccessToken)
}
