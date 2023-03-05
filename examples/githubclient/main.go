package main

import (
	"fmt"
	"github.com/gume1a/oauthproxy/pkg/defaults"
	"github.com/gume1a/oauthproxy/pkg/identity"
	"github.com/gume1a/oauthproxy/pkg/oauthclient"
	"github.com/skratchdot/open-golang/open"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"log"
	"net/url"
)

func main() {
	redirectURL, _ := url.Parse("http://127.0.0.1:1420/oauth2/callback")
	proxyURL, _ := url.Parse(fmt.Sprintf("http://%s:%d", defaults.DefaultHost, defaults.DefaultPort))
	authURL, _ := url.Parse(github.Endpoint.AuthURL)

	clt := oauthclient.
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
