package proxy

import (
	"github.com/gume1a/oauthproxy/internal/oauthtransport"
	"github.com/gume1a/oauthproxy/internal/provider"
	"net/http"
	"net/http/httputil"
)

type (
	oauthProxy struct {
		provider provider.Provider
		proxy    *httputil.ReverseProxy
	}

	OAuthProxy interface {
		ServeHTTP(rw http.ResponseWriter, req *http.Request)
	}
)

func New(provider provider.Provider) OAuthProxy {
	proxy := httputil.NewSingleHostReverseProxy(provider.TokenEndpoint)
	defaultDirector := proxy.Director

	// Director handles the modification of the request before it is sent to the OAut2 server.
	proxy.Director = func(req *http.Request) {
		defaultDirector(req)
		req.Header.Add("X-Proxy", "oauthproxy")
		req.Host = provider.GetEndpointHost()
	}

	// Handling of the response.
	proxy.ModifyResponse = func(res *http.Response) error {
		return nil
	}

	// Override the round trip function.
	proxy.Transport = oauthtransport.NewOauthTransport()

	return &oauthProxy{
		provider: provider,
		proxy:    proxy,
	}
}

func (p *oauthProxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Drop the URL path, so that it does not append to the provider token endpoint.
	// If you remove this line, the provider path and proxy path will be joined.
	req.URL.Path = ""
	req.RequestURI = ""

	// Add the client secret to the request.
	err := p.provider.AddSecret(req)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		_, _ = rw.Write([]byte(err.Error()))
		return
	}

	// Forward the request to the proxy
	p.proxy.ServeHTTP(rw, req)
}
